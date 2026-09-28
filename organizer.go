package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var categories = map[string][]string{
	"Images": {
		".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg", ".bmp",
	},
	"Videos": {
		".mp4", ".mkv", ".avi", ".mov", ".webm", ".flv",
	},
	"Audio": {
		".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a",
	},
	"Documents": {
		".pdf", ".doc", ".docx", ".txt", ".rtf", ".odt",
	},
	"Spreadsheets": {
		".xls", ".xlsx", ".csv", ".ods",
	},
	"Presentations": {
		".ppt", ".pptx", ".odp",
	},
	"Archives": {
		".zip", ".rar", ".7z", ".tar", ".gz", ".bz2",
	},
	"Code": {
		".go", ".js", ".ts", ".py", ".java", ".c", ".cpp",
		".html", ".css", ".php", ".json", ".xml", ".sql",
	},
}

func getCategory(extension string) string {
	extension = strings.ToLower(extension)

	for category, extensions := range categories {
		for _, ext := range extensions {
			if extension == ext {
				return category
			}
		}
	}

	return "Others"
}

func uniquePath(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path
	}

	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	name := strings.TrimSuffix(filepath.Base(path), ext)

	for i := 1; ; i++ {
		newName := fmt.Sprintf("%s_%d%s", name, i, ext)
		newPath := filepath.Join(dir, newName)

		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
	}
}

func Organize(directory string, dryRun bool) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(directory, entry.Name())
		extension := filepath.Ext(entry.Name())

		category := getCategory(extension)
		categoryDir := filepath.Join(directory, category)
		destination := filepath.Join(categoryDir, entry.Name())

		if dryRun {
			fmt.Printf("[PREVIEW] %s -> %s\n", entry.Name(), category)
			continue
		}

		if err := os.MkdirAll(categoryDir, 0755); err != nil {
			return err
		}

		destination = uniquePath(destination)

		if err := os.Rename(filePath, destination); err != nil {
			return fmt.Errorf("moving %s: %w", entry.Name(), err)
		}

		fmt.Printf("Moved: %s -> %s\n",
			entry.Name(),
			filepath.Base(categoryDir),
		)
	}

	return nil
}
