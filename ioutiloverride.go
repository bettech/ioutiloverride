package ioutiloverride

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Iouo provides an interface so we do not access ioutil directly
type Iouo interface {
	ReadAll(r io.Reader) ([]byte, error)
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type realIoutil struct{}

func (rI *realIoutil) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func (rI *realIoutil) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (rI *realIoutil) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

type fakeIoutil struct{}

func (fI *fakeIoutil) ReadAll(r io.Reader) ([]byte, error) {
	return []byte(""), nil
}

func (fI *fakeIoutil) ReadFile(filename string) ([]byte, error) {
	return []byte(""), nil
}

func (fI *fakeIoutil) WriteFile(filename string, data []byte, perm os.FileMode) error {
	fmt.Println("\n\nFake WriteFile Called\n\n")
	return nil
}

// NewRealIoutil returns a realIoutil which delegates calls to the actual ioutil
// package; it should be used by packages in production.
func NewRealIoutil() Iouo {
	return &realIoutil{}
}

// NewFakeIoutil returns a fakeIoutil which does not access the disk and returns
// empty values; can be used for testing.
func NewFakeIoutil() Iouo {
	return &fakeIoutil{}
}
