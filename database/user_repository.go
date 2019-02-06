package database

import (
	"github.com/MTES-MCT/filharmonic-api/domain"
	"github.com/MTES-MCT/filharmonic-api/models"
	"github.com/go-pg/pg"
)

func (repo *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.client.Model(&user).
		Relation("Favoris.id").
		Relation("Favoris.date").
		Relation("Favoris.Etablissement.nom").
		Relation("Favoris.Etablissement.adresse").
		Where("email = ?", email).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (repo *Repository) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	err := repo.db.client.Model(&user).
		Relation("Favoris.id").
		Relation("Favoris.date").
		Relation("Favoris.Etablissement.nom").
		Relation("Favoris.Etablissement.adresse").
		Where("id = ?", id).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (repo *Repository) CheckUsersInspecteurs(ids []int64) (bool, error) {
	goodProfileCount, err := repo.db.client.Model(&models.User{}).
		Where("id IN (?)", pg.In(ids)).
		Where("profile <> ?", models.ProfilExploitant).
		Count()
	return goodProfileCount == len(ids), err
}

func (repo *Repository) FindUsers(filters domain.ListUsersFilters) ([]models.User, error) {
	users := []models.User{}
	query := repo.db.client.Model(&users)

	if filters.Inspecteurs {
		query.Where("profile = ?", models.ProfilInspecteur)
	}

	err := query.Select()
	if err == pg.ErrNoRows {
		return users, nil
	}
	return users, err
}
