package judge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/criyle/go-judge/cmd/go-judge/model"
)

// JudgeRequest 发送给判题机的请求体
type JudgeRequest struct {
	Cmd []model.Cmd `json:"cmd"`
}

// JudgeResponse 判题机的响应体
type JudgeResponse struct {
	Results []model.Result `json:"results"`
}

// SubmitToJudge 提交代码到判题机
func SubmitToJudge(judgeURL string, request JudgeRequest) (JudgeResponse, error) {
	var response JudgeResponse

	data, err := json.Marshal(request)
	if err != nil {
		return response, fmt.Errorf("failed to marshal judge request: %v", err)
	}

	resp, err := http.Post(judgeURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return response, fmt.Errorf("failed to submit to judge: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("failed to read judge response: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("failed to unmarshal judge response: %v", err)
	}

	return response, nil
}
