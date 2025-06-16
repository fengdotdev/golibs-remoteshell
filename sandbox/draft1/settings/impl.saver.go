package settings

import (
	"os"

	"github.com/fengdotdev/golibs-traits/trait"
)
const (
	name = ".remoteshell"
)


var _ trait.Saver = (*Settings)(nil)

// Load implements trait.Saver.
func (s *Settings) Load() error {
	panic("unimplemented")
}

// LoadFrom implements trait.Saver.
func (s *Settings) LoadFrom(path string) error {
	panic("unimplemented")
}

// Save implements trait.Saver.
func (s *Settings) Save() error {
	panic("unimplemented")
}

// SaveTo implements trait.Saver.
func (s *Settings) SaveTo(path string) error {

	// create folders if not exists
	os.MkdirAll(path, os.ModePerm) // ensure the directory exists}

	// create file if not exists
	file, err := os.Create(path + "/" + name)
	if err != nil {
		

}
