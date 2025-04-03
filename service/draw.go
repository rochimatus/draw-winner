package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const maxSeed = 1000

type Service struct{}

func New() (Service, error) {
	return Service{}, nil
}

func (s Service) Draw(list []string) (string, error) {
	if len(list) == 0 {
		return "", errors.New("list is empty, none to draw")
	}

	if len(list) == 1 {
		return list[0], nil
	}

	num, err := s.getRandomIndex(list)
	if err != nil {
		return "", fmt.Errorf("random index error, %w", err)
	}

	return list[num], nil
}

func (s Service) getRandomIndex(list []string) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxSeed))
	if err != nil {
		return 0, err
	}

	randomed := nBig.Int64()

	return int(randomed) % len(list), nil
}
