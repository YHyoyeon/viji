// vi-assistant CLI 애플리케이션의 메인 패키지
// 이 파일은 애플리케이션의 진입점 역할을 합니다
package main

import (
	"fmt"  // 표준 출력/입력 포맷팅을 위한 패키지
	"os"   // 운영체제 인터페이스를 위한 패키지

	"vi-assistant/cmd"  // CLI 명령어 처리를 위한 내부 패키지
)

// main 함수는 애플리케이션의 시작점입니다
// CLI 명령어를 실행하고 오류가 발생하면 적절히 처리합니다
func main() {
	// cmd.Execute()를 호출하여 CLI 명령어를 실행
	// 오류가 발생하면 stderr에 오류 메시지를 출력하고 종료 코드 1로 종료
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "오류가 발생했습니다: %v\n", err)
		os.Exit(1)  // 비정상 종료를 나타내는 종료 코드
	}
} 