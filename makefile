main_file = cmd/main.go
binary_name = auth-service
build_dir = ./bin
GO_ENV = CGO_ENABLED=1 GOARCH=amd64 GOOS=linux
install:
	@echo "Hello world"

build:
	@mkdir -p ${build_dir}
	GOARCH=amd64 GOOS=darwin go build -o ${build_dir}/${binary_name}-darwin ${main_file}
	GOARCH=amd64 GOOS=linux go build -o ${build_dir}/${binary_name}-linux ${main_file}

run-build:
	./${build_dir}/${binary_name}-darwin

size:
	@ls -lh $(build_dir)/$(binary_name)-darwin
	@ls -lh $(build_dir)/$(binary_name)-linux

run:
	go run cmd/main.go

clear:
	rm -rf ${build_dir}

