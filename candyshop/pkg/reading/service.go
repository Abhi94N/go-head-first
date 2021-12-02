package reading

type Repository interface {
	//function is tied to the internal function call bellow for all candies
	GetAllCandyNames() ([]string, error)

}

//service is tied to function below
type Service interface {
	GetAllCandyNames() ([]string, error)

}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}

}

func (s *service) GetAllCandyNames() ([]string, error) {
	cs, err := s.r.GetAllCandyNames()

	if err != nil {
		return nil, err
	}
	return cs, nil
}