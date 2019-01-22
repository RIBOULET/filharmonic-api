package database

import (
	"github.com/MTES-MCT/filharmonic-api/domain"
	"github.com/MTES-MCT/filharmonic-api/models"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func (repo *Repository) ListInspections(ctx *domain.UserContext) ([]models.Inspection, error) {
	var inspections []models.Inspection
	query := repo.db.client.Model(&inspections).Relation("Etablissement")
	if ctx.IsInspecteur() {
		query.Join("JOIN inspection_to_inspecteurs AS u").
			JoinOn("u.inspection_id = inspection.id").
			JoinOn("u.user_id = ?", ctx.User.Id)
	} else if ctx.IsExploitant() {
		query.Join("JOIN etablissement_to_exploitants AS u").
			JoinOn("u.etablissement_id = etablissement.id").
			JoinOn("u.user_id = ?", ctx.User.Id)
	}
	err := query.Select()
	return inspections, err
}

func (repo *Repository) CreateInspection(ctx *domain.UserContext, inspection models.Inspection) (int64, error) {
	inspection.Id = 0
	inspection.Etat = models.EtatPreparation
	err := repo.db.client.Insert(&inspection)
	if err != nil {
		return 0, err
	}

	for _, inspecteur := range inspection.Inspecteurs {
		err = repo.db.client.Insert(&models.InspectionToInspecteur{
			InspectionId: inspection.Id,
			UserId:       inspecteur.Id,
		})
		if err != nil {
			return 0, err
		}
	}
	return inspection.Id, nil
}

func (repo *Repository) GetInspectionByID(ctx *domain.UserContext, id int64) (*models.Inspection, error) {
	var inspection models.Inspection
	query := repo.db.client.Model(&inspection).
		Relation("Etablissement").
		Relation("PointsDeControle").
		Relation("PointsDeControle.Messages", func(q *orm.Query) (*orm.Query, error) {
			if ctx.IsExploitant() {
				q.Where("interne = FALSE")
			}
			return q.Order("date ASC"), nil
		}).
		Relation("PointsDeControle.Messages.Auteur").
		Relation("PointsDeControle.Messages.Auteur.id").
		Relation("PointsDeControle.Messages.Auteur.prenom").
		Relation("PointsDeControle.Messages.Auteur.nom").
		Relation("PointsDeControle.Messages.Auteur.email").
		Relation("PointsDeControle.Messages.Auteur.profile").
		Relation("Inspecteurs.id").
		Relation("Inspecteurs.prenom").
		Relation("Inspecteurs.nom").
		Relation("Inspecteurs.email").
		Relation("Inspecteurs.profile").
		Where(`inspection.id = ?`, id)
	if ctx.IsExploitant() {
		query.Join("JOIN etablissement_to_exploitants AS u").
			JoinOn("u.etablissement_id = etablissement.id").
			JoinOn("u.user_id = ?", ctx.User.Id)
	} else {
		query.Relation("Commentaires", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("date ASC"), nil
		}).
			Relation("Commentaires.Auteur").
			Relation("Commentaires.Auteur.id").
			Relation("Commentaires.Auteur.prenom").
			Relation("Commentaires.Auteur.nom").
			Relation("Commentaires.Auteur.email").
			Relation("Commentaires.Auteur.profile")
	}
	err := query.Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &inspection, err
}