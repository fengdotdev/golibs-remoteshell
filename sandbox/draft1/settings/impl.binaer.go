package settings

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Binarer = (*Settings)(nil)

// FromBinary implements trait.Binarer.
func (s *Settings) FromBinary([]byte) error {
	panic("unimplemented")
}

// ToBinary implements trait.Binarer.
func (s *Settings) ToBinary() ([]byte, error) {
	panic("unimplemented")
}
