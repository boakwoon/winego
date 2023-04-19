package productcontroller

import (
	"github.com/boakwoon/winego/models/productmodel"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	products, _ := productmodel.FindByPath(w, r)
	data := map[string]interface{}{
		"products": products,
	}

	ts, err := template.ParseFiles("../static/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
