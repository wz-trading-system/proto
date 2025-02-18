.PHONY: proto clean init lint tools

# 默认目标
all: tools proto

# 安装必要的工具
tools:
	GOBIN=$(HOME)/go/bin go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(HOME)/go/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(HOME)/go/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

# 生成 protobuf 代码
proto:
	PATH="$(PATH):$(HOME)/go/bin" buf generate

# 清理生成的代码
clean:
	rm -rf gen

# 初始化或更新 buf 配置
init:
	GOPATH= buf mod update

# 检查 proto 文件格式
lint:
	GOPATH= buf lint

# 帮助信息
help:
	@echo "Available targets:"
	@echo "  make tools    - Install required tools"
	@echo "  make proto    - Generate protobuf code"
	@echo "  make clean    - Remove generated code"
	@echo "  make init     - Update buf dependencies"
	@echo "  make lint     - Check proto files format" 