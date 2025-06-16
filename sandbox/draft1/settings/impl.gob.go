package settings

type GOB interface {
	FromGOB(data []byte) error
	ToGOB() ([]byte, error)
}

var _ GOB = (*Settings)(nil)

func (s *Settings) FromGOB(data []byte) error {

	panic("unimplemented")
}

func (s *Settings) ToGOB() ([]byte, error) {
	
	dto, err := s.ToDTO()
	if err != nil {
		return nil, err
	}


	


}
