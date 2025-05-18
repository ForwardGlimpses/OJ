# OJ 在线判题系统

本项目是一个基于 Go 语言开发的在线判题系统（Online Judge，OJ），支持多种编程语言的自动编译与评测。系统支持 HTTP 和 gRPC 两种通信方式，适用于编程竞赛、在线编程学习等场景。项目结构清晰，易于扩展和维护。

## 主要功能

- 支持多种编程语言的代码自动评测
- 支持 HTTP 和 gRPC 两种判题服务通信协议
- RESTful API，支持代码提交、评测结果查询等操作
- 评测过程自动限时、限内存、限进程数
- 详细的日志记录与错误处理
- 完善的单元测试与集成测试

## 目录结构

```
OJ/
├── server/
│   ├── pkg/
│   │   ├── api/         # API 层，路由与接口
│   │   ├── judge/       # 判题核心逻辑，支持 HTTP/gRPC
│   │   ├── schema/      # 数据库模型
│   │   ├── global/      # 全局变量
│   ├── main.go          # 启动入口
├── README.md
```

## 快速开始

1. **环境准备**
   - Go 1.18 及以上
   - 配置好 go mod 依赖
   - 可选：Docker（用于部署判题后端）

2. **启动判题服务**
   - 判题服务支持 HTTP 和 gRPC 两种方式，需先启动判题后端（如 [go-judge](https://github.com/criyle/go-judge)）。

3. **启动 OJ 服务**
   ```bash
   cd server
   go run main.go
   ```

4. **提交代码评测**
   - 通过 API 或 gRPC 客户端提交代码，具体接口文档见 `api/` 目录或参考代码注释。

## 测试

- 单元测试：覆盖 judge 层、API 层等核心逻辑
- 集成测试：`pkg/api/integration_test.go`，模拟真实 API 调用，验证系统整体功能

运行测试：

```bash
go test ./...
```

## 主要技术栈

- Go
- gRPC / HTTP
- Gin（Web 框架）
- GORM（ORM）
- go-judge（判题后端）
- Ginkgo （测试）

## 贡献

欢迎提交 issue 和 PR，完善功能或修复 bug。

## License

MIT

---

**本项目适合用于高校课程、编程竞赛平台、在线编程学习等场景，欢迎二次开发和定制！**
