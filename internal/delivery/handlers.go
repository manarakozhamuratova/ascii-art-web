package delivery

import (
	"ascii/internal/service"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const (
	bannerStandard   = "standard"
	bannerShadow     = "shadow"
	bannerThinkertoy = "thinkertoy"
)

type Data struct {
	Art string
}

type Error struct {
	StatusCode int
	StatusText string
}

func handleError(w http.ResponseWriter, statusError int) {
	tmp, err := template.ParseFiles("./static/html/errors.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	data := Error{
		StatusCode: statusError,
		StatusText: http.StatusText(statusError),
	}
	w.WriteHeader(statusError)
	tmp.Execute(w, data)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		handleError(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("./static/html/home.html")
	if err != nil {
		log.Println(err.Error())
		handleError(w, http.StatusInternalServerError)
		return
	}
	if err := tmp.ExecuteTemplate(w, "home.html", nil); err != nil {
		log.Println(err.Error())
		handleError(w, http.StatusInternalServerError)
		return
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, http.StatusMethodNotAllowed)
		return
	}
	input := &input{
		text:   r.FormValue("comment"),
		banner: r.FormValue("banner"),
	}
	if !input.validate() {
		handleError(w, http.StatusBadRequest)
		return
	}
	if service.IsValid(input.text) {
		handleError(w, http.StatusBadRequest)
		return
	}
	if !input.checkBanner(bannerShadow, bannerStandard, bannerThinkertoy) {
		handleError(w, http.StatusBadRequest)
		return
	}
	res, err := service.App(input.text, input.banner)
	if err != nil {
		log.Printf("error from service.App : %s", err.Error())
		handleError(w, http.StatusInternalServerError)
		return
	}
	download := r.FormValue("download")
	if download == "DOWNLOAD" {
		w.Header().Set("Content-Disposition", "attachment; name=text; filename=ascii.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(res)))
		_, err := w.Write([]byte(res))
		if err != nil {
			handleError(w, http.StatusInternalServerError)
			return
		}
	}

	tmp, err := template.ParseFiles("./static/html/home.html")
	if err != nil {
		return
	}
	data := Data{
		Art: res,
	}
	err1 := tmp.ExecuteTemplate(w, "home.html", data)
	if err1 != nil {
		return
	}
}
