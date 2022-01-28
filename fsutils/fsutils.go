package fsutils

import "os"

func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return !fi.Mode().IsDir()
}
