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

	// 模拟更新比赛
	updateContest(baseURL, token, contestID, schema.ContestItem{
		Title:         "Updated Contest",
		Private:       "Yes",
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(3 * time.Hour),
		Password:      "newpassword",
		Administrator: "admin",
		Description:   "This is an updated contest",
	})

	// 模拟删除比赛
	//deleteContest(baseURL, token, contestID)

	// 模拟创建比赛问题
	contestProblemID := createContestProblem(baseURL, token, schema.ContestProblemItem{
		ContestID: contestID,
		Title:     "New Contest Problem",
		Accepted:  0,
		Submited:  0,
	})
	fmt.Println("比赛问题:", contestProblemID)

	// 模拟更新比赛问题
	updateContestProblem(baseURL, token, contestProblemID, schema.ContestProblemItem{
		ContestID: contestID,
		Title:     "Updated Contest Problem",
		Accepted:  0,
		Submited:  0,
	})

	// 测试 Query 方法
	queryContestProblems(baseURL, token)

	// 模拟删除比赛问题
	//deleteContestProblem(baseURL, token, contestProblemID)
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

func updateContest(baseURL, token string, id int, contest schema.ContestItem) {
	reader, err := marshalToReader(contest)
	if err != nil {
		fmt.Println("Error marshaling contest:", err)
		return
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/contest/%d", baseURL, id), reader)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error updating contest:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Update contest response:", string(body)) // 添加调试信息
}

func deleteContest(baseURL, token string, id int) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/contest/%d", baseURL, id), nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting contest:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Delete contest response:", string(body)) // 添加调试信息
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
func updateContestProblem(baseURL, token string, id int, contestProblem schema.ContestProblemItem) {
	reader, err := marshalToReader(contestProblem)
	if err != nil {
		fmt.Println("Error marshaling contest problem:", err)
		return
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/contestProblem/%d", baseURL, id), reader)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error updating contest problem:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Update contest problem response:", string(body)) // 添加调试信息
}

func deleteContestProblem(baseURL, token string, id int) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/contestProblem/%d", baseURL, id), nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting contest problem:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Delete contest problem response:", string(body)) // 添加调试信息
}

func queryContestProblems(baseURL, token string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contestProblem/", baseURL), nil)
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

// // 模拟创建题目
// problemID := createProblem(baseURL, token, schema.ProblemItem{
// 	Title:        "New Problem",
// 	Description:  "This is a new problem",
// 	Input:        "1 1",
// 	Output:       "2",
// 	SampleInput:  "Sample input",
// 	SampleOutput: "Sample output",
// })
// fmt.Println("问题:", problemID)

// // 模拟更新题目
// updateProblem(baseURL, token, problemID, schema.ProblemItem{
// 	Title:        "Updated Problem",
// 	Description:  "This is an updated problem",
// 	Input:        "2 2",
// 	Output:       "4",
// 	SampleInput:  "Updated sample input",
// 	SampleOutput: "Updated sample output",
// })

// // 模拟提交解决方案
// submitSolution(baseURL, token, schema.Submit{
// 	ID:        problemID,
// 	UserID:    userID,
// 	InputCode: "#include <iostream>\nusing namespace std;\nint main() {\nint a, b;\ncin >> a >> b;\ncout << a + b << endl;\n}",
// })

// // 模拟删除题目
// deleteProblem(baseURL, token, problemID)

// // 模拟更新用户
// updateUser(baseURL, token, userID, schema.UsersItem{
// 	Name:     "John Doe Updated",
// 	Email:    "john.doe@example.com",
// 	Password: "newpassword123",
// 	School:   "Updated School",
// })

// // 模拟删除用户
// deleteUser(baseURL, token, userID)
//}

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

// func createProblem(baseURL, token string, problem schema.ProblemItem) int {
// 	reader, err := marshalToReader(problem)
// 	if err != nil {
// 		fmt.Println("Error marshaling problem:", err)
// 		return 0
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/problem", baseURL), reader)
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return 0
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error creating problem:", err)
// 		return 0
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return 0
// 	}

// 	fmt.Println("Create problem response:", string(body)) // 添加调试信息

// 	var createProblemResponse CreateProblemResponse
// 	err = json.Unmarshal(body, &createProblemResponse)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling create problem response:", err)
// 		return 0
// 	}

// 	fmt.Println("Created Problem:", string(body))
// 	return createProblemResponse.Data
// }

// func updateProblem(baseURL, token string, id int, problem schema.ProblemItem) {
// 	reader, err := marshalToReader(problem)
// 	if err != nil {
// 		fmt.Println("Error marshaling problem:", err)
// 		return
// 	}

// 	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/problem/%d", baseURL, id), reader)
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error updating problem:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("Update problem response:", string(body)) // 添加调试信息
// }

// func deleteProblem(baseURL, token string, id int) {
// 	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/problem/%d", baseURL, id), nil)
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error deleting problem:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("Delete problem response:", string(body)) // 添加调试信息
// }

// func submitSolution(baseURL, token string, submit schema.Submit) {
// 	reader, err := marshalToReader(submit)
// 	if err != nil {
// 		fmt.Println("Error marshaling solution:", err)
// 		return
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/problem/%d", baseURL, submit.ID), reader)
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error submitting solution:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("Submit solution response:", string(body)) // 添加调试信息
// }

// func updateUser(baseURL, token string, id int, user schema.UsersItem) {
// 	reader, err := marshalToReader(user)
// 	if err != nil {
// 		fmt.Println("Error marshaling user:", err)
// 		return
// 	}

// 	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%d", baseURL, id), reader)
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error updating user:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("Update user response:", string(body)) // 添加调试信息
// }

// func deleteUser(baseURL, token string, id int) {
// 	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%d", baseURL, id), nil)
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error deleting user:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("Delete user response:", string(body)) // 添加调试信息
// }

// 将结构体编码成 JSON 并返回一个 io.Reader
func marshalToReader(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		logs.Error("Failed to marshal JSON:", err)
		return nil, err
	}
	return bytes.NewReader(data), nil
}
