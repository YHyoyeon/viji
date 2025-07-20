// search 패키지는 vi 명령어 검색 기능을 제공합니다
// JSON 파일에서 명령어 데이터를 로드하고 키워드 기반 검색을 수행합니다
package search

import (
	"encoding/json"  // JSON 데이터 처리를 위한 패키지
	"fmt"            // 표준 출력/입력 포맷팅을 위한 패키지
	"io/ioutil"      // 파일 읽기/쓰기를 위한 패키지
	"os"             // 운영체제 인터페이스를 위한 패키지
	"path/filepath"  // 파일 경로 조작을 위한 패키지
	"strings"        // 문자열 조작을 위한 패키지
)

// Command 구조체는 vi 명령어의 정보를 담는 데이터 구조입니다
// JSON 파일에서 로드되는 각 명령어의 구조를 정의합니다
type Command struct {
	Keyword     string `json:"keyword"`     // 검색 키워드 (복수 가능)
	Command     string `json:"command"`     // 실제 vi 명령어 (예: :wq, yy)
	Description string `json:"description"` // 명령어에 대한 설명
	Example     string `json:"example"`     // 사용 예제
	Category    string `json:"category"`    // 명령어 카테고리 (file, edit, navigation 등)
}

// SearchResult 구조체는 검색 결과를 담는 데이터 구조입니다
// 검색된 명령어 목록과 개수 정보를 포함합니다
type SearchResult struct {
	Commands []Command  // 검색된 명령어들의 슬라이스
	Count    int        // 검색된 명령어의 총 개수
}

// Search 함수는 키워드를 사용하여 vi 명령어를 검색합니다
// 키워드와 일치하는 모든 명령어를 찾아서 결과를 반환합니다
func Search(keyword string) (*SearchResult, error) {
	// JSON 파일에서 모든 명령어 데이터를 로드합니다
	commands, err := loadCommands()
	if err != nil {
		return nil, fmt.Errorf("명령어 데이터를 로드할 수 없습니다: %v", err)
	}

	var results []Command  // 검색 결과를 저장할 슬라이스
	keyword = strings.ToLower(keyword)  // 키워드를 소문자로 변환하여 대소문자 구분 없이 검색

	// 모든 명령어를 순회하면서 키워드와 일치하는 항목을 찾습니다
	for _, cmd := range commands {
		// 다음 항목들에서 키워드 검색:
		// - 명령어 키워드
		// - 실제 명령어
		// - 설명
		// - 카테고리
		if strings.Contains(strings.ToLower(cmd.Keyword), keyword) ||
			strings.Contains(strings.ToLower(cmd.Command), keyword) ||
			strings.Contains(strings.ToLower(cmd.Description), keyword) ||
			strings.Contains(strings.ToLower(cmd.Category), keyword) {
			results = append(results, cmd)  // 일치하는 명령어를 결과에 추가
		}
	}

	// 검색 결과를 SearchResult 구조체로 반환합니다
	return &SearchResult{
		Commands: results,  // 검색된 명령어들
		Count:    len(results),  // 검색된 명령어의 개수
	}, nil
}

// SearchByCategory 함수는 카테고리별로 vi 명령어를 검색합니다
// 특정 카테고리에 속하는 모든 명령어를 찾아서 결과를 반환합니다
func SearchByCategory(category string) (*SearchResult, error) {
	// JSON 파일에서 모든 명령어 데이터를 로드합니다
	commands, err := loadCommands()
	if err != nil {
		return nil, fmt.Errorf("명령어 데이터를 로드할 수 없습니다: %v", err)
	}

	var results []Command  // 검색 결과를 저장할 슬라이스
	category = strings.ToLower(category)  // 카테고리를 소문자로 변환하여 대소문자 구분 없이 검색

	// 모든 명령어를 순회하면서 지정된 카테고리와 일치하는 항목을 찾습니다
	for _, cmd := range commands {
		if strings.ToLower(cmd.Category) == category {
			results = append(results, cmd)  // 일치하는 명령어를 결과에 추가
		}
	}

	// 검색 결과를 SearchResult 구조체로 반환합니다
	return &SearchResult{
		Commands: results,  // 검색된 명령어들
		Count:    len(results),  // 검색된 명령어의 개수
	}, nil
}

