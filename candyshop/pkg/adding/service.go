package adding

type Repository interface {
	//function is tied to the internal function call bellow for all candies
	AddCandy(Candy) (string, error)
}

//service is tied to function below
type Service interface {
	AddCandy(Candy) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	//with repository interface as variable of a struct
	return &service{r}

}

func (s *service) AddCandy(c Candy) (string, error) {
	id, err := s.r.AddCandy(c)
	if err != nil {
		return "", err
	}

	return id, nil
}
