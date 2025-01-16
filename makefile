# 默认目标
.PHONY: all
all: server

# 启动模拟测试
.PHONY: mockclient
mockclient:
    cd ./server && go run ./tool/mockclient/main.go

# 启动后端服务器
.PHONY: server
server:
    cd ./server && go run ./cmd/main.go start -c ./configs/configs.json