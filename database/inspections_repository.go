package database

import (
	"time"

	"github.com/MTES-MCT/filharmonic-api/domain"
	"github.com/MTES-MCT/filharmonic-api/models"
	"github.com/MTES-MCT/filharmonic-api/util"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/go-pg/pg/types"
)

func (repo *Repository) ListInspections(ctx *domain.UserContext, filter domain.ListInspectionsFilter) ([]models.Inspection, error) {
	inspections := []models.Inspection{}
	query := repo.db.client.Model(&models.Inspection{}).
		Relation("Etablissement")
	if ctx.IsExploitant() {
		query.Where(`inspection.etat <> ?`, models.EtatPreparation).
			Join("JOIN etablissement_to_exploitants AS u").
			JoinOn("u.etablissement_id = inspection.etablissement_id").
			JoinOn("u.user_id = ?", ctx.User.Id)
	} else {
		if filter.Assigned {
			query.Join("JOIN inspection_to_inspecteurs AS u").
				JoinOn("u.inspection_id = inspection.id").
				JoinOn("u.user_id = ?", ctx.User.Id)
		}
	}

	query.ColumnExpr("inspection.*").
		ColumnExpr(`(
		SELECT count(m.id)
		FROM point_de_controles AS p
		JOIN messages AS m
			ON m.point_de_controle_id = p.id
			AND m.lu IS FALSE
			AND m.interne IS FALSE
		JOIN users AS auteur
			ON auteur.id = m.auteur_id
			AND auteur.profile in (?)
		WHERE p.inspection_id = inspection.id
		AND p.publie IS TRUE
		) AS nb_messages_non_lus`, pg.In(getDestinataires(ctx))).
		Order("inspection.id ASC")

	err := query.Select(&inspections)
	if err != nil {
		return nil, err
	}
	if ctx.IsExploitant() {
		for i, _ := range inspections {
			inspections[i].ValidationRejetee = false
			inspections[i].MotifRejetValidation = ""
		}
	}
	return inspections, nil
}

func (repo *Repository) ListInspectionsFavorites(ctx *domain.UserContext) ([]models.Inspection, error) {
	inspections := []models.Inspection{}
	err := repo.db.client.Model(&models.Inspection{}).
		Relation("Etablissement").
		Join("JOIN user_to_favoris AS favoris").
		JoinOn("favoris.inspection_id = inspection.id").
		JoinOn("favoris.user_id = ?", ctx.User.Id).
		Select(&inspections)
	return inspections, err
}

func (repo *Repository) CreateInspection(ctx *domain.UserContext, inspection models.Inspection) (int64, error) {
	inspectionId := int64(0)
	err := repo.db.client.RunInTransaction(func(tx *pg.Tx) error {
		inspection.Id = 0
		inspection.Etat = models.EtatPreparation
		err := tx.Insert(&inspection)
		if err != nil {
			return err
		}

		for _, inspecteur := range inspection.Inspecteurs {
			err = tx.Insert(&models.InspectionToInspecteur{
				InspectionId: inspection.Id,
				UserId:       inspecteur.Id,
			})
			if err != nil {
				return err
			}
		}
		inspectionId = inspection.Id
		err = repo.CreateEvenementTx(tx, ctx, models.EvenementCreationInspection, inspectionId, nil)
		return err
	})
	return inspectionId, err
}

func (repo *Repository) UpdateInspection(ctx *domain.UserContext, inspection models.Inspection) error {
	return repo.db.client.RunInTransaction(func(tx *pg.Tx) error {
		columns := []string{"date", "type", "origine", "annonce", "circonstance", "detail_circonstance", "contexte", "themes"}
		_, err := tx.Model(&inspection).Column(columns...).WherePK().Update()
		if err != nil {
			return err
		}
		_, err = tx.Model(&models.InspectionToInspecteur{}).Where("inspection_id = ?", inspection.Id).Delete()
		if err != nil {
			return err
		}
		for _, inspecteur := range inspection.Inspecteurs {
			err = tx.Insert(&models.InspectionToInspecteur{
				InspectionId: inspection.Id,
				UserId:       inspecteur.Id,
			})
			if err != nil {
				return err
			}
		}
		err = repo.CreateEvenementTx(tx, ctx, models.EvenementModificationInspection, inspection.Id, nil)
		return err
	})
}

