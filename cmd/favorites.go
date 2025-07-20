package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"vi-assistant/internal/favorites"
	"vi-assistant/internal/explain"
)

var favoritesCmd = &cobra.Command{
	Use:   "fav",
	Short: "즐겨찾기 명령어를 관리합니다",
	Long: `자주 사용하는 vi 명령어를 즐겨찾기에 추가하고 관리합니다.

하위 명령어:
  add    - 명령어를 즐겨찾기에 추가
  list   - 즐겨찾기 목록 보기
  remove - 즐겨찾기에서 제거
  clear  - 모든 즐겨찾기 삭제

사용 예시:
  vi-assistant fav add :wq
  vi-assistant fav list
  vi-assistant fav remove :wq`,
}

var favAddCmd = &cobra.Command{
	Use:   "add [command]",
	Short: "명령어를 즐겨찾기에 추가합니다",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		lang := viper.GetString("lang")

		// 명령어 설명 가져오기
		result, err := explain.Explain(command)
		if err != nil {
			fmt.Printf("명령어 확인 오류: %v\n", err)
			return
		}

		if !result.Found {
			if lang == "en" {
				fmt.Printf("Command not found: %s\n", command)
			} else {
				fmt.Printf("명령어를 찾을 수 없습니다: %s\n", command)
			}
			return
		}

		// 즐겨찾기 매니저 생성
		fm, err := favorites.NewFavoritesManager()
		if err != nil {
			fmt.Printf("즐겨찾기 매니저 오류: %v\n", err)
			return
		}

		// 즐겨찾기에 추가
		err = fm.Add(command, result.Command.Description)
		if err != nil {
			fmt.Printf("즐겨찾기 추가 오류: %v\n", err)
			return
		}

		if lang == "en" {
			fmt.Printf("Added '%s' to favorites.\n", command)
		} else {
			fmt.Printf("'%s'을(를) 즐겨찾기에 추가했습니다.\n", command)
		}
	},
}

var favListCmd = &cobra.Command{
	Use:   "list",
	Short: "즐겨찾기 목록을 표시합니다",
	Run: func(cmd *cobra.Command, args []string) {
		lang := viper.GetString("lang")

		// 즐겨찾기 매니저 생성
		fm, err := favorites.NewFavoritesManager()
		if err != nil {
			fmt.Printf("즐겨찾기 매니저 오류: %v\n", err)
			return
		}

		// 즐겨찾기 목록 가져오기
		favList, err := fm.List()
		if err != nil {
			fmt.Printf("즐겨찾기 목록 오류: %v\n", err)
			return
		}

		// 결과 출력
		output := favorites.FormatFavorites(favList, lang)
		fmt.Print(output)
	},
}

var favRemoveCmd = &cobra.Command{
	Use:   "remove [command]",
	Short: "즐겨찾기에서 명령어를 제거합니다",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		lang := viper.GetString("lang")

		// 즐겨찾기 매니저 생성
		fm, err := favorites.NewFavoritesManager()
		if err != nil {
			fmt.Printf("즐겨찾기 매니저 오류: %v\n", err)
			return
		}

		// 즐겨찾기에서 제거
		err = fm.Remove(command)
		if err != nil {
			fmt.Printf("즐겨찾기 제거 오류: %v\n", err)
			return
		}

		if lang == "en" {
			fmt.Printf("Removed '%s' from favorites.\n", command)
		} else {
			fmt.Printf("'%s'을(를) 즐겨찾기에서 제거했습니다.\n", command)
		}
	},
}

var favClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "모든 즐겨찾기를 삭제합니다",
	Run: func(cmd *cobra.Command, args []string) {
		lang := viper.GetString("lang")

		// 즐겨찾기 매니저 생성
		fm, err := favorites.NewFavoritesManager()
		if err != nil {
			fmt.Printf("즐겨찾기 매니저 오류: %v\n", err)
			return
		}

		// 모든 즐겨찾기 삭제
		err = fm.Clear()
		if err != nil {
			fmt.Printf("즐겨찾기 삭제 오류: %v\n", err)
			return
		}

		if lang == "en" {
			fmt.Println("All favorites cleared.")
		} else {
			fmt.Println("모든 즐겨찾기가 삭제되었습니다.")
		}
	},
}

func init() {
	favoritesCmd.AddCommand(favAddCmd)
	favoritesCmd.AddCommand(favListCmd)
	favoritesCmd.AddCommand(favRemoveCmd)
	favoritesCmd.AddCommand(favClearCmd)
} 