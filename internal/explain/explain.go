package explain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Command represents a vi command structure
type Command struct {
	Keyword     string `json:"keyword"`
	Command     string `json:"command"`
	Description string `json:"description"`
	Example     string `json:"example"`
	Category    string `json:"category"`
}

// ExplainResult represents explanation result
type ExplainResult struct {
	Command     Command
	Found       bool
	Suggestions []Command
}

// Explain explains a specific vi command
func Explain(command string) (*ExplainResult, error) {
	commands, err := loadCommands()
	if err != nil {
		return nil, fmt.Errorf("명령어 데이터를 로드할 수 없습니다: %v", err)
	}

	command = strings.TrimSpace(command)
	var result ExplainResult
	var suggestions []Command

	for _, cmd := range commands {
		if strings.EqualFold(cmd.Command, command) {
			result.Command = cmd
			result.Found = true
			break
		}
		
		// 유사한 명령어들을 제안으로 추가
		if strings.Contains(strings.ToLower(cmd.Command), strings.ToLower(command)) ||
			strings.Contains(strings.ToLower(cmd.Keyword), strings.ToLower(command)) {
			suggestions = append(suggestions, cmd)
		}
	}

	result.Suggestions = suggestions
	return &result, nil
}

// GetQuickReference returns a quick reference for common commands
func GetQuickReference(lang string) string {
	commonCommands := []struct {
		Command     string
		Description string
		Category    string
	}{
		{":w", "파일 저장", "file"},
		{":q", "종료", "file"},
		{":wq", "저장 후 종료", "file"},
		{":q!", "강제 종료", "file"},
		{"i", "삽입 모드", "mode"},
		{"Esc", "명령 모드", "mode"},
		{"yy", "줄 복사", "copy"},
		{"dd", "줄 삭제", "delete"},
		{"p", "붙여넣기", "paste"},
		{"u", "실행 취소", "edit"},
		{"h", "왼쪽 이동", "navigation"},
		{"j", "아래 이동", "navigation"},
		{"k", "위 이동", "navigation"},
		{"l", "오른쪽 이동", "navigation"},
		{"/pattern", "검색", "search"},
		{":s/old/new", "바꾸기", "edit"},
	}

	var output strings.Builder
	
	if lang == "en" {
		output.WriteString("Quick Reference - Common vi Commands:\n\n")
	} else {
		output.WriteString("빠른 참조 - 자주 사용하는 vi 명령어:\n\n")
	}

	output.WriteString("📁 File Operations:\n")
	for _, cmd := range commonCommands {
		if cmd.Category == "file" {
			output.WriteString(fmt.Sprintf("  %-10s %s\n", cmd.Command, cmd.Description))
		}
	}

	output.WriteString("\n🎯 Mode Switching:\n")
	for _, cmd := range commonCommands {
		if cmd.Category == "mode" {
			output.WriteString(fmt.Sprintf("  %-10s %s\n", cmd.Command, cmd.Description))
		}
	}

	output.WriteString("\n✂️ Edit Operations:\n")
	for _, cmd := range commonCommands {
		if cmd.Category == "copy" || cmd.Category == "delete" || cmd.Category == "paste" || cmd.Category == "edit" {
			output.WriteString(fmt.Sprintf("  %-10s %s\n", cmd.Command, cmd.Description))
		}
	}

	output.WriteString("\n🧭 Navigation:\n")
	for _, cmd := range commonCommands {
		if cmd.Category == "navigation" {
			output.WriteString(fmt.Sprintf("  %-10s %s\n", cmd.Command, cmd.Description))
		}
	}

	output.WriteString("\n🔍 Search & Replace:\n")
	for _, cmd := range commonCommands {
		if cmd.Category == "search" {
			output.WriteString(fmt.Sprintf("  %-10s %s\n", cmd.Command, cmd.Description))
		}
	}

	return output.String()
}

// GetCommandByCategory returns commands grouped by category
func GetCommandByCategory(category string) ([]Command, error) {
	commands, err := loadCommands()
	if err != nil {
		return nil, fmt.Errorf("명령어 데이터를 로드할 수 없습니다: %v", err)
	}

	var results []Command
	category = strings.ToLower(category)

	for _, cmd := range commands {
		if strings.ToLower(cmd.Category) == category {
			results = append(results, cmd)
		}
	}

	return results, nil
}

// loadCommands loads commands from the JSON file
func loadCommands() ([]Command, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Construct the path to commands.json
	dataPath := filepath.Join(wd, "data", "commands.json")

	// Read the JSON file
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("commands.json 파일을 읽을 수 없습니다: %v", err)
	}

	// Parse JSON
	var commands []Command
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("JSON 파싱 오류: %v", err)
	}

	return commands, nil
}

// FormatExplanation formats explanation result for display
func FormatExplanation(result *ExplainResult, lang string) string {
	var output strings.Builder

	if result.Found {
		if lang == "en" {
			output.WriteString(fmt.Sprintf("Command: %s\n", result.Command.Command))
			output.WriteString(fmt.Sprintf("Category: %s\n", result.Command.Category))
			output.WriteString(fmt.Sprintf("Description: %s\n", result.Command.Description))
			output.WriteString(fmt.Sprintf("Example: %s\n", result.Command.Example))
		} else {
			output.WriteString(fmt.Sprintf("명령어: %s\n", result.Command.Command))
			output.WriteString(fmt.Sprintf("카테고리: %s\n", result.Command.Category))
			output.WriteString(fmt.Sprintf("설명: %s\n", result.Command.Description))
			output.WriteString(fmt.Sprintf("예제: %s\n", result.Command.Example))
		}
	} else {
		if lang == "en" {
			output.WriteString("Command not found.\n\n")
			if len(result.Suggestions) > 0 {
				output.WriteString("Did you mean one of these?\n")
			}
		} else {
			output.WriteString("명령어를 찾을 수 없습니다.\n\n")
			if len(result.Suggestions) > 0 {
				output.WriteString("다음 중 하나를 찾으셨나요?\n")
			}
		}

		for i, suggestion := range result.Suggestions {
			if i >= 5 { // 최대 5개만 표시
				break
			}
			if lang == "en" {
				output.WriteString(fmt.Sprintf("  %s - %s\n", suggestion.Command, suggestion.Description))
			} else {
				output.WriteString(fmt.Sprintf("  %s - %s\n", suggestion.Command, suggestion.Description))
			}
		}
	}

	return output.String()
} 