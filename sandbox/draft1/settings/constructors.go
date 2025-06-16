package settings

func New(key, user, password string) *Settings {
	return &Settings{
		key:      key,
		user:     user,
		password: password,
	}
}

func NewSettingsFromDTO(dto SettingsDTO) (*Settings, error) {
	s := &Settings{}
	if err := s.FromDTO(dto); err != nil {
		return nil, err
	}
	return s, nil
}

// DTO Contructors

func NewSettingsDTO(key, user, password string) SettingsDTO {
	return SettingsDTO{
		Version:  version,
		Key:      key,
		User:     user,
		Password: password,
	}
}

func NewDTOFromGOB(data []byte) (*SettingsDTO, error) {
	dto := &SettingsDTO{}
	if err := dto.FromGOB(data); err != nil {
		return nil, err // Return an error if decoding fails
	}
	return dto, nil // Return the decoded DTO
}

func NewDTOFromJSON(s string) (*SettingsDTO, error) {
	dto := &SettingsDTO{}
	if err := dto.FromJSON(s); err != nil {
		return nil, err // Return an error if JSON parsing fails
	}
	return dto, nil // Return the parsed DTO
}
