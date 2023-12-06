package service

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"video/internal/data"
	"video/internal/helpers"
)

type Application struct {
	*sql.DB
}

type Answer struct {
	message string
	status  string
}

func (a Application) StreamPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	device := vars["id"]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	link, _ := data.LinkModel.GetByStatusAndDevice(data.LinkModel{DB: a.DB}, true, device)
	err := helpers.WriteJSON(w, 200, link, nil)
	if err != nil {
		return
	}
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	answer := Answer{
		message: links.Name + "was updated",
		status:  "Ok",
	}
	err = helpers.WriteJSON(w, 200, answer, nil)
	if err != nil {
		return
	}
}

// AdminPageHandler
//
//	func (a Application) InsertPageHandler(w http.ResponseWriter, r *http.Request) {
//		tmpl, err := template.New("").ParseFiles("cmd/api/static/templates/insert.html", "cmd/api/static/templates/base.html")
//		if err != nil {
//			fmt.Println("error at insert html loader	")
//		}
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		err = tmpl.ExecuteTemplate(w, "base", nil)
//		if err != nil {
//			return
//		}
//	}
func (a Application) AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	links, err := data.LinkModel.GetAllLinks(data.LinkModel{DB: a.DB})
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = helpers.WriteJSON(w, 200, links, nil)
	if err != nil {
		return
	}
}

func (a Application) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	links, err := data.LinkModel.GetAllWorkingLinks(data.LinkModel{DB: a.DB})
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = helpers.WriteJSON(w, 200, links, nil)
	if err != nil {
		return
	}
}

func (a Application) SelectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	link, _ := data.LinkModel.GetByName(data.LinkModel{DB: a.DB}, id)
	err := data.LinkModel.UpdateAllLinks(data.LinkModel{DB: a.DB}, link.Device)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_ = data.LinkModel.UpdateLink(data.LinkModel{DB: a.DB}, id)
	http.Redirect(w, r, "/", 303)
}
