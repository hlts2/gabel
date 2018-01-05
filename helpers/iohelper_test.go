package helpers

import (
	"os"
	"testing"
)

//Temporary Settings
const (
	TmpDirName  = "gabel_tmp"
	TmpFileName = "gabel_tmp.txt"
)

func TestMkdir(t *testing.T) {
	if err := Mkdir(TmpDirName); err != nil {
		t.Errorf("Mkdir() is error: %v", err)
	}

	_, err := os.Stat(TmpDirName)
	if os.IsNotExist(err) {
		t.Errorf("%s generations faild", TmpDirName)
	}

	os.Remove(TmpDirName)
}

func TestCreateFile(t *testing.T) {
	f, err := CreateFile(TmpFileName, os.O_WRONLY)
	if err != nil {
		t.Errorf("CreateFile() is error: %v", err)
	}

	if f == nil {
		t.Error("CreateFile() file = nil")
	}

	os.Remove(TmpFileName)
}

func TestOpenFile(t *testing.T) {
	_, _ = CreateFile(TmpFileName, os.O_RDWR)
	f, err := OpenFile(TmpFileName, os.O_RDONLY)
	if err != nil {
		t.Errorf("OpenFile() is error: %v", err)
	}

	if f == nil {
		t.Error("OpenFile() file = nil")
	}

	os.Remove(TmpFileName)
}

func TestIsExist(t *testing.T) {
	_, _ = CreateFile(TmpFileName, os.O_RDWR)
	_, err := os.Stat(TmpFileName)
	if os.IsNotExist(err) {
		t.Errorf("%s does not exist", TmpFileName)
	}

	os.Remove(TmpFileName)
}
