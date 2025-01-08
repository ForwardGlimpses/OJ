package service

import (
	"bytes"
	"context"
	"encoding/json" // 导入 json 包
	"fmt"
	"io"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/config"
	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	model "github.com/criyle/go-judge/cmd/go-judge/model"
	//"github.com/google/uuid"
)

type ProblemServiceInterface interface {
	Query(params schema.ProblemParams) (schema.ProblemItems, int64, error)
	Get(id int) (*schema.ProblemItem, error)
	Create(item *schema.ProblemItem) (int, error)
	Update(id int, item *schema.ProblemItem) error
	Delete(id int) error
	//TODO 提交 判断 存储提交结果 返回提交结果id给前端 前端拿提交的id来查询status
	Submit(id int, userId int, inputCode string) (int, error)
}

var ProblemSvc ProblemServiceInterface = &ProblemService{}

type ProblemService struct{}

// Query根据条件和分页查询获取用户列表
func (a *ProblemService) Query(params schema.ProblemParams) (schema.ProblemItems, int64, error) {
	// 初始化查询
	query := global.DB.Model(&schema.ProblemDBItem{})

	// 应用过滤条件
	if params.ProblemID != 0 {
		query = query.Where("Problem_id = ?", params.ProblemID)
	}

	// 使用通用分页函数并指定返回类型
	problems, total, err := gormx.GetPaginatedData[schema.ProblemDBItem](query, params.P, "id ASC")
	if err != nil {
		logs.Error("Failed to query problems:", err)
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.ProblemItems
	for _, problem := range problems {
		items = append(items, problem.ToItem())
	}

	return items, total, nil
}

// Get 通过ID从数据库获取题目
func (a *ProblemService) Get(id int) (*schema.ProblemItem, error) {
	db := global.DB.WithContext(context.Background())
	item := &schema.ProblemDBItem{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		logs.Error("Failed to get problem with ID:", id, "Error:", err)
		return nil, err
	}
	logs.Info("Successfully retrieved problem with ID:", id)
	return item.ToItem(), nil
}

// Create 将 ProblemItem 转换为 ProblemDBItem 并存入数据库
func (a *ProblemService) Create(item *schema.ProblemItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	dbItem := item.ToDBItem()
	err := db.Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create problem:", err)
		return 0, err
	}
	return dbItem.ID, nil
}

