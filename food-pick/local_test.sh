 #!/bin/bash

# 테스트할 함수 이름
FUNCTION_NAME="RecommendFoodsFunction"

# 이벤트 파일 경로
EVENT_FILE="foods/event.json"

# SAM CLI 설치 여부 확인
if ! command -v sam &> /dev/null; then
    echo "SAM CLI가 설치되어 있지 않습니다. 설치 후 다시 시도해주세요."
    exit 1
fi

# 이벤트 파일이 존재하는지 확인
if [ ! -f "$EVENT_FILE" ]; then
    echo "이벤트 파일($EVENT_FILE)이 없습니다. event.json 파일을 생성해주세요."
    exit 1
fi

# SAM local invoke 실행
echo "로컬 테스트를 실행합니다: $FUNCTION_NAME"
sam local invoke "$FUNCTION_NAME" --event "$EVENT_FILE"

# 테스트 결과 출력
if [ $? -eq 0 ]; then
    echo "로컬 테스트 성공: $FUNCTION_NAME"
else
    echo "로컬 테스트 실패: $FUNCTION_NAME"
    exit 1
fi
