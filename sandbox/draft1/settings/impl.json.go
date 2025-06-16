package settings

import (
	"fmt"

	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.JSONSerializer = (*Settings)(nil)

// FromJSON implements trait.JSONSerializer.
func (s *Settings) FromJSON(data string) error {

	dto, err := NewDTOFromJSON(data)
	if err != nil {
		return err // Return an error if JSON parsing fails
	}

	if dto.Version != version {
		return fmt.Errorf("version mismatch: expected %s, got %s", version, dto.Version)
	}

	s.key = dto.Key
	s.user = dto.User
	s.password = dto.Password
	return nil

}

// ToJSON implements trait.JSONSerializer.
func (s *Settings) ToJSON() (string, error) {
	dto, err := s.ToDTO()
	if err != nil {
		return "", err // Return an error if DTO conversion fails
	}

	data, err := dto.ToJSON()
	if err != nil {
		return "", err // Return an error if JSON encoding fails
	}
	return data, nil
}
