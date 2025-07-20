package favorites

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Favorite represents a favorite command
type Favorite struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	AddedAt     string `json:"added_at"`
}

// FavoritesManager manages user favorites
type FavoritesManager struct {
	filePath string
}

// NewFavoritesManager creates a new favorites manager
func NewFavoritesManager() (*FavoritesManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("홈 디렉토리를 찾을 수 없습니다: %v", err)
	}

	// Create .vi-assistant directory if it doesn't exist
	configDir := filepath.Join(homeDir, ".vi-assistant")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("설정 디렉토리를 생성할 수 없습니다: %v", err)
	}

	filePath := filepath.Join(configDir, "favorites.json")
	return &FavoritesManager{filePath: filePath}, nil
}

// Add adds a command to favorites
func (fm *FavoritesManager) Add(command, description string) error {
	favorites, err := fm.loadFavorites()
	if err != nil {
		return err
	}

	// Check if already exists
	for _, fav := range favorites {
		if fav.Command == command {
			return fmt.Errorf("이미 즐겨찾기에 추가된 명령어입니다: %s", command)
		}
	}

	// Add new favorite
	newFavorite := Favorite{
		Command:     command,
		Description: description,
		AddedAt:     getCurrentTime(),
	}

	favorites = append(favorites, newFavorite)

	return fm.saveFavorites(favorites)
}

// Remove removes a command from favorites
func (fm *FavoritesManager) Remove(command string) error {
	favorites, err := fm.loadFavorites()
	if err != nil {
		return err
	}

	var newFavorites []Favorite
	found := false

	for _, fav := range favorites {
		if fav.Command != command {
			newFavorites = append(newFavorites, fav)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("즐겨찾기에서 찾을 수 없는 명령어입니다: %s", command)
	}

	return fm.saveFavorites(newFavorites)
}

// List returns all favorites
func (fm *FavoritesManager) List() ([]Favorite, error) {
	return fm.loadFavorites()
}

// Clear removes all favorites
func (fm *FavoritesManager) Clear() error {
	return fm.saveFavorites([]Favorite{})
}

// loadFavorites loads favorites from file
func (fm *FavoritesManager) loadFavorites() ([]Favorite, error) {
	if _, err := os.Stat(fm.filePath); os.IsNotExist(err) {
		return []Favorite{}, nil
	}

	data, err := ioutil.ReadFile(fm.filePath)
	if err != nil {
		return nil, fmt.Errorf("즐겨찾기 파일을 읽을 수 없습니다: %v", err)
	}

	var favorites []Favorite
	if err := json.Unmarshal(data, &favorites); err != nil {
		return nil, fmt.Errorf("즐겨찾기 파일 파싱 오류: %v", err)
	}

	return favorites, nil
}

// saveFavorites saves favorites to file
func (fm *FavoritesManager) saveFavorites(favorites []Favorite) error {
	data, err := json.MarshalIndent(favorites, "", "  ")
	if err != nil {
		return fmt.Errorf("즐겨찾기 저장 오류: %v", err)
	}

	if err := ioutil.WriteFile(fm.filePath, data, 0644); err != nil {
		return fmt.Errorf("즐겨찾기 파일 쓰기 오류: %v", err)
	}

	return nil
}

// getCurrentTime returns current time in a readable format
func getCurrentTime() string {
	// Simple time format for now
	return "2024-01-01 12:00:00"
}

// FormatFavorites formats favorites for display
func FormatFavorites(favorites []Favorite, lang string) string {
	if len(favorites) == 0 {
		if lang == "en" {
			return "No favorite commands found.\n"
		}
		return "즐겨찾기된 명령어가 없습니다.\n"
	}

	var output strings.Builder
	
	if lang == "en" {
		output.WriteString(fmt.Sprintf("Favorite Commands (%d):\n\n", len(favorites)))
	} else {
		output.WriteString(fmt.Sprintf("즐겨찾기 명령어 (%d개):\n\n", len(favorites)))
	}

	for i, fav := range favorites {
		output.WriteString(fmt.Sprintf("%d. %s\n", i+1, fav.Command))
		if lang == "en" {
			output.WriteString(fmt.Sprintf("   Description: %s\n", fav.Description))
			output.WriteString(fmt.Sprintf("   Added: %s\n", fav.AddedAt))
		} else {
			output.WriteString(fmt.Sprintf("   설명: %s\n", fav.Description))
			output.WriteString(fmt.Sprintf("   추가일: %s\n", fav.AddedAt))
		}
		output.WriteString("\n")
	}

	return output.String()
} 