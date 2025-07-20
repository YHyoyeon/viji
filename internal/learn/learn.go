package learn

import (
	"fmt"
	"strings"
)

// Lesson represents a learning lesson
type Lesson struct {
	Title       string
	Description string
	Commands    []LessonCommand
	Tips        []string
}

// LessonCommand represents a command in a lesson
type LessonCommand struct {
	Command     string
	Description string
	Example     string
	Practice    string
}

// GetBeginnerLessons returns beginner level lessons
func GetBeginnerLessons(lang string) []Lesson {
	if lang == "en" {
		return getBeginnerLessonsEN()
	}
	return getBeginnerLessonsKO()
}

// GetIntermediateLessons returns intermediate level lessons
func GetIntermediateLessons(lang string) []Lesson {
	if lang == "en" {
		return getIntermediateLessonsEN()
	}
	return getIntermediateLessonsKO()
}

// getBeginnerLessonsKO returns Korean beginner lessons
func getBeginnerLessonsKO() []Lesson {
	return []Lesson{
		{
			Title:       "1. vi 시작하기 - 기본 모드 이해",
			Description: "vi의 두 가지 주요 모드와 기본 이동 명령어를 배워봅시다.",
			Commands: []LessonCommand{
				{
					Command:     "vi filename",
					Description: "vi 에디터로 파일 열기",
					Example:     "vi test.txt",
					Practice:    "터미널에서 'vi test.txt'를 입력해보세요",
				},
				{
					Command:     "i",
					Description: "삽입 모드로 전환 (텍스트 입력 가능)",
					Example:     "i를 누르면 하단에 -- INSERT -- 표시",
					Practice:    "i를 누르고 텍스트를 입력해보세요",
				},
				{
					Command:     "Esc",
					Description: "명령 모드로 돌아가기",
					Example:     "Esc를 누르면 -- INSERT -- 표시가 사라짐",
					Practice:    "텍스트 입력 후 Esc를 눌러 명령 모드로 전환",
				},
				{
					Command:     "h, j, k, l",
					Description: "기본 커서 이동 (왼쪽, 아래, 위, 오른쪽)",
					Example:     "h: 왼쪽, j: 아래, k: 위, l: 오른쪽",
					Practice:    "명령 모드에서 h, j, k, l로 커서를 움직여보세요",
				},
			},
			Tips: []string{
				"💡 vi는 항상 명령 모드에서 시작합니다",
				"💡 텍스트를 입력하려면 반드시 'i'로 삽입 모드로 전환해야 합니다",
				"💡 명령 모드에서는 모든 키가 명령어로 인식됩니다",
				"💡 h, j, k, l은 키보드의 왼쪽에 있어서 한 손으로 조작하기 편합니다",
			},
		},
		{
			Title:       "2. 파일 저장과 종료",
			Description: "작업한 내용을 저장하고 vi를 종료하는 방법을 배워봅시다.",
			Commands: []LessonCommand{
				{
					Command:     ":w",
					Description: "현재 파일 저장",
					Example:     ":w를 입력하고 Enter",
					Practice:    "텍스트를 입력한 후 :w로 저장해보세요",
				},
				{
					Command:     ":q",
					Description: "변경사항 없이 종료",
					Example:     ":q를 입력하고 Enter",
					Practice:    "저장 후 :q로 종료해보세요",
				},
				{
					Command:     ":wq",
					Description: "저장하고 종료 (가장 많이 사용)",
					Example:     ":wq를 입력하고 Enter",
					Practice:    "작업 완료 후 :wq로 저장하고 종료",
				},
				{
					Command:     ":q!",
					Description: "변경사항 무시하고 강제 종료",
					Example:     ":q!를 입력하고 Enter",
					Practice:    "실수로 변경한 경우 :q!로 강제 종료",
				},
			},
			Tips: []string{
				"💡 :wq는 'write and quit'의 줄임말입니다",
				"💡 :q!는 변경사항을 저장하지 않고 나가는 긴급 탈출 명령어입니다",
				"💡 저장하지 않고 종료하려고 하면 vi가 경고를 표시합니다",
				"💡 :w filename으로 다른 이름으로 저장할 수 있습니다",
			},
		},
		{
			Title:       "3. 텍스트 편집 기본",
			Description: "텍스트를 삭제하고 복사하는 기본적인 편집 명령어를 배워봅시다.",
			Commands: []LessonCommand{
				{
					Command:     "x",
					Description: "커서 위치의 문자 삭제",
					Example:     "x를 누르면 커서 위치의 문자가 삭제됨",
					Practice:    "텍스트에서 x를 눌러 문자를 삭제해보세요",
				},
				{
					Command:     "dd",
					Description: "현재 줄 전체 삭제",
					Example:     "dd를 누르면 현재 줄이 삭제됨",
					Practice:    "dd를 눌러 현재 줄을 삭제해보세요",
				},
				{
					Command:     "yy",
					Description: "현재 줄 복사 (야크)",
					Example:     "yy를 누르면 현재 줄이 복사됨",
					Practice:    "yy를 눌러 줄을 복사해보세요",
				},
				{
					Command:     "p",
					Description: "복사한 내용을 다음 위치에 붙여넣기",
					Example:     "yy 후 p를 누르면 다음 줄에 붙여넣어짐",
					Practice:    "yy로 복사한 후 p로 붙여넣어보세요",
				},
			},
			Tips: []string{
				"💡 dd는 'delete line'의 줄임말입니다",
				"💡 yy는 'yank'의 줄임말로, 복사 기능입니다",
				"💡 p는 'paste'의 줄임말입니다",
				"💡 P(대문자)를 누르면 이전 위치에 붙여넣어집니다",
			},
		},
	}
}

