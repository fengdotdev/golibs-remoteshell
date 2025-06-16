package settings

import (
	"bytes"
	"encoding/gob"
)

func (dto *SettingsDTO) ToGOB() ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer) // Create a new GOB encoder

	if err := encoder.Encode(dto); err != nil {
		return nil, err // Return an error if encoding fails
	}

	return buffer.Bytes(), nil // Return the encoded data as a byte slice
}

func (dto *SettingsDTO) FromGOB(data []byte) error {
	var buffer bytes.Buffer
	buffer.Write(data) // Write the byte slice to the buffer

	decoder := gob.NewDecoder(&buffer) // Create a new GOB decoder

	if err := decoder.Decode(dto); err != nil {
		return err // Return an error if decoding fails
	}

	return nil // Return nil if decoding is successful
}
