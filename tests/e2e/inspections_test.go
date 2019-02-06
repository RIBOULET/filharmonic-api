package e2e

import (
	"net/http"
	"testing"

	"github.com/MTES-MCT/filharmonic-api/models"
	"github.com/MTES-MCT/filharmonic-api/tests"
	"github.com/MTES-MCT/filharmonic-api/util"
)

func TestListInspectionsOwnedByInspecteur(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	results := tests.AuthInspecteur(e.GET("/inspections")).
		Expect().
		Status(http.StatusOK).
		JSON().Array()
	results.Length().Equal(4)
	results.First().Object().ValueEqual("id", 1)
	results.First().Object().Value("etablissement").Object().ValueEqual("id", 1)
	results.Element(1).Object().ValueEqual("id", 2)
	results.Element(1).Object().Value("etablissement").Object().ValueEqual("id", 3)
}

func TestListInspectionsOwnedByExploitant(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	results := tests.AuthExploitant(e.GET("/inspections")).
		Expect().
		Status(http.StatusOK).
		JSON().Array()
	results.Length().Equal(1)
	results.First().Object().ValueEqual("id", 1)
	results.First().Object().Value("etablissement").Object().ValueEqual("id", 1)
}

func TestGetInspectionAsInspecteur(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	inspection := tests.AuthInspecteur(e.GET("/inspections/{id}")).WithPath("id", "1").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	inspection.ValueEqual("id", 1)
	inspection.ValueEqual("date", "2018-09-01")
	inspection.Value("etablissement").Object().ValueEqual("id", 1)
	inspection.Value("themes").Array().Contains("Produits chimiques")
	inspecteurs := inspection.Value("inspecteurs").Array()
	inspecteurs.Length().Equal(1)
	inspecteurs.First().Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	suite := inspection.Value("suite").Object()
	suite.ValueEqual("delai", 30)
	suite.ValueEqual("type", "observation")
	suite.ValueEqual("synthese", "Observations à traiter")
	firstPointDeControle := inspection.Value("points_de_controle").Array().First().Object()
	firstPointDeControle.Value("references_reglementaires").Array().Contains("Article 3.2.3. de l'arrêté préfectoral du 28 juin 2017")
	firstPointDeControle.ValueEqual("sujet", "Mesure des émissions atmosphériques canalisées par un organisme extérieur")
	constat := firstPointDeControle.Value("constat").Object()
	constat.ValueEqual("type", "observation")
	constat.ValueEqual("remarques", "Il manque des choses")
	messages := firstPointDeControle.Value("messages").Array()
	messages.Length().Equal(4)
	firstMessage := messages.First().Object()
	firstMessage.ValueEqual("message", "Auriez-vous l'obligeance de me fournir le document approprié ?")
	firstMessage.Value("auteur").Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	firstMessage.Value("auteur").Object().NotContainsKey("password")
	firstCommentaire := inspection.Value("commentaires").Array().First().Object()
	firstCommentaire.ValueEqual("message", "Attention à l'article 243.")
	firstCommentaire.Value("auteur").Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	firstCommentaire.Value("auteur").Object().NotContainsKey("password")
	messageAvecPieceJointe := messages.Element(1).Object()
	piecesJointes := messageAvecPieceJointe.Value("pieces_jointes").Array().NotEmpty()
	firstPieceJointe := piecesJointes.First().Object()
	firstPieceJointe.ValueEqual("id", 1)
	firstPieceJointe.ValueEqual("nom", "photo-cuve.pdf")
	firstPieceJointe.ValueEqual("type", "application/pdf")
	firstPieceJointe.ValueEqual("taille", 7945)
}

func TestGetInspectionAsExploitantNotAllowed(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	result := tests.AuthExploitant(e.GET("/inspections/{id}")).WithPath("id", "2").
		Expect().
		Status(http.StatusNotFound).
		JSON().Object()
	result.ValueEqual("message", "not_found")
}

func TestGetInspectionAsExploitantAllowed(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	inspection := tests.AuthExploitant(e.GET("/inspections/{id}")).WithPath("id", "1").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	inspection.ValueEqual("id", 1)
	inspection.ValueEqual("date", "2018-09-01")
	inspection.Value("etablissement").Object().ValueEqual("id", 1)
	inspection.Value("themes").Array().Contains("Produits chimiques")
	firstPointDeControle := inspection.Value("points_de_controle").Array().First().Object()
	firstPointDeControle.Value("references_reglementaires").Array().Contains("Article 3.2.3. de l'arrêté préfectoral du 28 juin 2017")
	firstPointDeControle.ValueEqual("sujet", "Mesure des émissions atmosphériques canalisées par un organisme extérieur")
	messages := firstPointDeControle.Value("messages").Array()
	messages.Length().Equal(3)
	firstMessage := messages.First().Object()
	firstMessage.ValueEqual("message", "Auriez-vous l'obligeance de me fournir le document approprié ?")
	firstMessage.Value("auteur").Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	firstMessage.Value("auteur").Object().NotContainsKey("password")
	lastMessage := messages.Last().Object()
	lastMessage.ValueEqual("message", "Il manque un document.")
	lastMessage.Value("auteur").Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	lastMessage.Value("auteur").Object().NotContainsKey("password")
	inspection.NotContainsKey("commentaires")
}

func TestCreateInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	inspectionInput := models.Inspection{
		Date:            util.Date("2019-01-22"),
		Type:            models.TypeCourant,
		Annonce:         true,
		Origine:         models.OriginePlanControle,
		EtablissementId: 1,
		Inspecteurs: []models.User{
			models.User{
				Id: 3,
			},
			models.User{
				Id: 4,
			},
		},
		Themes: []string{
			"Incendie",
			"Produits chimiques",
		},
		Contexte: "Contrôles de début d'année",
	}

	inspectionId := tests.AuthInspecteur(e.POST("/inspections")).WithJSON(inspectionInput).
		Expect().
		Status(http.StatusOK).
		JSON().Object().Value("id").Raw()

	inspection := tests.AuthInspecteur(e.GET("/inspections/{id}")).WithPath("id", inspectionId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	inspection.ValueEqual("id", inspectionId)
	inspection.ValueEqual("etat", models.EtatPreparation)
	inspecteurs := inspection.Value("inspecteurs").Array()
	inspecteurs.Length().Equal(2)
	inspecteurs.First().Object().ValueEqual("id", 3)
	inspecteurs.First().Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	inspecteurs.Last().Object().ValueEqual("id", 4)
	inspecteurs.Last().Object().ValueEqual("email", "inspecteur2@filharmonic.com")
}

func TestUpdateInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	inspectionInput := models.Inspection{
		Id:      1,
		Date:    util.Date("2019-01-30"),
		Type:    models.TypeCourant,
		Annonce: true,
		Origine: models.OriginePlanControle,
		Inspecteurs: []models.User{
			models.User{
				Id: 3,
			},
			models.User{
				Id: 5,
			},
		},
		Themes: []string{
			"Incendie",
			"Produits chimiques",
		},
		Contexte: "Contrôles de début d'année",
	}

	tests.AuthInspecteur(e.PUT("/inspections/{id}")).WithPath("id", 1).WithJSON(inspectionInput).
		Expect().
		Status(http.StatusOK)

	inspection := tests.AuthInspecteur(e.GET("/inspections/{id}")).WithPath("id", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	inspection.ValueEqual("id", 1)
	inspection.ValueEqual("etat", models.EtatEnCours)
	inspection.ValueEqual("date", "2019-01-30")
	themes := inspection.Value("themes").Array()
	themes.Length().Equal(2)
	themes.Equal([]string{"Incendie", "Produits chimiques"})
	inspecteurs := inspection.Value("inspecteurs").Array()
	inspecteurs.Length().Equal(2)
	inspecteurs.First().Object().ValueEqual("id", 3)
	inspecteurs.First().Object().ValueEqual("email", "inspecteur1@filharmonic.com")
	inspecteurs.Last().Object().ValueEqual("id", 5)
	inspecteurs.Last().Object().ValueEqual("email", "inspecteur3@filharmonic.com")
}

func TestGetInspectionAsExploitantPointDeControleNonPublie(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	inspection := tests.AuthExploitant(e.GET("/inspections/{id}")).WithPath("id", "1").
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection.Value("points_de_controle").Array().Length().Equal(1)
}

func TestValidateInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthApprobateur(e.POST("/inspections/{id}/valider")).WithPath("id", 3).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection := tests.AuthApprobateur(e.GET("/inspections/{id}")).WithPath("id", 3).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection.ValueEqual("etat", models.EtatValide)
}
func TestPublishInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthInspecteur(e.POST("/inspections/{id}/publier")).WithPath("id", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection := tests.AuthInspecteur(e.GET("/inspections/{id}")).WithPath("id", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection.ValueEqual("etat", models.EtatEnCours)
}
func TestAskValidateInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthInspecteur(e.POST("/inspections/{id}/demandervalidation")).WithPath("id", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection := tests.AuthInspecteur(e.GET("/inspections/{id}")).WithPath("id", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection.ValueEqual("etat", models.EtatAttenteValidation)
}
func TestRejectInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthApprobateur(e.POST("/inspections/{id}/rejeter")).WithPath("id", 3).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection := tests.AuthApprobateur(e.GET("/inspections/{id}")).WithPath("id", 3).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inspection.ValueEqual("etat", models.EtatEnCours)
}

func TestAddFavoriToInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthInspecteur(e.POST("/inspections/{id}/favori")).WithPath("id", 2).
		Expect().
		Status(http.StatusOK)

	user := tests.AuthInspecteur(e.GET("/user")).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	user.ValueEqual("id", 3)
	user.ValueEqual("email", "inspecteur1@filharmonic.com")
	user.ValueEqual("profile", "inspecteur")
	favoris := user.Value("favoris").Array()
	favoris.Length().Equal(2)
	favori := favoris.Last().Object()
	favori.NotContainsKey("inspecteurs")
	favori.NotContainsKey("commentaires")
	favori.NotContainsKey("points_de_controle")
	favori.NotContainsKey("suite")
	favori.ValueEqual("id", 2)
	favori.ValueEqual("date", "2018-11-15")
	etablissement := favori.Value("etablissement").Object()
	etablissement.ValueEqual("adresse", "1 rue des cordeliers 69000 Lyon")
	etablissement.ValueEqual("nom", "Nom 3")
}
func TestRemoveFavoriToInspection(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthInspecteur(e.DELETE("/inspections/{id}/favori")).WithPath("id", 1).
		Expect().
		Status(http.StatusOK)

	user := tests.AuthInspecteur(e.GET("/user")).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	user.ValueEqual("id", 3)
	user.ValueEqual("email", "inspecteur1@filharmonic.com")
	user.ValueEqual("profile", "inspecteur")
	user.NotContainsKey("favoris")
}
