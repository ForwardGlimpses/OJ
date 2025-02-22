package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	router *gin.Engine
	ctx    context.Context
	token  string
)

var _ = ginkgo.Describe("API Integration Test", func() {
	ginkgo.BeforeEach(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		// 初始化数据库
		err = db.AutoMigrate(
			&schema.ContestProblemDBItem{},
			&schema.ContestSolutionDBItem{},
			&schema.ContestDBItem{},
			&schema.ProblemDBItem{},
			&schema.SolutionDBItem{},
			&schema.SourceCodeDBItem{},
			&schema.UsersDBItem{},
		)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		// 初始化 global.DB
		global.DB = db

		router = gin.Default()
		RegisterRoutes(router)

		ctx = context.Background()
	})

	ginkgo.Describe("LoginAPI", func() {
		ginkgo.It("should login a user", func() {
			// 创建用户
			createReq := schema.UsersItem{
				Email:    "testuser@qq.com",
				Password: "password",
				Level:    2,
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/users", bytes.NewReader(createBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			//id := int(createResp["data"].(float64))

			// 登录用户
			loginReq := map[string]string{
				"email":    "testuser@qq.com",
				"password": "password",
				"level":    "2",
			}
			loginBody, _ := json.Marshal(loginReq)
			req, _ = http.NewRequest("POST", "/api/login", bytes.NewReader(loginBody))
			req.Header.Set("Content-Type", "application/json")
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var loginResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &loginResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			token = loginResp["token"].(string)
		})
	})

	ginkgo.Describe("ContestAPI", func() {
		ginkgo.It("should create, get, update, and delete a contest", func() {
			// 创建比赛
			createReq := schema.ContestItem{
				Title: "New Contest",
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/contest", bytes.NewReader(createBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(createResp["data"].(float64))

			// 获取比赛
			req, _ = http.NewRequest("GET", "/api/contest/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新比赛
			updateReq := schema.ContestItem{
				Title: "Updated Contest",
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/contest/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 查询比赛问题
			req, _ = http.NewRequest("GET", "/api/contest", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var queryResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &queryResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(queryResp["data"]).NotTo(gomega.BeNil())

			// 删除比赛
			req, _ = http.NewRequest("DELETE", "/api/contest/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})

	ginkgo.Describe("ContestProblemAPI", func() {
		ginkgo.It("should create, get, update, and delete a contest problem", func() {
			// 创建比赛问题
			createReq := schema.ContestProblemItem{
				ContestID: 1,
				ProblemID: 1,
				Title:     "New Contest Problem",
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/contestProblem", bytes.NewReader(createBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(createResp["data"].(float64))

			// 获取比赛问题
			req, _ = http.NewRequest("GET", "/api/contestProblem/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新比赛问题
			updateReq := schema.ContestProblemItem{
				Title: "Updated Contest Problem",
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/contestProblem/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 查询比赛问题
			req, _ = http.NewRequest("GET", "/api/contestProblem", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var queryResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &queryResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(queryResp["data"]).NotTo(gomega.BeNil())

			// 删除比赛问题
			req, _ = http.NewRequest("DELETE", "/api/contestProblem/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})

	ginkgo.Describe("ContestSolutionAPI", func() {
		ginkgo.It("should create, get, update, and delete a contest solution", func() {
			// 创建比赛解决方案
			createReq := schema.ContestSolutionItem{
				ContestID:  1,
				ProblemID:  1,
				UserID:     1,
				IsAccepted: false,
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/contestSolution", bytes.NewReader(createBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(createResp["data"].(float64))

			// 获取比赛解决方案
			req, _ = http.NewRequest("GET", "/api/contestSolution/"+strconv.Itoa(id)+"/rank", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新比赛解决方案
			updateReq := schema.ContestSolutionItem{
				IsAccepted: true,
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/contestSolution/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 查询比赛问题
			req, _ = http.NewRequest("GET", "/api/contestSolution", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var queryResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &queryResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(queryResp["data"]).NotTo(gomega.BeNil())

			// 删除比赛解决方案
			req, _ = http.NewRequest("DELETE", "/api/contestSolution/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})

	ginkgo.Describe("ProblemAPI", func() {
		ginkgo.It("should create, get, update, and delete a problem", func() {
			// 创建问题
			createReq := schema.ProblemItem{
				Title:       "New Problem",
				Description: "Problem description",
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/problem", bytes.NewReader(createBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(createResp["data"].(float64))

			// 获取问题
			req, _ = http.NewRequest("GET", "/api/problem/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新问题
			updateReq := schema.ProblemItem{
				Title:       "Updated Problem",
				Description: "Updated description",
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/problem/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 查询比赛问题
			req, _ = http.NewRequest("GET", "/api/problem", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var queryResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &queryResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(queryResp["data"]).NotTo(gomega.BeNil())

			// 删除问题
			req, _ = http.NewRequest("DELETE", "/api/problem/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})

	ginkgo.Describe("SolutionAPI", func() {
		ginkgo.It("should create, get, update, and delete a solution", func() {
			// 创建解决方案
			createReq := schema.SolutionItem{
				ProblemID: 1,
				UserID:    1,
				Language:  "Python",
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/solution", bytes.NewReader(createBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(createResp["data"].(float64))

			// 获取解决方案
			req, _ = http.NewRequest("GET", "/api/solution/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新解决方案
			updateReq := schema.SolutionItem{
				Language: "C",
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/solution/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 查询比赛问题
			req, _ = http.NewRequest("GET", "/api/solution", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var queryResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &queryResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(queryResp["data"]).NotTo(gomega.BeNil())

			// 删除解决方案
			req, _ = http.NewRequest("DELETE", "/api/solution/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})

	ginkgo.Describe("SourceCodeAPI", func() {
		ginkgo.It("should create, get, update, and delete source code", func() {
			// 创建源代码
			createReq := schema.SourceCodeItem{
				Source: "print('Hello, world!')",
			}
			createBody, _ := json.Marshal(createReq)
			req, _ := http.NewRequest("POST", "/api/sourceCode", bytes.NewReader(createBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var createResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &createResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(createResp["data"].(float64))

			// 获取源代码
			req, _ = http.NewRequest("GET", "/api/sourceCode/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新源代码
			updateReq := schema.SourceCodeItem{
				Source: "print('Updated code')",
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/sourceCode/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 查询比赛问题
			req, _ = http.NewRequest("GET", "/api/sourceCode", nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var queryResp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &queryResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(queryResp["data"]).NotTo(gomega.BeNil())

			// 删除源代码
			req, _ = http.NewRequest("DELETE", "/api/sourceCode/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})

	ginkgo.Describe("UsersAPI", func() {
		ginkgo.It("should create, get, update, and delete a user", func() {
			// 使用 Register 方法创建用户
			registerReq := schema.UsersItem{
				Email:    "testuser@qq.com",
				Password: "password",
			}
			registerBody, _ := json.Marshal(registerReq)
			req, _ := http.NewRequest("POST", "/api/users/register", bytes.NewReader(registerBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			var registerResp map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &registerResp)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			id := int(registerResp["data"].(float64))

			// 获取用户
			req, _ = http.NewRequest("GET", "/api/users/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 更新用户
			updateReq := schema.UsersItem{
				Email: "updateduser@qq.com",
			}
			updateBody, _ := json.Marshal(updateReq)
			req, _ = http.NewRequest("PUT", "/api/users/"+strconv.Itoa(id), bytes.NewReader(updateBody))
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

			// 删除用户
			req, _ = http.NewRequest("DELETE", "/api/users/"+strconv.Itoa(id), nil)
			req.AddCookie(&http.Cookie{Name: "token", Value: token})
			req.Header.Set("Authorization", "Bearer "+token) // 添加 token
			w = httptest.NewRecorder()
			router.ServeHTTP(w, req)
			gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		})
	})
})

func TestIntegration(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "API Integration Test Suite")
}

func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		RegisterProblem(apiGroup)
		RegisterContestSolution(apiGroup)
		RegisterContestProblem(apiGroup)
		RegisterContest(apiGroup)
		RegisterSolution(apiGroup)
		RegisterSourceCode(apiGroup)
		RegisterUsers(apiGroup)
		RegisterLogin(apiGroup)
	}
}
