# AWS Lambda Go Build Makefile

BINARY_NAME=bootstrap# Lambda에서 실행될 바이너리 이름은 항상 bootstrap이어야 합니다.
TARGET_FOLDER=$(ARTIFACTS_DIR)# SAM 빌드에서 전달받는 빌드 아티팩트 경로
SRC_FILE=main.go

.PHONY: build build-FoodsFunction clean

# SAM CLI가 요구하는 함수 이름 기반 타겟 추가
build-FoodsFunction: build

# 빌드 명령어
build:
	GOOS=linux GOARCH=amd64 go build -o $(TARGET_FOLDER)/$(BINARY_NAME) $(SRC_FILE)

# 클린업 명령어
clean:
	rm -rf $(TARGET_FOLDER)/*