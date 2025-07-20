// cmd 패키지의 설명 명령어를 정의합니다
package cmd

import (
	"fmt"  // 표준 출력/입력 포맷팅을 위한 패키지

	"github.com/spf13/cobra"  // CLI 명령어 프레임워크
	"github.com/spf13/viper"  // 설정 관리 라이브러리
	"vi-assistant/internal/explain"  // 설명 기능을 위한 내부 패키지
)

// explainCmd는 vi 명령어 설명을 위한 Cobra 명령어입니다
// 사용자가 특정 명령어를 입력하면 상세한 설명과 예제를 제공합니다
var explainCmd = &cobra.Command{
	Use:   "explain [command]",  // 명령어 사용법 - command는 필수 인수
	Short: "vi 명령어에 대한 상세 설명을 제공합니다",  // 짧은 설명
	Long: `특정 vi/vim 명령어에 대한 상세한 설명과 사용 예제를 제공합니다.

명령어를 정확히 입력하면 상세 설명을, 
유사한 명령어가 있으면 제안 목록을 보여줍니다.

사용 예시:
  vi-assistant explain :wq
  vi-assistant explain yy
  vi-assistant explain /pattern`,  // 긴 설명 (도움말에 표시됨)
	Args: cobra.ExactArgs(1),  // 정확히 1개의 인수가 필요함을 지정
	Run: func(cmd *cobra.Command, args []string) {
		// 명령어 실행 시 호출되는 함수
		command := args[0]  // 첫 번째 인수를 명령어로 사용
		lang := viper.GetString("lang")  // 현재 언어 설정을 가져옴

		// 명령어 설명 기능을 실행합니다
		result, err := explain.Explain(command)
		if err != nil {
			fmt.Printf("설명 오류: %v\n", err)  // 오류 발생 시 메시지 출력
			return
		}

		// 설명 결과를 포맷팅하여 출력합니다
		output := explain.FormatExplanation(result, lang)
		fmt.Print(output)
	},
} 