// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/MTES-MCT/filharmonic-api/domain"
import mock "github.com/stretchr/testify/mock"
import models "github.com/MTES-MCT/filharmonic-api/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddFavoriToInspection provides a mock function with given fields: ctx, idInspection
func (_m *Repository) AddFavoriToInspection(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanCloreInspection provides a mock function with given fields: ctx, idInspection
func (_m *Repository) CanCloreInspection(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanCreateConstat provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) CanCreateConstat(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanCreatePointDeControle provides a mock function with given fields: ctx, idInspection
func (_m *Repository) CanCreatePointDeControle(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanDeleteConstat provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) CanDeleteConstat(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanUpdatePointDeControle provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) CanUpdatePointDeControle(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckCanCreateSuite provides a mock function with given fields: ctx, idInspection
func (_m *Repository) CheckCanCreateSuite(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckCanDeleteSuite provides a mock function with given fields: ctx, idInspection
func (_m *Repository) CheckCanDeleteSuite(ctx *domain.UserContext, idInspection int64) (bool, error) {
	ret := _m.Called(ctx, idInspection)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) bool); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, idInspection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckEtatInspection provides a mock function with given fields: idInspection, etats
func (_m *Repository) CheckEtatInspection(idInspection int64, etats []models.EtatInspection) (bool, error) {
	ret := _m.Called(idInspection, etats)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, []models.EtatInspection) bool); ok {
		r0 = rf(idInspection, etats)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, []models.EtatInspection) error); ok {
		r1 = rf(idInspection, etats)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckInspecteurAllowedInspection provides a mock function with given fields: ctx, idInspection
func (_m *Repository) CheckInspecteurAllowedInspection(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckInspectionHasNonConformites provides a mock function with given fields: idInspection
func (_m *Repository) CheckInspectionHasNonConformites(idInspection int64) (bool, error) {
	ret := _m.Called(idInspection)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(idInspection)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idInspection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUserAllowedMessage provides a mock function with given fields: ctx, idMessage
func (_m *Repository) CheckUserAllowedMessage(ctx *domain.UserContext, idMessage int64) (bool, error) {
	ret := _m.Called(ctx, idMessage)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) bool); ok {
		r0 = rf(ctx, idMessage)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, idMessage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUserAllowedPointDeControle provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) CheckUserAllowedPointDeControle(ctx *domain.UserContext, idPointDeControle int64) (bool, error) {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) bool); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, idPointDeControle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUserIsRecipient provides a mock function with given fields: ctx, idMessage
func (_m *Repository) CheckUserIsRecipient(ctx *domain.UserContext, idMessage int64) (bool, error) {
	ret := _m.Called(ctx, idMessage)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) bool); ok {
		r0 = rf(ctx, idMessage)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, idMessage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUsersInspecteurs provides a mock function with given fields: ids
func (_m *Repository) CheckUsersInspecteurs(ids []int64) (bool, error) {
	ret := _m.Called(ids)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]int64) bool); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCommentaire provides a mock function with given fields: ctx, idInspection, commentaire
