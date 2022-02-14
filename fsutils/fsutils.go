package fsutils

import "os"

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return true
}

func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return !fi.Mode().IsDir()
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return fi.Mode().IsDir()
}