// Update 更新题目信息
func (a *ProblemService) Update(id int, item *schema.ProblemItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		logs.Error("Failed to update problem with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 根据ID删除题目
func (a *ProblemService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ProblemDBItem{})
	if err.Error != nil {
		logs.Error("Failed to delete problem with ID:", id, "Error:", err.Error)
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("no record found with id %d", id)
	}
	return nil
}

func (a *ProblemService) Submit(id int, userId int, inputCode string) (int, error) {

	//db := global.DB.WithContext(context.Background())

	// Step 1: 获取题目详细信息
	problem, err := a.Get(id)
	if err != nil {
		logs.Error("Failed to get problem with ID:", id, "Error:", err)
		return 0, err
	}

	submission := &schema.SolutionItem{
		ProblemID: id,
		UserID:    userId,
		Status:    "Pending",
	}

	submissionId, err := SolutionSvc.Create(submission)
	if err != nil {
		logs.Error("Failed to create submission:", err)
		return 0, err
	}

	fileName := fmt.Sprintf("%d", submissionId)
	fileNameWithExtension := fileName + ".cc"
	content := problem.Input
	Stdout := "stdout"
	Stderr := "stderr"
	StoutMax := int64(10240)
	StderrMax := int64(10240)

	// Step 2: 准备发送给 Judge 的请求体
	judgeRequest1 := model.Request{
		Cmd: []model.Cmd{
			{
				Args: []string{"/usr/bin/g++", fileNameWithExtension, "-o", fileName},
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*model.CmdFile{
					{Content: &content},
					{Name: &Stdout, Max: &StoutMax},
					{Name: &Stderr, Max: &StderrMax},
				},
				CPULimit:    uint64(10 * time.Second), // 10 seconds
				MemoryLimit: 104857600,                // 100 MB
				ProcLimit:   50,
				CopyIn: map[string]model.CmdFile{
					fileNameWithExtension: {
						Content: &inputCode,
					},
				},
				CopyOut:       []string{"stdout", "stderr"},
				CopyOutCached: []string{fileName}},
		},
	}
	//fmt.Println(fileNameWithExtension, "---", fileName)

	// Step 3: 发送请求到 Judge 系统
	body1, err := marshalToReader(judgeRequest1)
	if err != nil {
		logs.Error("Failed to marshal request body:", err)
		return 0, err
	}
	ojBaseURL := config.C.OJ.BaseURL()
	resp1, err := global.HttpClient.Post(ojBaseURL+"/run", "application/json", body1)

	if err != nil {
		logs.Error("Failed to send request to Judge system:", err)
		return 0, err
	}
	defer resp1.Body.Close()

	bodya1, err := io.ReadAll(resp1.Body)
	if err != nil {
		logs.Error("Failed to read response body:", err)
		return 0, err
	}

	logs.Info("Response Body:", string(bodya1)) //记录返回的 JSON 数据

	var judgeResponses1 []model.Result
	err = json.Unmarshal(bodya1, &judgeResponses1)
	if err != nil {
		logs.Error("Error parsing JSON response:", err)
		return 0, err
	}

	aValue := judgeResponses1[0].FileIDs[fileName]
	judgeRequest2 := model.Request{
		Cmd: []model.Cmd{
			{
				Args: []string{fileName},
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*model.CmdFile{
					{Content: &content},
					{Name: &Stdout, Max: &StoutMax},
					{Name: &Stderr, Max: &StderrMax},
				},
				CPULimit:    uint64(10 * time.Second), // 10 seconds
				MemoryLimit: 104857600,                // 100 MB
				ProcLimit:   50,
				CopyIn: map[string]model.CmdFile{
					fileName: {
						FileID: &aValue,
					},
				},
			},
		},
	}

	body2, err := marshalToReader(judgeRequest2)
	if err != nil {
		logs.Error("Failed to marshal request body:", err)
		return 0, err
	}

	resp2, err := global.HttpClient.Post(ojBaseURL+"/run", "application/json", body2)

	if err != nil {
		logs.Error("Failed to send request to Judge system:", err)
		return 0, err
	}
	defer resp2.Body.Close()

	bodya2, err := io.ReadAll(resp2.Body)
	if err != nil {
		logs.Error("Failed to read response body:", err)
		return 0, err
	}

	logs.Info("Response Body:", string(bodya2)) //记录返回的 JSON 数据
	// Step 4: 解析 Judge 系统返回的结果
	var judgeResponse2 []model.Result
	err = json.Unmarshal(bodya2, &judgeResponse2)
	if err != nil {
		logs.Error("Error parsing JSON response:", err)
		return 0, err
	}

	// Step 5: 对比评测结果和标准答案
	contentOutput := judgeResponse2[0].Files["stdout"]
	isCorrect := contentOutput == problem.Output
	status := "Accepted"
	if !isCorrect {
		status = "Wrong Answer"
	}
	fmt.Println("程序输出", contentOutput)
	fmt.Println("样例输出", problem.Output)

	// Step 6: 将评测结果存储到数据库
	submission = &schema.SolutionItem{
		ProblemID: id,
		UserID:    userId,
		Status:    status,
		Time:      judgeResponse2[0].Time,
		Memory:    judgeResponse2[0].Memory,
		Indate:    time.Now(),
		Language:  "C",
		Passrate:  judgeResponse2[0].RunTime,
	}
	//改成solution的creat创建
	err = SolutionSvc.Update(submissionId, submission)
	if err != nil {
		logs.Error("Failed to update submission:", err)
		return 0, err
	}

	// Step 7: 返回提交记录的 ID
	return submissionId, nil
}

// 将结构体编码成 JSON 并返回一个 io.Reader
func marshalToReader(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		logs.Error("Failed to marshal JSON:", err)
		return nil, err
	}
	return bytes.NewReader(data), nil
}
