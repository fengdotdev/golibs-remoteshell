package settings

type Getter interface {
	GetKey() string
	GetUser() string
	GetPassword() string
}

var _ Getter = (*Settings)(nil)

func (s *Settings) GetKey() string {
	return s.key
}

func (s *Settings) GetUser() string {
	return s.user
}

func (s *Settings) GetPassword() string {
	return s.password
}
