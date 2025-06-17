package settings

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fengdotdev/golibs-traits/trait"
)

const (
	name = ".remoteshell"
)

var _ trait.Saver = (*Settings)(nil)

// Load implements trait.Saver.
func (s *Settings) Load() error {
	return s.LoadFrom(name)
}

// LoadFrom implements trait.Saver.
func (s *Settings) LoadFrom(fullpath string) error {
	// check if file exists
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", fullpath)
	}
	// open file
	file, err := os.Open(fullpath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", fullpath, err)
	}
	defer file.Close()
	// read data from file
	data := make([]byte, 0)
	if _, err := file.Read(data); err != nil {
		return fmt.Errorf("failed to read data from file %s: %w", fullpath, err)
	}
	return s.FromGOB(data)
}

// Save implements trait.Saver.
func (s *Settings) Save() error {
	return s.SaveTo(name)
}

// SaveTo implements trait.Saver.
func (s *Settings) SaveTo(fullpath string) error {

	dirpath := filepath.Dir(fullpath)
	// create folders if not exists
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return err // return error if folders cannot be created
	}

	// create file if not exists
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := s.ToGOB()
	if err != nil {
		return err

	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
