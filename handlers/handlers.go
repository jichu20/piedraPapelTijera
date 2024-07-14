package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

const (
	templateDir   = "templates/"
	indexPage     = "index.html"
	newGamePage   = "new-game.html"
	gamePage      = "game.html"
	aboutPage     = "about.html"
	templateBase  = templateDir + "/base.html"
	templateIndex = templateDir + "/" + indexPage
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, indexPage, nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, newGamePage, nil)
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return

		}
		player.Name = r.Form.Get("name")

	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
		return

	}

	renderTemplate(w, gamePage, player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, aboutPage, nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Reiniciar Valores
func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0

}
