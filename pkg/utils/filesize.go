package utils

import "os"

// GetFileSize returns the size in bytes of the given path
func GetFileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}
