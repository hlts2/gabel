package utils

import "os"

//Mkdir generate directory of given path
func Mkdir(path string) error {
	if ok := isExist(path); ok {
		return nil
	}
	return os.Mkdir(path, os.ModePerm)
}

//CreateFile generate file of given path
func CreateFile(name string, flag int) (*os.File, error) {
	if ok := isExist(name); ok {
		file, err := OpenFile(name, flag)
		file.Truncate(0)
		return file, err
	}

	return OpenFile(name, flag)
}

//OpenFile opens file
func OpenFile(name string, flag int) (*os.File, error) {
	return os.OpenFile(name, flag, os.ModePerm)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
