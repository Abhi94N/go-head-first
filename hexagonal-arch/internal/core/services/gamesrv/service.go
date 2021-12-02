package gamesrv

import (
	"errors"

	"github.com/Abhi94N/go-head-first/hexagonal-arch/internal/core/domain"
	"github.com/Abhi94N/go-head-first/hexagonal-arch/internal/core/ports"
)

type service struct {
	gamesRepository ports.GamesRepository
	uidGen          uidgen.uidGen
}

func New(gamesRepository ports.GamesRepository) *service {
	return &service{
		gamesRepository: gamesRepository,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("get game from repository has failed")
	}

	return game, nil
}
