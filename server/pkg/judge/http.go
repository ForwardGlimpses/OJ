package judge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/config"
	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/criyle/go-judge/cmd/go-judge/model"
)

type httpClient struct {
	// 需要存储 judge 服务地址
	baseURL string
}

func newHTTPClient() judgeInterface {
	return &httpClient{
		baseURL: fmt.Sprintf("http://%s:%d/run", config.C.Judge.Host, config.C.Judge.Port),
	}
}

func (a *httpClient) Submit(req Request) (Response, error) {
	// 提交 + 判断逻辑

	fileName := fmt.Sprintf("%d", req.ID)
	fileNameWithExtension := fileName + ".cc"
	content := req.Input
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
						Content: &req.Code,
					},
				},
				CopyOut:       []string{"stdout", "stderr"},
				CopyOutCached: []string{fileName}},
		},
	}
	//fmt.Println(fileNameWithExtension, "---", fileName)

	// Step 3: 发送请求到 Judge 系统

	body1, err := json.Marshal(judgeRequest1)
	if err != nil {
		logs.Error("Failed to marshal request body:", err)
		return Response{}, err
	}
	ojBaseURL := config.C.Judge.BaseURL()
	resp1, err := global.HttpClient.Post(ojBaseURL, "application/json", bytes.NewReader(body1))

	if err != nil {
		logs.Error("Failed to send request to Judge system:", err)
		return Response{}, err
	}
	defer resp1.Body.Close()

	bodya1, err := io.ReadAll(resp1.Body)
	if err != nil {
		logs.Error("Failed to read response body:", err)
		return Response{}, err
	}

	logs.Info("Response Body:", string(bodya1)) //记录返回的 JSON 数据

	var judgeResponses1 []model.Result
	err = json.Unmarshal(bodya1, &judgeResponses1)
	if err != nil {
		logs.Error("Error parsing JSON response:", err)
		return Response{}, err
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

	body2, err := json.Marshal(judgeRequest2)
	if err != nil {
		logs.Error("Failed to marshal request body:", err)
		return Response{}, err
	}

	resp2, err := global.HttpClient.Post(ojBaseURL, "application/json", bytes.NewReader(body2))

	if err != nil {
		logs.Error("Failed to send request to Judge system:", err)
		return Response{}, err
	}
	defer resp2.Body.Close()

	bodya2, err := io.ReadAll(resp2.Body)
	if err != nil {
		logs.Error("Failed to read response body:", err)
		return Response{}, err
	}

	logs.Info("Response Body:", string(bodya2)) //记录返回的 JSON 数据
	// Step 4: 解析 Judge 系统返回的结果
	var judgeResponse2 []model.Result
	err = json.Unmarshal(bodya2, &judgeResponse2)
	if err != nil {
		logs.Error("Error parsing JSON response:", err)
		return Response{}, err
	}

	// Step 5: 对比评测结果和标准答案
	contentOutput := judgeResponse2[0].Files["stdout"]
	isCorrect := contentOutput == req.Output
	status := "Accepted"
	if !isCorrect {
		status = "Wrong Answer"
	}
	fmt.Println("程序输出：", contentOutput)
	fmt.Println("样例输出：", req.Output)

	response := Response{
		Status:  status,
		Memory:  judgeResponse2[0].Memory,
		Time:    judgeResponse2[0].Time,
		RunTime: judgeResponse2[0].RunTime,
	}

	return response, nil
	//return Response{}, nil
}
