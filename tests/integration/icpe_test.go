package integration

import (
	"testing"

	"github.com/MTES-MCT/filharmonic-api/database"
	"github.com/MTES-MCT/filharmonic-api/database/icpe"
	"github.com/MTES-MCT/filharmonic-api/models"
	"github.com/MTES-MCT/filharmonic-api/tests"
	"github.com/stretchr/testify/require"
)

func TestLoadEtablissementsFromCSV(t *testing.T) {
	assert, application := tests.InitFuncDB(t, func(db *database.Database, assert *require.Assertions) {
		assert.NoError(db.Insert(&models.Etablissement{
			S3IC:     "0521.00217",
			Nom:      "FROMAGERIE OLD",
			Raison:   "FROMAGERIE OLD",
			Activite: "Fromage",
			Seveso:   "Haut",
			Iedmtd:   true,
			Adresse:  "PLACE DE LA LAITERIE 56000 VANNES",
		}))
	})

	err := icpe.LoadEtablissementsCSV("s3ic_ic_gen_fabnum.mini.csv", application.DB)
	assert.NoError(err)
	etablissements, err := application.Repo.ListEtablissements()
	assert.NoError(err)
	assert.NotEmpty(etablissements)
	assert.Len(etablissements, 54)
	etablissement := etablissements[len(etablissements)-1-3]
	assert.Equal("0521.00217", etablissement.S3IC)
	assert.Equal("FROMAGERIE BERTHAUT", etablissement.Raison)
	assert.Equal("FROMAGERIE BERTHAUT", etablissement.Nom)
	assert.Equal("10\",\"Fabrication de fromage", etablissement.Activite)
	assert.Equal("PLACE DU CHAMP DE FOIRE BP 5 21460 EPOISSES", etablissement.Adresse)
	assert.Equal("Non Seveso", etablissement.Seveso)
	assert.False(etablissement.Iedmtd)
}