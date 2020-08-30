package setting

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type App struct {
	Title    string
	MinDay   int
}
var AppSetting = &App{}

type Server struct {
	RunMode      string
	Port         string
}
var ServerSetting = &Server{}

type Trello struct {
	AppKey     string
	Token      string
	BoardID    string
	ListNumber int
}
var TrelloSetting = &Trello{}

var GoogleTokenJson = &oauth2.Token{}

type GoogleDrive struct {
	FolderId	string
}
var GoogleDriveSetting = &GoogleDrive{}

type GoogleCredentials struct {
	Credentials	string
}
var GoogleCredentialJson = &GoogleCredentials{}

func Setup() {
	log.Println("Start getting environment variables...")

	mapTo("APP", AppSetting)
	mapTo("SERVER", ServerSetting)
	mapTo("TRELLO", TrelloSetting)
	mapTo("GOOGLE_TOKEN_JSON", GoogleTokenJson)
	mapTo("GOOGLE_DRIVE", GoogleDriveSetting)
	mapTo("GOOGLE", GoogleCredentialJson)

	if os.Getenv("PORT") != "" {
		ServerSetting.Port = os.Getenv("PORT")
		log.Println("PORT automatic defined by service...")
	}

	log.Printf("Using PORT: %v\n", ServerSetting.Port)
}

func mapTo(section string, v interface{}) {
	log.Printf("Reading envvars: %v\n", section)
	err := envconfig.Process(section, v)
	if err != nil {
		log.Fatalf("envconfig.MapTo err: %v\n", err.Error())
	}
}
