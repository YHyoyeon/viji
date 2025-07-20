// cmd 패키지는 CLI 명령어들을 정의하고 관리합니다
// Cobra 프레임워크를 사용하여 강력한 CLI 인터페이스를 제공합니다
package cmd

import (
	"fmt"  // 표준 출력/입력 포맷팅을 위한 패키지
	"os"   // 운영체제 인터페이스를 위한 패키지

	"github.com/spf13/cobra"  // CLI 명령어 프레임워크
	"github.com/spf13/viper"  // 설정 관리 라이브러리
	"vi-assistant/internal/learn"  // 학습 모드 기능을 위한 내부 패키지
)

// 전역 변수들 - CLI 플래그와 설정을 저장합니다
var (
	cfgFile string  // 설정 파일 경로를 저장하는 변수
	lang    string  // 출력 언어 설정 (ko/en)을 저장하는 변수
	learnLevel string  // 학습 모드 레벨 (beginner/intermediate)을 저장하는 변수
)

// rootCmd는 하위 명령어 없이 호출될 때의 기본 명령어를 나타냅니다
// 애플리케이션의 메인 진입점 역할을 합니다
var rootCmd = &cobra.Command{
	Use:   "vi-assistant",  // 명령어 사용법
	Short: "vi/vim 명령어 도우미 CLI 도구",  // 짧은 설명
	Long: `Vi Assistant는 vi/vim 명령어를 빠르게 검색하고 학습할 수 있는 CLI 도구입니다.

주요 기능:
- 🔍 키워드로 vi 명령어 검색
- 📖 명령어 상세 설명 및 예제
- 🎓 단계별 학습 모드
- ⭐ 즐겨찾기 기능
- 🌍 한국어/영어 지원

사용 예시:
  vi-assistant search copy
  vi-assistant explain :wq
  vi-assistant --learn beginner
  vi-assistant help`,  // 긴 설명 (도움말에 표시됨)
	Version: "1.0.0",  // 애플리케이션 버전
}

// Execute 함수는 모든 하위 명령어를 루트 명령어에 추가하고 플래그를 적절히 설정합니다
// 이 함수는 main.go에서 호출되어 CLI 애플리케이션을 실행합니다
func Execute() error {
	return rootCmd.Execute()  // Cobra 명령어 실행
}

// init 함수는 패키지가 로드될 때 자동으로 호출됩니다
// CLI 명령어와 플래그를 초기화하고 설정합니다
func init() {
	// Cobra 초기화 시 설정 파일을 로드하도록 설정
	cobra.OnInitialize(initConfig)

	// 전역 플래그 설정 - 모든 하위 명령어에서 사용 가능한 플래그들
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "설정 파일 (기본값: $HOME/.vi-assistant.yaml)")
	rootCmd.PersistentFlags().StringVar(&lang, "lang", "ko", "출력 언어 (ko/en)")
	rootCmd.PersistentFlags().StringVar(&learnLevel, "learn", "", "학습 모드 시작 (beginner/intermediate)")

	// 로컬 플래그 설정 - 루트 명령어에서만 사용 가능한 플래그
	rootCmd.Flags().BoolP("toggle", "t", false, "도움말 토글")

	// Viper 설정 바인딩 - 플래그 값을 설정으로 연결
	viper.BindPFlag("lang", rootCmd.PersistentFlags().Lookup("lang"))
	viper.BindPFlag("learn", rootCmd.PersistentFlags().Lookup("learn"))

	// 하위 명령어들을 루트 명령어에 추가
	rootCmd.AddCommand(searchCmd)    // 검색 명령어
	rootCmd.AddCommand(explainCmd)   // 설명 명령어
	rootCmd.AddCommand(helpCmd)      // 도움말 명령어
	rootCmd.AddCommand(favoritesCmd) // 즐겨찾기 명령어

	// 학습 모드 플래그 처리 - 명령어 실행 전에 학습 모드가 설정되었는지 확인
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if learnLevel != "" {
			handleLearnMode(learnLevel)  // 학습 모드가 설정되면 해당 함수 호출
		}
	}
}

// initConfig 함수는 설정 파일과 환경 변수를 읽어들입니다
// 애플리케이션 시작 시 자동으로 호출되어 설정을 초기화합니다
func initConfig() {
	if cfgFile != "" {
		// 플래그로 지정된 설정 파일을 사용
		viper.SetConfigFile(cfgFile)
	} else {
		// 홈 디렉토리를 찾습니다
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)  // 오류가 발생하면 프로그램 종료

		// 홈 디렉토리에서 ".vi-assistant" 이름의 설정 파일을 찾습니다 (확장자 없음)
		viper.AddConfigPath(home)      // 설정 파일 경로 추가
		viper.SetConfigType("yaml")    // 설정 파일 타입을 YAML로 설정
		viper.SetConfigName(".vi-assistant")  // 설정 파일 이름 설정
	}

	// 환경 변수를 자동으로 읽어들입니다 (viper_로 시작하는 변수들)
	viper.AutomaticEnv()

	// 설정 파일이 발견되면 읽어들입니다
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "설정 파일 사용:", viper.ConfigFileUsed())
	}
}

