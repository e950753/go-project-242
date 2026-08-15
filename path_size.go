package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, human, all, recursive bool) (string, error) {
	size, err := getSize(path, all, recursive)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

func getSize(path string, all, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range entries {
		if !recursive && entry.IsDir() {
			continue
		}

		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			dirSize, err := getSize(entryPath, all, recursive)
			if err != nil {
				return 0, err
			}

			total += dirSize
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		total += entryInfo.Size()
	}

	return total, nil
}

func formatSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0
	floatSize := float64(size)

	if !human {
		return fmt.Sprintf("%d%s", size, units[0])
	}

	for floatSize >= 1024 && unitIndex < len(units)-1 {
		floatSize /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%.0f%s", floatSize, units[unitIndex])
	}

	return fmt.Sprintf("%.1f%s", floatSize, units[unitIndex])
}
