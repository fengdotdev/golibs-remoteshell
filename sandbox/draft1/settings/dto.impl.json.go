package settings

import "encoding/json"

func (dto *SettingsDTO) ToJSON() (string, error) {
	// Convert the DTO to JSON format
	jsonData, err := json.Marshal(dto)
	if err != nil {
		return "", err // Return an error if JSON marshaling fails
	}
	return string(jsonData), nil // Return the JSON string
}
func (dto *SettingsDTO) FromJSON(s string) error {
	// Parse the JSON string into the DTO
	if err := json.Unmarshal([]byte(s), dto); err != nil {
		return err // Return an error if JSON unmarshaling fails
	}
	return nil // Return nil if unmarshaling is successful
}
