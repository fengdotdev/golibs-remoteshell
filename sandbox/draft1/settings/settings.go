package settings

const version = "draft1"

type Settings struct {
	key      string
	user     string
	password string
}

// FromJSON implements trait.JSONSerializer.
func (*Settings) FromJSON(s string) error {
	panic("unimplemented")
}

// ToJSON implements trait.JSONSerializer.
func (s *Settings) ToJSON() (string, error) {
	panic("unimplemented")
}
