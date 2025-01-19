package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

// LoginResponse 结构体
type LoginResponse struct {
	Token string `json:"token"`
}

// CreateProblemResponse 结构体
type CreateProblemResponse struct {
	Data int `json:"data"`
}

// CreateUserResponse 结构体
type CreateUserResponse struct {
	Data int `json:"data"`
}

type GetProblemIDResponse struct {
	Success bool `json:"success"`
	Data    struct {
		ProblemID int `json:"ProblemID"`
	} `json:"data"`
}

type SubmitResponse struct {
	Data int `json:"data"`
}

func main() {
	baseURL := "http://127.0.0.1:8080/api"

	// 模拟用户注册
	userID := registerUser(baseURL, schema.UsersItem{
		Name:     "joe",
		Level:    2,
		Email:    "joe@example.com",
		Password: "password123",
		School:   "zzsd",
	})

	// 模拟用户登录
	token := login(baseURL, "joe@example.com", "password123")

	fmt.Println("用户:", userID)

	fmt.Println("Token值:", token)

	// 模拟创建比赛
	contestID := createContest(baseURL, token, schema.ContestItem{
		Title:         "New Contest",
		Private:       "No",
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(2 * time.Hour),
		Password:      "123456",
		Administrator: "admin",
		Description:   "This is a new contest",
	})
	fmt.Println("比赛:", contestID)

	// 模拟创建题目
	problemID := createProblem(baseURL, token, schema.ProblemItem{
		Title:        "New Problem",
		Description:  "This is a new problem",
		Input:        "1 1",
		Output:       "2/n",
		Indate:       time.Now(),
		SampleInput:  "Sample input",
		SampleOutput: "Sample output",
	})
	fmt.Println("问题:", problemID)

	// 模拟创建比赛问题
	contestProblemID := createContestProblem(baseURL, token, schema.ContestProblemItem{
		ContestID: contestID,
		ProblemID: problemID,
		Title:     "New Contest Problem",
		Accepted:  0,
		Submited:  0,
	})
	fmt.Println("比赛问题:", contestProblemID)

	// 测试 Query 方法
	queryContestProblems(baseURL, token)

	problemID = getProblemID(baseURL, token, contestProblemID)

	fmt.Println("获取的问题ID:", problemID)

	// 模拟提交解决方案
	submissionID := submitSolution(baseURL, token, contestID, problemID, userID, "#include <iostream>\nusing namespace std;\nint main() {\nint a, b;\ncin >> a >> b;\ncout << a + b << endl;\n}")
	fmt.Println("提交记录:", submissionID)

	// 获取比赛的实时排名
	getContestRanking(baseURL, token, contestID)
}
func createContest(baseURL, token string, contest schema.ContestItem) int {
	reader, err := marshalToReader(contest)
	if err != nil {
		fmt.Println("Error marshaling contest:", err)
		return 0
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/contest", baseURL), reader)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating contest:", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	fmt.Println("Create contest response:", string(body)) // 添加调试信息

	var createContestResponse CreateProblemResponse
	err = json.Unmarshal(body, &createContestResponse)
	if err != nil {
		fmt.Println("Error unmarshaling create contest response:", err)
		return 0
	}

	fmt.Println("Created Contest:", string(body))
	return createContestResponse.Data
}

func createContestProblem(baseURL, token string, contestProblem schema.ContestProblemItem) int {
	reader, err := marshalToReader(contestProblem)
	if err != nil {
		fmt.Println("Error marshaling contest problem:", err)
		return 0
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/contestProblem", baseURL), reader)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating contest problem:", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	fmt.Println("Create contest problem response:", string(body)) // 添加调试信息

	var createContestProblemResponse CreateProblemResponse
	err = json.Unmarshal(body, &createContestProblemResponse)
	if err != nil {
		fmt.Println("Error unmarshaling create contest problem response:", err)
		return 0
	}

	fmt.Println("Created Contest Problem:", string(body))
	return createContestProblemResponse.Data
}

func getProblemID(baseURL, token string, contestProblemID int) int {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contestProblem/%d", baseURL, contestProblemID), nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting problem ID:", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	fmt.Println("Get problem ID response:", string(body))

	var getProblemIDResponse GetProblemIDResponse

	err = json.Unmarshal(body, &getProblemIDResponse)
	if err != nil {
		fmt.Println("Error unmarshaling get problem ID response:", err)
		return 0
	}

	return getProblemIDResponse.Data.ProblemID
}

func queryContestProblems(baseURL, token string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contestProblem", baseURL), nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error querying contest problems:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Query contest problems response:", string(body))
}

func registerUser(baseURL string, user schema.UsersItem) int {
	reader, err := marshalToReader(user)
	if err != nil {
		fmt.Println("Error marshaling user data:", err)
		return 0
	}

	resp, err := http.Post(fmt.Sprintf("%s/users", baseURL), "application/json", reader)
	if err != nil {
		fmt.Println("Error registering user:", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	fmt.Println("Register response:", string(body)) // 添加调试信息

	var createUserResponse CreateUserResponse
	err = json.Unmarshal(body, &createUserResponse)
	if err != nil {
		fmt.Println("Error unmarshaling create user response:", err)
		return 0
	}

	fmt.Println("Registered User:", string(body))
	return createUserResponse.Data
}

func login(baseURL, email, password string) string {
	data := map[string]string{
		"email":    email,
		"password": password,
	}
	reader, err := marshalToReader(data)
	if err != nil {
		fmt.Println("Error marshaling login data:", err)
		return ""
	}

	resp, err := http.Post(fmt.Sprintf("%s/login", baseURL), "application/json", reader)
	if err != nil {
		fmt.Println("Error logging in:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	fmt.Println("Login response:", string(body)) // 添加调试信息

	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Println("Error unmarshaling login response:", err)
		return ""
	}

	fmt.Println("Logged in successfully, token:", loginResponse.Token)
	return loginResponse.Token
}

func createProblem(baseURL, token string, problem schema.ProblemItem) int {
	reader, err := marshalToReader(problem)
	if err != nil {
		fmt.Println("Error marshaling problem:", err)
		return 0
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/problem", baseURL), reader)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating problem:", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	fmt.Println("Create problem response:", string(body)) // 添加调试信息

	var createProblemResponse CreateProblemResponse
	err = json.Unmarshal(body, &createProblemResponse)
	if err != nil {
		fmt.Println("Error unmarshaling create problem response:", err)
		return 0
	}

	fmt.Println("Created Problem:", string(body))
	return createProblemResponse.Data
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

func submitSolution(baseURL, token string, contestID, problemID, userID int, inputCode string) int {
	data := map[string]interface{}{
		"contest_id": contestID,
		"id":         problemID,
		"user_id":    userID,
		"inputcode":  inputCode,
		"indate":     time.Now(),
	}
	reader, err := marshalToReader(data)
	if err != nil {
		fmt.Println("Error marshaling submission data:", err)
		return 0
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/contestSolution/%d", baseURL, problemID), reader)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error submitting solution:", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	fmt.Println("Submit response:", string(body)) // 添加调试信息

	var submitResponse SubmitResponse

	err = json.Unmarshal(body, &submitResponse)
	if err != nil {
		fmt.Println("Error unmarshaling submit response:", err)
		return 0
	}

	return submitResponse.Data
}

func getContestRanking(baseURL, token string, contestID int) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contestSolution/%d/rank", baseURL, contestID), nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting contest ranking:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Contest ranking response:", string(body))
}
