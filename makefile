# 定义变量
FRONTEND_DIR := web
BACKEND_DIR := server
CONFIG_FILE := $(BACKEND_DIR)/configs/config.json

# 默认目标
.PHONY: all
all: frontend backend

# 前端目标
.PHONY: frontend
frontend:
    @echo "Building frontend..."
    cd $(FRONTEND_DIR) && npm install && npm run build

# 后端目标
.PHONY: backend
backend:
    @echo "Building backend..."
    cd $(BACKEND_DIR) && go build -o oj-backend ./cmd/main.go

# 启动前端开发服务器
.PHONY: start-frontend
start-frontend:
    @echo "Starting frontend development server..."
    cd $(FRONTEND_DIR) && npm run serve

# 启动后端服务器
.PHONY: start-backend
start-backend:
    @echo "Starting backend server..."
    cd $(BACKEND_DIR) && go run ./cmd/main.go start -c $(CONFIG_FILE)

# 初始化数据库
.PHONY: initdb
initdb:
    @echo "Initializing database..."
    cd $(BACKEND_DIR)/tool/initdb && go run main.go

# 清理目标
.PHONY: clean
clean:
    @echo "Cleaning up..."
    cd $(FRONTEND_DIR) && rm -rf node_modules dist
    cd $(BACKEND_DIR) && rm -f oj-backend