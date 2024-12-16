package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/criyle/go-judge/cmd/go-judge/model"
)

func main() {
	content := ""
	Stdout := "stdout"
	Stderr := "stderr"
	StoutMax := int64(10240)
	StderrMax := int64(10240)
	Copycontent := "#include <iostream>\nusing namespace std;\nint main() {\nint a, b;\ncin >> a >> b;\ncout << a + b << endl;\n}"
	// Step 2: 准备发送给 Judge 的请求体
	judgeRequest := model.Request{
		Cmd: []model.Cmd{
			{
				Args: []string{"/usr/bin/g++", "a.cc", "-o", "a"},
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
					"a.cc": {
						Content: &Copycontent,
					},
				},
				CopyOut:       []string{"stdout", "stderr"},
				CopyOutCached: []string{"a"}},
		},
	}
	// Step 3: 发送请求到 Judge 系统
	body, err := marshalToReader(judgeRequest)
	if err != nil {
		panic(err)
	}
	resp, err := global.HttpClient.Post("http://localhost:5050/run", "application/json", body)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodya, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(bodya)) // 打印返回的 JSON 数据

	// Step 4: 解析 Judge 系统返回的结果
	var judgeResponse []model.Result
	if err := json.Unmarshal(bodya, &judgeResponse); err != nil {
		panic(err)
	}
	// Step 5: 对比评测结果和标准答案
	isCorrect := judgeResponse[0].Files["stdout"] == "2"
	status := "Accepted"
	if !isCorrect {
		status = "Wrong Answer"
	}

	fmt.Println(status)
}

// 将结构体编码成 JSON 并返回一个 io.Reader
func marshalToReader(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
