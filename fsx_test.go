package fsx

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/satori/go.uuid"
)

func TestFile_Exists(t *testing.T) {
	file := File{
		Name: "/etc/hosts",
	}
	if err := file.Exists(); err != nil {
		t.Error("file doesn't exists")
	}
}

func TestFile_Create(t *testing.T) {
	file := File{
		Name: filepath.Join(os.TempDir(), uuid.NewV4().String()),
	}

	if err := file.Create(); err != nil {
		t.Errorf("failed to create file [%s], err: %v", file.Name, err)
	}

	// cleanup
	os.Remove(file.Name)
}

func TestSymlink_Exists(t *testing.T) {
	sl := Symlink{
		Src:    "/tmp/yolo",
		Target: "/tmp/yolo-target",
	}
	if err := sl.Exists(); err != nil {
		t.Error("symlink doesn't exits, err:", err)
	}
}
