package api

import (
	"L0/pkg/cache"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)
type Api struct {
	Router *mux.Router
	cache  *cache.Cache
}


// InitRouter инициализирует роутер, кэш и методы
func (a *Api) InitRouter(cache *cache.Cache) {
	a.cache = cache
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/order/", a.getOrder)
}

// getOrder хэндлер для "/order/"
func (a *Api) getOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
		id:=r.URL.Query().Get("id")
		order, err := a.cache.Get(id)
		if err != nil {
			h:="<body bgcolor=\"#FFE4B5\">\n<h1><font size=\"16\" face=\"serif\">{{ .}}</font></h1>\n</body>"
			tmpl, _ := template.New(err.Error()).Parse(h)
			w.WriteHeader(http.StatusNotFound)
			tmpl.Execute(w, err.Error())
			return
		}

		tmpl, _ := template.ParseFiles("templates/index.html")
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, order)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

}