// getBeginnerLessonsEN returns English beginner lessons
func getBeginnerLessonsEN() []Lesson {
	return []Lesson{
		{
			Title:       "1. Getting Started with vi - Understanding Basic Modes",
			Description: "Learn about vi's two main modes and basic movement commands.",
			Commands: []LessonCommand{
				{
					Command:     "vi filename",
					Description: "Open file with vi editor",
					Example:     "vi test.txt",
					Practice:    "Type 'vi test.txt' in terminal",
				},
				{
					Command:     "i",
					Description: "Switch to insert mode (can input text)",
					Example:     "Press i to see -- INSERT -- at bottom",
					Practice:    "Press i and type some text",
				},
				{
					Command:     "Esc",
					Description: "Return to command mode",
					Example:     "Press Esc to remove -- INSERT --",
					Practice:    "After typing text, press Esc to switch to command mode",
				},
				{
					Command:     "h, j, k, l",
					Description: "Basic cursor movement (left, down, up, right)",
					Example:     "h: left, j: down, k: up, l: right",
					Practice:    "In command mode, move cursor with h, j, k, l",
				},
			},
			Tips: []string{
				"💡 vi always starts in command mode",
				"💡 You must press 'i' to switch to insert mode to type text",
				"💡 In command mode, every key is treated as a command",
				"💡 h, j, k, l are on the left side of keyboard for easy one-hand operation",
			},
		},
		{
			Title:       "2. Saving Files and Exiting",
			Description: "Learn how to save your work and exit vi.",
			Commands: []LessonCommand{
				{
					Command:     ":w",
					Description: "Save current file",
					Example:     "Type :w and press Enter",
					Practice:    "After typing text, save with :w",
				},
				{
					Command:     ":q",
					Description: "Exit without changes",
					Example:     "Type :q and press Enter",
					Practice:    "After saving, exit with :q",
				},
				{
					Command:     ":wq",
					Description: "Save and exit (most commonly used)",
					Example:     "Type :wq and press Enter",
					Practice:    "After completing work, save and exit with :wq",
				},
				{
					Command:     ":q!",
					Description: "Force exit ignoring changes",
					Example:     "Type :q! and press Enter",
					Practice:    "If you made mistakes, force exit with :q!",
				},
			},
			Tips: []string{
				"💡 :wq stands for 'write and quit'",
				"💡 :q! is an emergency escape command that exits without saving",
				"💡 vi will warn you if you try to exit without saving",
				"💡 You can save with different name using :w filename",
			},
		},
	}
}