// GetCategories 함수는 사용 가능한 모든 카테고리를 반환합니다
// 중복을 제거하여 고유한 카테고리 목록을 제공합니다
func GetCategories() ([]string, error) {
	// JSON 파일에서 모든 명령어 데이터를 로드합니다
	commands, err := loadCommands()
	if err != nil {
		return nil, fmt.Errorf("명령어 데이터를 로드할 수 없습니다: %v", err)
	}

	// 중복을 제거하기 위해 맵을 사용합니다
	categoryMap := make(map[string]bool)
	for _, cmd := range commands {
		categoryMap[cmd.Category] = true  // 각 카테고리를 맵에 추가 (중복 자동 제거)
	}

	// 맵의 키들을 슬라이스로 변환합니다
	var categories []string
	for category := range categoryMap {
		categories = append(categories, category)
	}

	return categories, nil
}

// loadCommands 함수는 JSON 파일에서 vi 명령어 데이터를 로드합니다
// 내부적으로 사용되는 헬퍼 함수로, 모든 검색 함수에서 공통으로 사용됩니다
func loadCommands() ([]Command, error) {
	// 현재 작업 디렉토리를 가져옵니다
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// commands.json 파일의 경로를 구성합니다
	dataPath := filepath.Join(wd, "data", "commands.json")

	// JSON 파일을 읽어들입니다
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("commands.json 파일을 읽을 수 없습니다: %v", err)
	}

	// JSON 데이터를 Command 구조체 슬라이스로 파싱합니다
	var commands []Command
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("JSON 파싱 오류: %v", err)
	}

	return commands, nil
}

// FormatSearchResults 함수는 검색 결과를 사용자에게 보여주기 위한 형태로 포맷팅합니다
// 언어 설정에 따라 한국어 또는 영어로 결과를 표시합니다
func FormatSearchResults(results *SearchResult, lang string) string {
	// 검색 결과가 없는 경우 처리
	if results.Count == 0 {
		if lang == "en" {
			return "No commands found matching your search criteria."
		}
		return "검색 조건에 맞는 명령어를 찾을 수 없습니다."
	}

	// 결과를 효율적으로 구성하기 위해 strings.Builder를 사용합니다
	var output strings.Builder
	
	// 검색 결과 개수를 표시합니다
	if lang == "en" {
		output.WriteString(fmt.Sprintf("Found %d command(s):\n\n", results.Count))
	} else {
		output.WriteString(fmt.Sprintf("%d개의 명령어를 찾았습니다:\n\n", results.Count))
	}

	// 각 검색 결과를 순회하면서 포맷팅합니다
	for i, cmd := range results.Commands {
		output.WriteString(fmt.Sprintf("%d. %s\n", i+1, cmd.Command))  // 명령어 번호와 실제 명령어
		if lang == "en" {  // 영어 버전
			output.WriteString(fmt.Sprintf("   Category: %s\n", cmd.Category))
			output.WriteString(fmt.Sprintf("   Description: %s\n", cmd.Description))
			output.WriteString(fmt.Sprintf("   Example: %s\n", cmd.Example))
		} else {  // 한국어 버전
			output.WriteString(fmt.Sprintf("   카테고리: %s\n", cmd.Category))
			output.WriteString(fmt.Sprintf("   설명: %s\n", cmd.Description))
			output.WriteString(fmt.Sprintf("   예제: %s\n", cmd.Example))
		}
		output.WriteString("\n")  // 각 명령어 사이에 빈 줄 추가
	}

	return output.String()  // 포맷팅된 문자열 반환
} 