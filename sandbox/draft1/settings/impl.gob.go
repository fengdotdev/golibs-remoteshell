package settings

import "fmt"

type GOB interface {
	FromGOB(data []byte) error
	ToGOB() ([]byte, error)
}

var _ GOB = (*Settings)(nil)

func (s *Settings) FromGOB(data []byte) error {
	dto, err := NewDTOFromGOB(data)
	if err != nil {
		return fmt.Errorf("failed to read GOB data: %w", err)
	}

	if dto.Version != version {
		return fmt.Errorf("version mismatch: expected %s, got %s", version, dto.Version)
	}

	s.key = dto.Key
	s.user = dto.User
	s.password = dto.Password

	return nil
}

func (s *Settings) ToGOB() ([]byte, error) {

	dto, err := s.ToDTO()
	if err != nil {
		return nil, err
	}

	data, err := dto.ToGOB()
	if err != nil {
		return nil, err
	}
	return data, nil

}
