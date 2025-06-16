package settings

type SettingsDTO struct {
	Version  string `json:"version"` // Version of the settings, e.g., "draft1"
	Key      string `json:"key"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (dto *SettingsDTO) String() string {
	return "SettingsDTO{" +
		"Version: " + dto.Version + ", " +
		"Key: " + dto.Key + ", " +
		"User: " + dto.User + ", " +
		"Password: " + dto.Password +
		"}"
}
