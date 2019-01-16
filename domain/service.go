package domain

import (
	"github.com/MTES-MCT/filharmonic-api/models"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ListEtablissements(s3ic string) ([]models.Etablissement, error) {
	return s.repo.FindEtablissementsByS3IC(s3ic)
}
