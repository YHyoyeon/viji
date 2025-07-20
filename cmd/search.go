// cmd 패키지의 검색 명령어를 정의합니다
package cmd

import (
	"fmt"  // 표준 출력/입력 포맷팅을 위한 패키지
	"strings"  // 문자열 조작을 위한 패키지

	"github.com/spf13/cobra"  // CLI 명령어 프레임워크
	"github.com/spf13/viper"  // 설정 관리 라이브러리
	"vi-assistant/internal/search"  // 검색 기능을 위한 내부 패키지
)

// searchCmd는 vi 명령어 검색을 위한 Cobra 명령어입니다
// 사용자가 키워드를 입력하면 관련된 vi 명령어들을 검색하여 표시합니다
var searchCmd = &cobra.Command{
	Use:   "search [keyword]",  // 명령어 사용법 - keyword는 필수 인수
	Short: "키워드로 vi 명령어를 검색합니다",  // 짧은 설명
	Long: `키워드를 사용하여 vi/vim 명령어를 검색합니다.

검색은 다음 항목에서 수행됩니다:
- 명령어 키워드
- 실제 명령어
- 설명
- 카테고리

사용 예시:
  vi-assistant search copy
  vi-assistant search save
  vi-assistant search navigation`,  // 긴 설명 (도움말에 표시됨)
	Args: cobra.ExactArgs(1),  // 정확히 1개의 인수가 필요함을 지정
	Run: func(cmd *cobra.Command, args []string) {
		// 명령어 실행 시 호출되는 함수
		keyword := args[0]  // 첫 번째 인수를 키워드로 사용
		lang := viper.GetString("lang")  // 현재 언어 설정을 가져옴

		// 검색 기능을 실행합니다
		results, err := search.Search(keyword)
		if err != nil {
			fmt.Printf("검색 오류: %v\n", err)  // 오류 발생 시 메시지 출력
			return
		}

		// 검색 결과를 포맷팅하여 출력합니다
		output := search.FormatSearchResults(results, lang)
		fmt.Print(output)
	},
} 