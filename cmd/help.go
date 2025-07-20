// cmd 패키지의 도움말 명령어를 정의합니다
package cmd

import (
	"fmt"  // 표준 출력/입력 포맷팅을 위한 패키지

	"github.com/spf13/cobra"  // CLI 명령어 프레임워크
	"github.com/spf13/viper"  // 설정 관리 라이브러리
	"vi-assistant/internal/explain"  // 설명 기능을 위한 내부 패키지
)

// helpCmd는 vi 명령어 빠른 참조를 위한 Cobra 명령어입니다
// 자주 사용하는 vi 명령어들을 카테고리별로 정리하여 보여줍니다
var helpCmd = &cobra.Command{
	Use:   "help",  // 명령어 사용법
	Short: "자주 사용하는 vi 명령어의 빠른 참조를 제공합니다",  // 짧은 설명
	Long: `vi/vim에서 자주 사용하는 명령어들의 빠른 참조를 제공합니다.

카테고리별로 정리된 명령어 목록과 간단한 설명을 보여줍니다.
초보자가 가장 먼저 알아야 할 명령어들로 구성되어 있습니다.

사용 예시:
  vi-assistant help
  vi-assistant help --lang en`,  // 긴 설명 (도움말에 표시됨)
	Run: func(cmd *cobra.Command, args []string) {
		// 명령어 실행 시 호출되는 함수
		lang := viper.GetString("lang")  // 현재 언어 설정을 가져옴

		// 빠른 참조 내용을 출력합니다
		quickRef := explain.GetQuickReference(lang)
		fmt.Print(quickRef)

		// 추가 도움말 팁을 출력합니다
		if lang == "en" {  // 영어 버전
			fmt.Println("\n💡 Tips:")
			fmt.Println("  - Use 'vi-assistant search <keyword>' to find specific commands")
			fmt.Println("  - Use 'vi-assistant explain <command>' for detailed explanations")
			fmt.Println("  - Use 'vi-assistant --learn beginner' for interactive tutorial")
		} else {  // 한국어 버전
			fmt.Println("\n💡 팁:")
			fmt.Println("  - 'vi-assistant search <키워드>'로 특정 명령어 검색")
			fmt.Println("  - 'vi-assistant explain <명령어>'로 상세 설명 보기")
			fmt.Println("  - 'vi-assistant --learn beginner'로 대화형 튜토리얼 시작")
		}
	},
} 