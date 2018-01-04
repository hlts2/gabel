package utils

import "os"

//Mkdir generate directory of given path
func Mkdir(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	return os.Mkdir(path, os.ModePerm)
}