func (_m *Repository) CreateCommentaire(ctx *domain.UserContext, idInspection int64, commentaire models.Commentaire) (int64, error) {
	ret := _m.Called(ctx, idInspection, commentaire)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.Commentaire) int64); ok {
		r0 = rf(ctx, idInspection, commentaire)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64, models.Commentaire) error); ok {
		r1 = rf(ctx, idInspection, commentaire)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateConstat provides a mock function with given fields: ctx, idPointDeControle, constat
func (_m *Repository) CreateConstat(ctx *domain.UserContext, idPointDeControle int64, constat models.Constat) (int64, error) {
	ret := _m.Called(ctx, idPointDeControle, constat)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.Constat) int64); ok {
		r0 = rf(ctx, idPointDeControle, constat)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64, models.Constat) error); ok {
		r1 = rf(ctx, idPointDeControle, constat)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEvenement provides a mock function with given fields: ctx, typeEvenement, idInspection, data
func (_m *Repository) CreateEvenement(ctx *domain.UserContext, typeEvenement models.TypeEvenement, idInspection int64, data map[string]interface{}) error {
	ret := _m.Called(ctx, typeEvenement, idInspection, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, models.TypeEvenement, int64, map[string]interface{}) error); ok {
		r0 = rf(ctx, typeEvenement, idInspection, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInspection provides a mock function with given fields: ctx, inspection
func (_m *Repository) CreateInspection(ctx *domain.UserContext, inspection models.Inspection) (int64, error) {
	ret := _m.Called(ctx, inspection)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.UserContext, models.Inspection) int64); ok {
		r0 = rf(ctx, inspection)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, models.Inspection) error); ok {
		r1 = rf(ctx, inspection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMessage provides a mock function with given fields: ctx, idPointDeControle, message
func (_m *Repository) CreateMessage(ctx *domain.UserContext, idPointDeControle int64, message models.Message) (int64, error) {
	ret := _m.Called(ctx, idPointDeControle, message)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.Message) int64); ok {
		r0 = rf(ctx, idPointDeControle, message)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64, models.Message) error); ok {
		r1 = rf(ctx, idPointDeControle, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePieceJointe provides a mock function with given fields: pieceJointe
func (_m *Repository) CreatePieceJointe(pieceJointe models.PieceJointe) (int64, error) {
	ret := _m.Called(pieceJointe)

	var r0 int64
	if rf, ok := ret.Get(0).(func(models.PieceJointe) int64); ok {
		r0 = rf(pieceJointe)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.PieceJointe) error); ok {
		r1 = rf(pieceJointe)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePointDeControle provides a mock function with given fields: ctx, idInspection, pointDeControle
func (_m *Repository) CreatePointDeControle(ctx *domain.UserContext, idInspection int64, pointDeControle models.PointDeControle) (int64, error) {
	ret := _m.Called(ctx, idInspection, pointDeControle)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.PointDeControle) int64); ok {
		r0 = rf(ctx, idInspection, pointDeControle)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64, models.PointDeControle) error); ok {
		r1 = rf(ctx, idInspection, pointDeControle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRapport provides a mock function with given fields: idInspection, rapport
func (_m *Repository) CreateRapport(idInspection int64, rapport models.Rapport) error {
	ret := _m.Called(idInspection, rapport)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, models.Rapport) error); ok {
		r0 = rf(idInspection, rapport)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSuite provides a mock function with given fields: ctx, idInspection, suite
func (_m *Repository) CreateSuite(ctx *domain.UserContext, idInspection int64, suite models.Suite) (int64, error) {
	ret := _m.Called(ctx, idInspection, suite)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.Suite) int64); ok {
		r0 = rf(ctx, idInspection, suite)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64, models.Suite) error); ok {
		r1 = rf(ctx, idInspection, suite)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTheme provides a mock function with given fields: theme
func (_m *Repository) CreateTheme(theme models.Theme) (int64, error) {
	ret := _m.Called(theme)

	var r0 int64
	if rf, ok := ret.Get(0).(func(models.Theme) int64); ok {
		r0 = rf(theme)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Theme) error); ok {
		r1 = rf(theme)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConstat provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) DeleteConstat(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePointDeControle provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) DeletePointDeControle(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSuite provides a mock function with given fields: ctx, idInspection
func (_m *Repository) DeleteSuite(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTheme provides a mock function with given fields: idTheme
func (_m *Repository) DeleteTheme(idTheme int64) error {
	ret := _m.Called(idTheme)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(idTheme)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindEtablissements provides a mock function with given fields: ctx, filter
func (_m *Repository) FindEtablissements(ctx *domain.UserContext, filter domain.ListEtablissementsFilter) (*models.FindEtablissementResults, error) {
	ret := _m.Called(ctx, filter)

	var r0 *models.FindEtablissementResults
	if rf, ok := ret.Get(0).(func(*domain.UserContext, domain.ListEtablissementsFilter) *models.FindEtablissementResults); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FindEtablissementResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, domain.ListEtablissementsFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUsers provides a mock function with given fields: filters
func (_m *Repository) FindUsers(filters domain.ListUsersFilters) ([]models.User, error) {
	ret := _m.Called(filters)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(domain.ListUsersFilters) []models.User); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.ListUsersFilters) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEtablissementByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetEtablissementByID(ctx *domain.UserContext, id int64) (*models.Etablissement, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Etablissement
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) *models.Etablissement); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Etablissement)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEtatInspectionByPointDeControleID provides a mock function with given fields: idPointDeControle
func (_m *Repository) GetEtatInspectionByPointDeControleID(idPointDeControle int64) (models.EtatInspection, error) {
	ret := _m.Called(idPointDeControle)

	var r0 models.EtatInspection
	if rf, ok := ret.Get(0).(func(int64) models.EtatInspection); ok {
		r0 = rf(idPointDeControle)
	} else {
		r0 = ret.Get(0).(models.EtatInspection)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idPointDeControle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInspectionByID provides a mock function with given fields: ctx, id, filter
func (_m *Repository) GetInspectionByID(ctx *domain.UserContext, id int64, filter domain.InspectionFilter) (*models.Inspection, error) {
	ret := _m.Called(ctx, id, filter)

	var r0 *models.Inspection
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, domain.InspectionFilter) *models.Inspection); ok {
		r0 = rf(ctx, id, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Inspection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64, domain.InspectionFilter) error); ok {
		r1 = rf(ctx, id, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInspectionTypesConstatsSuiteByID provides a mock function with given fields: idInspection
func (_m *Repository) GetInspectionTypesConstatsSuiteByID(idInspection int64) (*models.Inspection, error) {
	ret := _m.Called(idInspection)

	var r0 *models.Inspection
	if rf, ok := ret.Get(0).(func(int64) *models.Inspection); ok {
		r0 = rf(idInspection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Inspection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idInspection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPieceJointe provides a mock function with given fields: ctx, id
func (_m *Repository) GetPieceJointe(ctx *domain.UserContext, id int64) (*models.PieceJointe, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.PieceJointe
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) *models.PieceJointe); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PieceJointe)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRapport provides a mock function with given fields: ctx, idInspection
func (_m *Repository) GetRapport(ctx *domain.UserContext, idInspection int64) (*models.Rapport, error) {
	ret := _m.Called(ctx, idInspection)

	var r0 *models.Rapport
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) *models.Rapport); ok {
		r0 = rf(ctx, idInspection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rapport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, int64) error); ok {
		r1 = rf(ctx, idInspection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecapsValidation provides a mock function with given fields: idInspection
func (_m *Repository) GetRecapsValidation(idInspection int64) ([]domain.RecapValidationInspection, error) {
	ret := _m.Called(idInspection)

	var r0 []domain.RecapValidationInspection
	if rf, ok := ret.Get(0).(func(int64) []domain.RecapValidationInspection); ok {
		r0 = rf(idInspection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RecapValidationInspection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idInspection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTypeConstatByPointDeControleID provides a mock function with given fields: idPointDeControle
func (_m *Repository) GetTypeConstatByPointDeControleID(idPointDeControle int64) (models.TypeConstat, error) {
	ret := _m.Called(idPointDeControle)

	var r0 models.TypeConstat
	if rf, ok := ret.Get(0).(func(int64) models.TypeConstat); ok {
		r0 = rf(idPointDeControle)
	} else {
		r0 = ret.Get(0).(models.TypeConstat)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(idPointDeControle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *Repository) GetUserByEmail(email string) (*models.User, error) {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *Repository) GetUserByID(id int64) (*models.User, error) {
	ret := _m.Called(id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(int64) *models.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LireMessage provides a mock function with given fields: ctx, idMessage
func (_m *Repository) LireMessage(ctx *domain.UserContext, idMessage int64) error {
	ret := _m.Called(ctx, idMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LireNotifications provides a mock function with given fields: ctx, ids
func (_m *Repository) LireNotifications(ctx *domain.UserContext, ids []int64) error {
	ret := _m.Called(ctx, ids)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, []int64) error); ok {
		r0 = rf(ctx, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListInspections provides a mock function with given fields: ctx, filter
func (_m *Repository) ListInspections(ctx *domain.UserContext, filter domain.ListInspectionsFilter) ([]models.Inspection, error) {
	ret := _m.Called(ctx, filter)

	var r0 []models.Inspection
	if rf, ok := ret.Get(0).(func(*domain.UserContext, domain.ListInspectionsFilter) []models.Inspection); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Inspection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, domain.ListInspectionsFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInspectionsEcheancesProches provides a mock function with given fields: seuilRappelEcheancservicees
func (_m *Repository) ListInspectionsEcheancesProches(seuilRappelEcheancservicees float32) ([]domain.InspectionEcheancesProches, error) {
	ret := _m.Called(seuilRappelEcheancservicees)

	var r0 []domain.InspectionEcheancesProches
	if rf, ok := ret.Get(0).(func(float32) []domain.InspectionEcheancesProches); ok {
		r0 = rf(seuilRappelEcheancservicees)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.InspectionEcheancesProches)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float32) error); ok {
		r1 = rf(seuilRappelEcheancservicees)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInspectionsExpirationDelais provides a mock function with given fields:
func (_m *Repository) ListInspectionsExpirationDelais() ([]domain.InspectionExpirationDelais, error) {
	ret := _m.Called()

	var r0 []domain.InspectionExpirationDelais
	if rf, ok := ret.Get(0).(func() []domain.InspectionExpirationDelais); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.InspectionExpirationDelais)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInspectionsFavorites provides a mock function with given fields: ctx
func (_m *Repository) ListInspectionsFavorites(ctx *domain.UserContext) ([]models.Inspection, error) {
	ret := _m.Called(ctx)

	var r0 []models.Inspection
	if rf, ok := ret.Get(0).(func(*domain.UserContext) []models.Inspection); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Inspection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNotifications provides a mock function with given fields: ctx, filter
func (_m *Repository) ListNotifications(ctx *domain.UserContext, filter *domain.ListNotificationsFilter) ([]models.Notification, error) {
	ret := _m.Called(ctx, filter)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(*domain.UserContext, *domain.ListNotificationsFilter) []models.Notification); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserContext, *domain.ListNotificationsFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNouveauxMessages provides a mock function with given fields:
func (_m *Repository) ListNouveauxMessages() ([]domain.NouveauxMessagesUser, error) {
	ret := _m.Called()

	var r0 []domain.NouveauxMessagesUser
	if rf, ok := ret.Get(0).(func() []domain.NouveauxMessagesUser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.NouveauxMessagesUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListThemes provides a mock function with given fields:
func (_m *Repository) ListThemes() ([]models.Theme, error) {
	ret := _m.Called()

	var r0 []models.Theme
	if rf, ok := ret.Get(0).(func() []models.Theme); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Theme)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishPointDeControle provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) PublishPointDeControle(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RejectInspection provides a mock function with given fields: id, motifRejet
func (_m *Repository) RejectInspection(id int64, motifRejet string) error {
	ret := _m.Called(id, motifRejet)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(id, motifRejet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveFavoriToInspection provides a mock function with given fields: ctx, idInspection
func (_m *Repository) RemoveFavoriToInspection(ctx *domain.UserContext, idInspection int64) error {
	ret := _m.Called(ctx, idInspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idInspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResolveConstat provides a mock function with given fields: ctx, idPointDeControle
func (_m *Repository) ResolveConstat(ctx *domain.UserContext, idPointDeControle int64) error {
	ret := _m.Called(ctx, idPointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64) error); ok {
		r0 = rf(ctx, idPointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEtatInspection provides a mock function with given fields: ctx, id, etat
func (_m *Repository) UpdateEtatInspection(ctx *domain.UserContext, id int64, etat models.EtatInspection) error {
	ret := _m.Called(ctx, id, etat)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.EtatInspection) error); ok {
		r0 = rf(ctx, id, etat)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateInspection provides a mock function with given fields: ctx, inspection
func (_m *Repository) UpdateInspection(ctx *domain.UserContext, inspection models.Inspection) error {
	ret := _m.Called(ctx, inspection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, models.Inspection) error); ok {
		r0 = rf(ctx, inspection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePointDeControle provides a mock function with given fields: ctx, idPointDeControle, pointDeControle
func (_m *Repository) UpdatePointDeControle(ctx *domain.UserContext, idPointDeControle int64, pointDeControle models.PointDeControle) error {
	ret := _m.Called(ctx, idPointDeControle, pointDeControle)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.PointDeControle) error); ok {
		r0 = rf(ctx, idPointDeControle, pointDeControle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRappelsEcheancesEnvoyes provides a mock function with given fields: constatsIds
func (_m *Repository) UpdateRappelsEcheancesEnvoyes(constatsIds []int64) error {
	ret := _m.Called(constatsIds)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int64) error); ok {
		r0 = rf(constatsIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSuite provides a mock function with given fields: ctx, idInspection, suite
func (_m *Repository) UpdateSuite(ctx *domain.UserContext, idInspection int64, suite models.Suite) error {
	ret := _m.Called(ctx, idInspection, suite)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserContext, int64, models.Suite) error); ok {
		r0 = rf(ctx, idInspection, suite)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateInspection provides a mock function with given fields: id, etatCible
func (_m *Repository) ValidateInspection(id int64, etatCible models.EtatInspection) error {
	ret := _m.Called(id, etatCible)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, models.EtatInspection) error); ok {
		r0 = rf(id, etatCible)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