func (repo *Repository) GetInspectionByID(ctx *domain.UserContext, id int64, filter domain.InspectionFilter) (*models.Inspection, error) {
	var inspection models.Inspection
	query := repo.db.client.Model(&inspection).
		Relation("Etablissement").
		Relation("Etablissement.Exploitants").
		Relation("PointsDeControle", func(q *orm.Query) (*orm.Query, error) {
			if ctx.IsExploitant() || filter.OmitPointsDeControleNonPublies {
				q.Where("publie = TRUE")
			}
			return q.Order("id ASC"), nil
		}).
		Relation("PointsDeControle.Constat").
		Relation("PointsDeControle.Messages", func(q *orm.Query) (*orm.Query, error) {
			if ctx.IsExploitant() {
				q.Where("interne = FALSE")
			}
			return q.Order("date ASC"), nil
		}).
		Relation("PointsDeControle.Messages.Auteur").
		Relation("PointsDeControle.Messages.PiecesJointes").
		Relation("Inspecteurs").
		Relation("Suite").
		Relation("Rapport").
		Where(`inspection.id = ?`, id)

	if ctx.IsExploitant() {
		query.Where(`inspection.etat <> ?`, models.EtatPreparation).
			Join("JOIN etablissement_to_exploitants AS u").
			JoinOn("u.etablissement_id = etablissement.id").
			JoinOn("u.user_id = ?", ctx.User.Id)
	} else {
		query.Relation("Commentaires", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("date ASC"), nil
		}).
			Relation("Commentaires.Auteur").
			Relation("Commentaires.PiecesJointes").
			Relation("Evenements").
			Relation("Evenements.Auteur")
	}
	err := query.Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	if ctx.IsExploitant() {
		inspection.ValidationRejetee = false
		inspection.MotifRejetValidation = ""
	}
	return &inspection, err
}

func (repo *Repository) CheckInspecteurAllowedInspection(ctx *domain.UserContext, id int64) error {
	count, err := repo.db.client.Model(&models.InspectionToInspecteur{}).
		Where("inspection_id = ?", id).
		Where("user_id = ?", ctx.User.Id).
		Count()
	if err != nil {
		return err
	}
	if count != 1 {
		return domain.ErrInspecteurNonAffecte
	}
	return nil
}

func (repo *Repository) CheckInspectionHasNonConformites(id int64) (bool, error) {
	count, err := repo.db.client.Model(&models.Inspection{}).
		Relation("Suite.Type").
		Where(`"inspection".id = ?`, id).
		Where(`"suite".type <> ?`, models.TypeSuiteAucune).
		Count()
	return count == 1, err
}

func (repo *Repository) GetInspectionTypesConstatsSuiteByID(id int64) (*models.Inspection, error) {
	inspection := &models.Inspection{
		Id: id,
	}
	err := repo.db.client.Model(inspection).
		Column("_").
		Relation("PointsDeControle.inspection_id").
		Relation("PointsDeControle.constat_id").
		Relation("PointsDeControle.publie").
		Relation("PointsDeControle.Constat.type").
		Relation("Suite.type").
		WherePK().
		Select()
	if err != nil {
		return nil, err
	}
	return inspection, nil
}

func (repo *Repository) CheckEtatInspection(id int64, etats []models.EtatInspection) (bool, error) {
	count, err := repo.db.client.Model(&models.Inspection{}).
		Where("id = ?", id).
		Where("etat in (?)", pg.In(etats)).
		Count()
	return count == 1, err
}

func (repo *Repository) CanCloreInspection(ctx *domain.UserContext, idInspection int64) error {
	count, err := repo.db.client.Model(&models.Inspection{}).
		Join("JOIN point_de_controles AS point_de_controle").
		JoinOn("inspection.id = point_de_controle.inspection_id").
		Join("JOIN constats as constat").
		JoinOn("constat.id = point_de_controle.constat_id").
		Where("inspection.id = ?", idInspection).
		Where("constat.type <> ?", models.TypeConstatConforme).
		Where("constat.date_resolution IS NULL").
		Where("constat.echeance_resolution > ?", time.Now()).
		Count()
	if err != nil {
		return err
	}
	if count >= 1 {
		return domain.ErrClotureInspectionImpossible
	}
	return nil
}

func (repo *Repository) UpdateEtatInspection(ctx *domain.UserContext, id int64, etat models.EtatInspection) error {
	inspection := models.Inspection{
		Id:   id,
		Etat: etat,
	}
	columns := []string{"etat"}
	_, err := repo.db.client.Model(&inspection).Column(columns...).WherePK().Update()
	return err
}

func (repo *Repository) ValidateInspection(id int64, etatCible models.EtatInspection) error {
	return repo.db.client.RunInTransaction(func(tx *pg.Tx) error {
		inspection := models.Inspection{
			Id:   id,
			Etat: etatCible,
			DateValidation: types.NullTime{
				Time: time.Now(),
			},
			ValidationRejetee:    false,
			MotifRejetValidation: "",
		}
		columns := []string{"etat", "date_validation", "validation_rejetee", "motif_rejet_validation"}
		_, err := tx.Model(&inspection).Column(columns...).WherePK().Update()
		if err != nil {
			return err
		}

		_, err = tx.Exec(`UPDATE constats as c
										 SET echeance_resolution = date(
												 i.date_validation + (case when c.delai_unite='jours' then make_interval(days => c.delai_nombre) else make_interval(months => c.delai_nombre) end)
												)
										 FROM point_de_controles as p
										 JOIN inspections as i
										   ON i.id = p.inspection_id
										 WHERE p.constat_id = c.id
										   AND i.id = ?`, id)
		return err
	})
}