// getIntermediateLessonsKO returns Korean intermediate lessons
func getIntermediateLessonsKO() []Lesson {
	return []Lesson{
		{
			Title:       "1. 고급 이동 명령어",
			Description: "더 효율적인 텍스트 탐색을 위한 고급 이동 명령어를 배워봅시다.",
			Commands: []LessonCommand{
				{
					Command:     "w",
					Description: "다음 단어의 시작으로 이동",
					Example:     "w를 누르면 다음 단어로 이동",
					Practice:    "w를 눌러 단어 단위로 이동해보세요",
				},
				{
					Command:     "b",
					Description: "이전 단어의 시작으로 이동",
					Example:     "b를 누르면 이전 단어로 이동",
					Practice:    "b를 눌러 뒤로 단어 단위 이동",
				},
				{
					Command:     "0",
					Description: "현재 줄의 시작으로 이동",
					Example:     "0을 누르면 줄의 맨 앞으로 이동",
					Practice:    "0을 눌러 줄의 시작으로 이동",
				},
				{
					Command:     "$",
					Description: "현재 줄의 끝으로 이동",
					Example:     "$를 누르면 줄의 맨 뒤로 이동",
					Practice:    "$를 눌러 줄의 끝으로 이동",
				},
				{
					Command:     "gg",
					Description: "파일의 첫 번째 줄로 이동",
					Example:     "gg를 누르면 파일 맨 위로 이동",
					Practice:    "gg를 눌러 파일의 시작으로 이동",
				},
				{
					Command:     "G",
					Description: "파일의 마지막 줄로 이동",
					Example:     "G를 누르면 파일 맨 아래로 이동",
					Practice:    "G를 눌러 파일의 끝으로 이동",
				},
			},
			Tips: []string{
				"💡 w는 'word'의 줄임말입니다",
				"💡 b는 'back'의 줄임말입니다",
				"💡 0은 숫자 0이지만 줄의 시작을 의미합니다",
				"💡 $는 줄의 끝을 의미하는 기호입니다",
				"💡 gg는 'go to beginning'의 줄임말입니다",
				"💡 G는 'go to end'의 줄임말입니다",
			},
		},
		{
			Title:       "2. 검색과 바꾸기",
			Description: "텍스트 내에서 특정 패턴을 찾고 바꾸는 방법을 배워봅시다.",
			Commands: []LessonCommand{
				{
					Command:     "/pattern",
					Description: "앞으로 패턴 검색",
					Example:     "/hello를 입력하면 'hello'를 앞으로 검색",
					Practice:    "/를 누르고 검색할 단어를 입력해보세요",
				},
				{
					Command:     "?pattern",
					Description: "뒤로 패턴 검색",
					Example:     "?hello를 입력하면 'hello'를 뒤로 검색",
					Practice:    "?를 누르고 검색할 단어를 입력해보세요",
				},
				{
					Command:     "n",
					Description: "다음 검색 결과로 이동",
					Example:     "n을 누르면 다음 검색 결과로 이동",
					Practice:    "검색 후 n을 눌러 다음 결과로 이동",
				},
				{
					Command:     "N",
					Description: "이전 검색 결과로 이동",
					Example:     "N을 누르면 이전 검색 결과로 이동",
					Practice:    "검색 후 N을 눌러 이전 결과로 이동",
				},
				{
					Command:     ":s/old/new",
					Description: "현재 줄의 첫 번째 'old'를 'new'로 바꾸기",
					Example:     ":s/cat/dog를 입력하면 첫 번째 'cat'이 'dog'로 바뀜",
					Practice:    ":s/를 사용해 현재 줄의 텍스트를 바꿔보세요",
				},
				{
					Command:     ":%s/old/new/g",
					Description: "파일 전체의 모든 'old'를 'new'로 바꾸기",
					Example:     ":%s/cat/dog/g를 입력하면 모든 'cat'이 'dog'로 바뀜",
					Practice:    ":%s/를 사용해 파일 전체의 텍스트를 바꿔보세요",
				},
			},
			Tips: []string{
				"💡 /는 앞으로, ?는 뒤로 검색합니다",
				"💡 n은 'next'의 줄임말입니다",
				"💡 N은 'previous'의 줄임말입니다",
				"💡 :s는 'substitute'의 줄임말입니다",
				"💡 g는 'global'의 줄임말로 모든 매치를 바꿉니다",
				"💡 %는 파일 전체를 의미합니다",
			},
		},
	}
}

