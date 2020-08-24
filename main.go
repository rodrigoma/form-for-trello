package main

import (
	"fmt"
	"formfortrello/lib"
	"formfortrello/setting"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type PageData struct {
	Title    string
	Msg      string
	MinDay	 int
}

func cleanup() {
	fmt.Println("Cleanup")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	setting.Setup()
	lib.Setup()

	// mux serves
	mx := mux.NewRouter()

	mx.Use(authMiddleware)

	mx.HandleFunc("/", Enter)
	mx.HandleFunc("/save", Save)
	mx.HandleFunc("/attach/{file}", Attach)

	// http serves
	http.Handle("/", mx)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":" + setting.ServerSetting.Port, nil)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		return
	})
}

func Enter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	log.Println("Request ->", r.URL.Path)

	tmpl := template.Must(template.ParseFiles(
		"templates/layouts/default.html",
				   "templates/form-trello.html"))

	r.ParseForm()

	tmpl.Execute(w, PageData{
		Title:    setting.AppSetting.Title,
		Msg:      r.Form.Get("msg"),
		MinDay:   setting.AppSetting.MinDay,
	})
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/save" {
		return
	}

	log.Println("Request ->", r.URL.Path)

	r.ParseMultipartForm(0)

	file, _, _ := r.FormFile("attachment")

	if file != nil {
		filename, err := lib.UploadFile(r)
		if err != nil {
			log.Fatalln("Error ->", err)
		}
		card := lib.CreateCard(r, filename)

		log.Println("Added Card with attach : ", card.Name)
	} else {
		card := lib.CreateCard(r, "")

		log.Println("Added Card : ", card.Name)
	}

	http.Redirect(w, r, "/?msg=Sua solicitação foi enviada com sucesso!", http.StatusPermanentRedirect)
}

func Attach(w http.ResponseWriter, r *http.Request) {

	log.Println("Request ->", r.URL.Path)

	vars := mux.Vars(r)

	log.Println("VARS ->", vars)

	dat, err := ioutil.ReadFile("/tmp/" + vars["file"])

	if err != nil {
		log.Fatalln("Error ->", err)
	}

	w.Write(dat)

}
