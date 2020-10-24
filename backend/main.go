package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pattapong1/app/controllers"
	_ "github.com/pattapong1/app/docs"
	"github.com/pattapong1/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	USERNAME  string
	USEREMAIL string
}

type Activities struct {
	Activity []Activity
}

type Activity struct {
	CLUBEACTIVITYNAME string
}

type Locations struct {
	Location []Location
}

type Location struct {
	CLUBELOCATIONNAME    string
	CLUBELOCATIONADDRESS string
}

type ClubTypesSlice struct {
	ClubTypes []ClubTypes
}

type ClubTypes struct {
	CLUBETYPENAME string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewActivityController(v1, client)
	controllers.NewClubController(v1, client)
	controllers.NewLocationController(v1, client)
	controllers.NewClubTypesController(v1, client)

	users := Users{
		User: []User{
			User{"Pattapong Sapmak", "pattapong999@gmail.com"},
			User{"Name Surname", "me@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUSERNAME(u.USERNAME).
			SetUSEREMAIL(u.USEREMAIL).
			Save(context.Background())
	}
	activitys := Activities{
		Activity: []Activity{
			Activity{"Planting Trees"},
			Activity{"Racing Car"},
			Activity{"Robot Development"},
		},
	}

	for _, a := range activitys.Activity {
		client.Activity.
			Create().
			SetCLUBEACTIVITYNAME(a.CLUBEACTIVITYNAME).
			Save(context.Background())
	}
	locations := Locations{
		Location: []Location{
			Location{"Suranaree University", "111, University Road, Suranaree Subdistrict, Mueang Nakhon Ratchasima District, Nakhon Ratchasima 30000"},
			Location{"Khon Kaen University", "Kaen Mo Din Daeng Building Khon Kaen University Mueang Khon Kaen District, Khon Kaen 40002"},
		},
	}

	for _, l := range locations.Location {
		client.Location.
			Create().
			SetCLUBELOCATIONNAME(l.CLUBELOCATIONNAME).
			SetCLUBELOCATIONADDRESS(l.CLUBELOCATIONADDRESS).
			Save(context.Background())
	}
	clubtypess := ClubTypesSlice{
		ClubTypes: []ClubTypes{
			ClubTypes{"Volunteer"},
			ClubTypes{"Sport"},
			ClubTypes{"Technology"},
		},
	}

	for _, t := range clubtypess.ClubTypes {
		client.ClubTypes.
			Create().
			SetCLUBETYPENAME(t.CLUBETYPENAME).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
