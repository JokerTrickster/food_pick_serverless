#!/bin/bash

# 테스트를 실행할 Lambda 함수와 이벤트 파일 목록
# sam local invoke FoodsFunction --event ./foods/events/dailyRecommend.json
declare -a tests=(
    "FoodsFunction ./foods/events/dailyRecommend.json"
    "FoodsFunction ./foods/events/recommend.json"
)

# 테스트 로그 디렉토리 생성
LOG_DIR="./test_logs"
mkdir -p $LOG_DIR

# 테스트 실행
for test in "${tests[@]}"; do
    # Lambda 함수 이름과 이벤트 파일 추출
    IFS=" " read -r functionName eventFile <<< "$test"

    # 로그 파일 이름 생성
    logFile="$LOG_DIR/${functionName}_$(basename $eventFile .json).log"

    # 테스트 실행
    echo "Invoking $functionName with event $eventFile..."
    sam local invoke "$functionName" --event "$eventFile" > "$logFile" 2>&1

    # 결과 확인
    if [ $? -eq 0 ]; then
        echo "✅ $functionName with $eventFile succeeded. Logs saved to $logFile"
    else
        echo "❌ $functionName with $eventFile failed. Check logs in $logFile"
    fi
done

echo "All tests completed. Check $LOG_DIR for logs."
