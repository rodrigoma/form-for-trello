package lib

import (
	"formfortrello/setting"
	"formfortrello/utils"
	"github.com/adlio/trello"
	"google.golang.org/api/drive/v3"
	"log"
	"net/http"
)

var client *trello.Client
var board *trello.Board
var list *trello.List

var err error

func Setup() {
	client = trello.NewClient(setting.TrelloSetting.AppKey, setting.TrelloSetting.Token)

	board, err = client.GetBoard(setting.TrelloSetting.BoardID, trello.Defaults())

	if err != nil {
		log.Fatalln("Error: Selecting Board:", err.Error())
	}

	lists, err := board.GetLists(trello.Defaults())

	if err != nil {
		log.Fatalln("Error: Selecting Lists of Board:", err.Error())
	}

	list = lists[setting.TrelloSetting.ListNumber]
}

func CreateCard(r *http.Request, fileGD *drive.File) (card *trello.Card) {
	form := r.PostForm

	dt := utils.FormatDate(form.Get("date"))

	name := "[" + dt + "] - " +
		form.Get("minister") + " - " +
		form.Get("event")

	desc := "- Dados Contato" + "\n\n" +
		"**NOME:** " + form.Get("name") + "\n" +
		"**E-MAIL:** " + form.Get("email") + "\n" +
		"**TEL:** " + form.Get("phone") + "\n\n" +
		"- Dados Evento/Ação" + "\n\n" +
		"**DATA/HORA:** " + dt + "\n" +
		"**MINISTÉRIO:** " + form.Get("minister") + "\n" +
		"**EVENTO/AÇÃO:** " + form.Get("event") + "\n" +
		"**AO VIVO?:** " + form.Get("broadcastOptions") + "\n" +
		"**TEMA:** " + form.Get("subject") + "\n" +
		"**VERSÍCULO:** " + form.Get("verse") + "\n" +
		"**INFOS AD.:** " + form.Get("infos")

	card = &trello.Card{
		IDList: list.ID,
		Name:   name,
		Desc:   desc,
	}

	err = client.CreateCard(card, trello.Defaults())

	if err != nil {
		log.Fatalln("Error on creating card", err.Error())
	}

	if fileGD != nil {
		attachUrl := utils.GetGDSharedUrl(fileGD.Id)
		addAttach(card, attachUrl, fileGD.Name)
	}

	return
}

func addAttach(card *trello.Card, url string, filenameGD string) {
	attach := trello.Attachment{URL: url, Name: filenameGD}

	err = card.AddURLAttachment(&attach)

	if err != nil {
		log.Fatalf("Add attachment error : %v", err)
	}
}