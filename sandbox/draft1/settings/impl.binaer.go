package settings

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Binarer = (*Settings)(nil)

// FromBinary implements trait.Binarer.
func (s *Settings) FromBinary(data []byte) error {
	return s.FromGOB(data)
}

// ToBinary implements trait.Binarer.
func (s *Settings) ToBinary() ([]byte, error) {
	return s.ToGOB()
}
