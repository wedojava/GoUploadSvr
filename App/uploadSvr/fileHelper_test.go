package main

import (
	"testing"
)

func TestGetFileList(t *testing.T) {
	//want := make([]string, 2)
	//want[0] = "..\\files\\gONOfI\\go_test.ipynb"
	//want[1] = "..\\files\\milDs8\\go_test.ipynb"
	GetFileList()
	//if !reflect.DeepEqual(want, got) {
	//	t.Errorf("want %s got %s", want, got)
	//}
}

func TestSaveFileLstInfo(t *testing.T) {
	files := GetFileList()
	SaveFileLstInfo(files, "../db.json")
}
