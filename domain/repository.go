package domain

import "github.com/MTES-MCT/filharmonic-api/models"

//go:generate mockery -name Repository

type Repository interface {
	FindEtablissementsByS3IC(ctx *UserContext, s3ic string) ([]models.Etablissement, error)
	GetEtablissementByID(ctx *UserContext, id int64) (*models.Etablissement, error)

	ListInspections(ctx *UserContext) ([]models.Inspection, error)
	CreateInspection(ctx *UserContext, inspection models.Inspection) (int64, error)
	UpdateInspection(ctx *UserContext, inspection models.Inspection) error
	GetInspectionByID(ctx *UserContext, id int64) (*models.Inspection, error)
	CheckEtatInspection(idInspection int64, etats []models.EtatInspection) (bool, error)
	ValidateInspection(ctx *UserContext, id int64) error

	CreateCommentaire(ctx *UserContext, idInspection int64, commentaire models.Commentaire) (int64, error)

	CreatePointDeControle(ctx *UserContext, idInspection int64, pointDeControle models.PointDeControle) (int64, error)
	UpdatePointDeControle(ctx *UserContext, idPointDeControle int64, pointDeControle models.PointDeControle) error
	DeletePointDeControle(ctx *UserContext, idPointDeControle int64) error
	PublishPointDeControle(ctx *UserContext, idPointDeControle int64) error
	CreateMessage(ctx *UserContext, idPointDeControle int64, message models.Message) (int64, error)
	LireMessage(ctx *UserContext, idMessage int64) error
	CheckUserAllowedMessage(ctx *UserContext, idMessage int64) (bool, error)
	CheckUserIsRecipient(ctx *UserContext, idMessage int64) (bool, error)

	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
	CheckUsersInspecteurs(ids []int64) (bool, error)
	CheckInspecteurAllowedInspection(ctx *UserContext, idInspection int64) (bool, error)
	CheckUserAllowedPointDeControle(ctx *UserContext, idPointDeControle int64) (bool, error)
}
