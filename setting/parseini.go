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

type Google struct {
	Credentials	string
}
var GoogleSetting = &Google{}

type GoogleDrive struct {
	FolderId	string
}
var GoogleDriveSetting = &GoogleDrive{}

var GoogleTokenJson = &oauth2.Token{}

func Setup() {
	mapTo("APP", AppSetting)
	mapTo("SERVER", ServerSetting)
	mapTo("TRELLO", TrelloSetting)
	mapTo("GOOGLE_TOKEN_JSON", GoogleTokenJson)
	mapTo("GOOGLE", GoogleSetting)
	mapTo("GOOGLE_DRIVE", GoogleDriveSetting)

	if os.Getenv("PORT") != "" {
		ServerSetting.Port = os.Getenv("PORT")
	}
}

func mapTo(section string, v interface{}) {
	err := envconfig.Process(section, v)
	if err != nil {
		log.Fatalf("envconfig.MapTo err: %v", err.Error())
	}
}
