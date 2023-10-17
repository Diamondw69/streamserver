package service

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"video/internal/data"
)

var tmpl *template.Template

type Application struct {
	*sql.DB
}

func (a Application) StreamPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	device := vars["id"]
	tmpl, err := template.New("").ParseFiles("cmd/api/static/templates/stream.html", "cmd/api/static/templates/base.html")
	if err != nil {
		fmt.Println("Problem")
	}
	link, _ := data.LinkModel.GetByStatusAndDevice(data.LinkModel{DB: a.DB}, true, device)
	tmpl.ExecuteTemplate(w, "base", link)
}

func (a Application) InsertHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	links := data.Links{Name: r.Form.Get("name"),
		Link:   r.Form.Get("link") + "?autoplay=1&mute=1&loop=1",
		Status: true,
		Device: r.Form.Get("device"),
	}
	err = data.LinkModel.UpdateAllLinks(data.LinkModel{DB: a.DB}, links.Device)
	err = data.LinkModel.Insert(data.LinkModel{DB: a.DB}, &links)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", 303)
}

func (a Application) InsertPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("cmd/api/static/templates/insert.html", "cmd/api/static/templates/base.html")
	if err != nil {
		fmt.Println("error at insert html loader	")
	}
	tmpl.ExecuteTemplate(w, "base", nil)
}
func (a Application) AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("cmd/api/static/templates/adminpage.html", "cmd/api/static/templates/base.html")
	links, err := data.LinkModel.GetAllLinks(data.LinkModel{DB: a.DB})
	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w, "base", links)
}

func (a Application) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("cmd/api/static/templates/main.html", "cmd/api/static/templates/base.html")
	links, err := data.LinkModel.GetAllWorkingLinks(data.LinkModel{DB: a.DB})
	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w, "base", links)
}

func (a Application) SelectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	link, _ := data.LinkModel.GetByName(data.LinkModel{DB: a.DB}, id)
	err := data.LinkModel.UpdateAllLinks(data.LinkModel{DB: a.DB}, link.Device)
	if err != nil {
		fmt.Println(err)
	}
	_ = data.LinkModel.UpdateLink(data.LinkModel{DB: a.DB}, id)
	http.Redirect(w, r, "/", 303)
}
