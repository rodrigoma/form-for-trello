package setting

import (
	"github.com/kelseyhightower/envconfig"
	"log"
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

func Setup() {
	mapTo("APP", AppSetting)
	mapTo("SERVER", ServerSetting)
	mapTo("TRELLO", TrelloSetting)
}

func mapTo(section string, v interface{}) {
	err := envconfig.Process(section, v)
	if err != nil {
		log.Fatalf("envconfig.MapTo err: %v", err.Error())
	}
}