func (repo *Repository) RejectInspection(id int64, motifRejet string) error {
	inspection := models.Inspection{
		Id:                   id,
		Etat:                 models.EtatEnCours,
		ValidationRejetee:    true,
		MotifRejetValidation: motifRejet,
	}
	columns := []string{"etat", "validation_rejetee", "motif_rejet_validation"}
	_, err := repo.db.client.Model(&inspection).Column(columns...).WherePK().Update()
	return err
}

func (repo *Repository) AddFavoriToInspection(ctx *domain.UserContext, idInspection int64) error {
	return repo.db.client.Insert(&models.UserToFavori{
		InspectionId: idInspection,
		UserId:       ctx.User.Id,
	})
}

func (repo *Repository) RemoveFavoriToInspection(ctx *domain.UserContext, idInspection int64) error {
	return repo.db.client.Delete(&models.UserToFavori{
		InspectionId: idInspection,
		UserId:       ctx.User.Id,
	})
}

func (repo *Repository) GetRecapsValidation(idInspection int64) ([]domain.RecapValidationInspection, error) {
	recaps := []domain.RecapValidationInspection{}
	_, err := repo.db.client.Query(&recaps, `select
		inspection.id as inspection_id,
		inspection.date as date_inspection,
		inspection.etat = ? as non_conformites,
		etablissement.raison as raison_etablissement,
		users.prenom || ' ' || users.nom as destinataire__nom,
		users.email as destinataire__email
	from inspections as inspection
	join etablissements as etablissement
		on etablissement.id = inspection.etablissement_id
	join etablissement_to_exploitants as exploitant
		on etablissement.id = exploitant.etablissement_id
	join users
		on users.id = exploitant.user_id
	where
		inspection.id = ?`, models.EtatTraitementNonConformites, idInspection)
	if err != nil {
		return nil, err
	}
	return recaps, nil
}

func (repo *Repository) ListInspectionsExpirationDelais() ([]domain.InspectionExpirationDelais, error) {
	inspectionExpirationDelais := []domain.InspectionExpirationDelais{}
	_, err := repo.db.client.Query(&inspectionExpirationDelais, `select
		distinct(inspection.id) as inspection_id,
		inspection.date as date_inspection,
		etablissement.raison as raison_etablissement,
		users.prenom || ' ' || users.nom as destinataire__nom,
		users.email as destinataire__email
	from inspections as inspection
	join etablissements as etablissement
		on etablissement.id = inspection.etablissement_id
	join inspection_to_inspecteurs as inspecteur
		on inspection.id = inspecteur.inspection_id
	join point_de_controles AS p
		ON p.inspection_id = inspection.id
	join constats AS constat
		on p.constat_id = constat.id
		and constat.date_resolution is null
		and constat.echeance_resolution <= ?
	join users
		on users.id = inspecteur.user_id`, util.Now())
	if err != nil {
		return nil, err
	}
	return inspectionExpirationDelais, nil
}

func (repo *Repository) ListInspectionsEcheancesProches(seuilRappelEcheances float32) ([]domain.InspectionEcheancesProches, error) {
	inspectionEcheancesProches := []domain.InspectionEcheancesProches{}
	_, err := repo.db.client.Query(&inspectionEcheancesProches, `select
		inspection.id as inspection_id,
		constat.id as constat_id,
		inspection.date as date_inspection,
		etablissement.raison as raison_etablissement,
		users.prenom || ' ' || users.nom as destinataire__nom,
		users.email as destinataire__email
	from inspections as inspection
	join etablissements as etablissement
		on etablissement.id = inspection.etablissement_id
	join inspection_to_inspecteurs as inspecteur
		on inspection.id = inspecteur.inspection_id
	join point_de_controles AS p
		ON p.inspection_id = inspection.id
	join constats AS constat
		on p.constat_id = constat.id
		and constat.date_resolution is null
		and constat.rappel_echeances_envoye is false
		and ? between date(
			constat.echeance_resolution - (case when constat.delai_unite='jours' then make_interval(days => cast(ceil(? * constat.delai_nombre) as int)) else make_interval(months => cast(ceil(? * constat.delai_nombre) as int)) end)
		   ) and constat.echeance_resolution
	join users
		on users.id = inspecteur.user_id
	order by inspection.id asc`, util.Now(), seuilRappelEcheances, seuilRappelEcheances)
	if err != nil {
		return nil, err
	}
	return inspectionEcheancesProches, nil
}

func (repo *Repository) UpdateRappelsEcheancesEnvoyes(constatIds []int64) error {
	if len(constatIds) == 0 {
		return nil
	}
	constat := models.Constat{
		RappelEcheancesEnvoye: true,
	}
	_, err := repo.db.client.Model(&constat).
		Column("rappel_echeances_envoye").
		Where("id in (?)", pg.In(constatIds)).
		Update()
	return err
}
