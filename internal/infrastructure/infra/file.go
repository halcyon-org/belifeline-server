package infra

import (
	"os"
)

type FilesInterface interface {
	GetWD() (string, error)
	MkdirAll(path string, perm os.FileMode) error
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte, perm os.FileMode) error
}

type FilesInterfaceImpl struct{}

func NewFilesInterface() FilesInterface {
	return &FilesInterfaceImpl{}
}

func (f *FilesInterfaceImpl) GetWD() (string, error) {
	return os.Getwd()
}

func (f *FilesInterfaceImpl) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (f *FilesInterfaceImpl) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *FilesInterfaceImpl) WriteFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}