// getMessage 함수는 현재 언어 설정에 따라 지역화된 메시지를 반환합니다
// 다국어 지원을 위한 메시지 관리 시스템입니다
func getMessage(key string) string {
	// 언어별 메시지 맵 - 한국어와 영어 메시지를 정의합니다
	messages := map[string]map[string]string{
		"ko": {  // 한국어 메시지들
			"welcome":        "Vi Assistant에 오신 것을 환영합니다! 🚀",
			"search_prompt":  "검색할 키워드를 입력하세요: ",
			"explain_prompt": "설명할 명령어를 입력하세요: ",
			"no_results":     "검색 결과가 없습니다.",
			"error":          "오류가 발생했습니다: ",
		},
		"en": {  // 영어 메시지들
			"welcome":        "Welcome to Vi Assistant! 🚀",
			"search_prompt":  "Enter keyword to search: ",
			"explain_prompt": "Enter command to explain: ",
			"no_results":     "No results found.",
			"error":          "An error occurred: ",
		},
	}

	// 현재 언어 설정을 가져옵니다 (기본값: 한국어)
	currentLang := viper.GetString("lang")
	if currentLang == "" {
		currentLang = "ko"  // 언어가 설정되지 않으면 한국어로 기본 설정
	}

	// 현재 언어의 메시지가 존재하는지 확인하고 반환
	if langMessages, exists := messages[currentLang]; exists {
		if message, exists := langMessages[key]; exists {
			return message
		}
	}

	// 현재 언어에서 메시지를 찾지 못한 경우 한국어로 폴백
	if langMessages, exists := messages["ko"]; exists {
		if message, exists := langMessages[key]; exists {
			return message
		}
	}

	// 메시지를 찾지 못한 경우 키 자체를 반환
	return key
}

// handleLearnMode 함수는 학습 모드 플래그를 처리합니다
// 사용자가 --learn 플래그를 사용했을 때 호출되어 대화형 튜토리얼을 시작합니다
func handleLearnMode(level string) {
	// 현재 언어 설정을 가져옵니다
	lang := viper.GetString("lang")
	
	// 레벨에 따라 다른 튜토리얼을 시작합니다
	switch level {
	case "beginner":  // 초보자 레벨
		// 초보자 강의 목록을 가져옵니다
		lessons := learn.GetBeginnerLessons(lang)
		
		// 시작 메시지를 출력합니다
		if lang == "en" {
			fmt.Println("🎓 Starting Beginner Level Tutorial")
		} else {
			fmt.Println("🎓 초보자 레벨 튜토리얼을 시작합니다")
		}
		
		// 각 강의를 순차적으로 표시합니다
		for i, lesson := range lessons {
			fmt.Print(learn.FormatLesson(lesson, i+1, lang))  // 강의 내용 출력
			if i < len(lessons)-1 {  // 마지막 강의가 아닌 경우
				if lang == "en" {
					fmt.Println("\nPress Enter to continue to next lesson...")
				} else {
					fmt.Println("\n다음 강의로 계속하려면 Enter를 누르세요...")
				}
				fmt.Scanln() // 사용자 입력을 기다립니다
			}
		}
		
		// 완료 메시지를 출력합니다
		if lang == "en" {
			fmt.Println("\n🎉 Congratulations! You've completed the beginner tutorial!")
		} else {
			fmt.Println("\n🎉 축하합니다! 초보자 튜토리얼을 완료했습니다!")
		}
		
	case "intermediate":  // 중급자 레벨
		// 중급자 강의 목록을 가져옵니다
		lessons := learn.GetIntermediateLessons(lang)
		
		// 시작 메시지를 출력합니다
		if lang == "en" {
			fmt.Println("🎓 Starting Intermediate Level Tutorial")
		} else {
			fmt.Println("🎓 중급자 레벨 튜토리얼을 시작합니다")
		}
		
		// 각 강의를 순차적으로 표시합니다
		for i, lesson := range lessons {
			fmt.Print(learn.FormatLesson(lesson, i+1, lang))  // 강의 내용 출력
			if i < len(lessons)-1 {  // 마지막 강의가 아닌 경우
				if lang == "en" {
					fmt.Println("\nPress Enter to continue to next lesson...")
				} else {
					fmt.Println("\n다음 강의로 계속하려면 Enter를 누르세요...")
				}
				fmt.Scanln() // 사용자 입력을 기다립니다
			}
		}
		
		// 완료 메시지를 출력합니다
		if lang == "en" {
			fmt.Println("\n🎉 Congratulations! You've completed the intermediate tutorial!")
		} else {
			fmt.Println("\n🎉 축하합니다! 중급자 튜토리얼을 완료했습니다!")
		}
		
	default:  // 알 수 없는 레벨
		// 오류 메시지를 출력합니다
		if lang == "en" {
			fmt.Printf("Unknown level: %s. Use 'beginner' or 'intermediate'\n", level)
		} else {
			fmt.Printf("알 수 없는 레벨입니다: %s. 'beginner' 또는 'intermediate'를 사용하세요\n", level)
		}
	}
	
	// 학습 모드 완료 후 프로그램을 종료합니다
	os.Exit(0)
} 