// getIntermediateLessonsEN returns English intermediate lessons
func getIntermediateLessonsEN() []Lesson {
	return []Lesson{
		{
			Title:       "1. Advanced Movement Commands",
			Description: "Learn advanced movement commands for more efficient text navigation.",
			Commands: []LessonCommand{
				{
					Command:     "w",
					Description: "Move to beginning of next word",
					Example:     "Press w to move to next word",
					Practice:    "Press w to move word by word",
				},
				{
					Command:     "b",
					Description: "Move to beginning of previous word",
					Example:     "Press b to move to previous word",
					Practice:    "Press b to move backward word by word",
				},
				{
					Command:     "0",
					Description: "Move to beginning of current line",
					Example:     "Press 0 to move to start of line",
					Practice:    "Press 0 to move to beginning of line",
				},
				{
					Command:     "$",
					Description: "Move to end of current line",
					Example:     "Press $ to move to end of line",
					Practice:    "Press $ to move to end of line",
				},
			},
			Tips: []string{
				"💡 w stands for 'word'",
				"💡 b stands for 'back'",
				"💡 0 is the number zero but means start of line",
				"💡 $ means end of line",
			},
		},
	}
}

// FormatLesson formats a lesson for display
func FormatLesson(lesson Lesson, lessonNumber int, lang string) string {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("\n📚 %s\n", lesson.Title))
	output.WriteString(fmt.Sprintf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"))
	output.WriteString(fmt.Sprintf("%s\n\n", lesson.Description))

	output.WriteString("🎯 Commands to Learn:\n")
	for i, cmd := range lesson.Commands {
		output.WriteString(fmt.Sprintf("\n%d. %s\n", i+1, cmd.Command))
		output.WriteString(fmt.Sprintf("   Description: %s\n", cmd.Description))
		output.WriteString(fmt.Sprintf("   Example: %s\n", cmd.Example))
		output.WriteString(fmt.Sprintf("   Practice: %s\n", cmd.Practice))
	}

	output.WriteString("\n💡 Tips:\n")
	for _, tip := range lesson.Tips {
		output.WriteString(fmt.Sprintf("   %s\n", tip))
	}

	return output.String()
}

// FormatLessonList formats lesson list for display
func FormatLessonList(lessons []Lesson, level string, lang string) string {
	var output strings.Builder

	if lang == "en" {
		output.WriteString(fmt.Sprintf("\n🎓 %s Level Lessons\n", strings.Title(level)))
	} else {
		output.WriteString(fmt.Sprintf("\n🎓 %s 레벨 강의\n", level))
	}
	output.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	for i, lesson := range lessons {
		if lang == "en" {
			output.WriteString(fmt.Sprintf("%d. %s\n", i+1, lesson.Title))
		} else {
			output.WriteString(fmt.Sprintf("%d. %s\n", i+1, lesson.Title))
		}
	}

	return output.String()
} 