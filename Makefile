.PHONY: proto clean

# 默认目标
all: proto

# 生成 protobuf 代码
proto:
	buf mod update
	buf generate

# 清理生成的代码
clean:
	rm -rf gen

# 初始化 buf 模块（首次使用时运行）
init:
	buf mod init

# 检查 proto 文件格式
lint:
	buf lint

# 检查破坏性变更
breaking:
	buf breaking --against '.git#branch=main'

# 帮助信息
help:
	@echo "Available targets:"
	@echo "  make proto     - Generate protobuf code"
	@echo "  make clean    - Remove generated code"
	@echo "  make init     - Initialize buf module"
	@echo "  make lint     - Check proto files format"
	@echo "  make breaking - Check for breaking changes" 