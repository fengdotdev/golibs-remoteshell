package settings

import (
	"fmt"

	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.DataTransferObject[SettingsDTO] = (*Settings)(nil)



// FromDTO implements trait.DataTransferObject.
func (s *Settings) FromDTO(dto SettingsDTO) error {
	if dto.Version != version {
		return fmt.Errorf("version mismatch: expected %s, got %s", version, dto.Version)
	}

	s.key = dto.Key
	s.user = dto.User
	s.password = dto.Password
	return nil
}

// ToDTO implements trait.DataTransferObject.
func (s *Settings) ToDTO() (SettingsDTO, error) {
	return SettingsDTO{
		Version:  version,
		Key:      s.key,
		User:     s.user,
		Password: s.password,
	}, nil
}
