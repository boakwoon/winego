package productcontroller

import (
	"github.com/boakwoon/winego/models/productmodel"
	"github.com/upper/db/v4"
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

func filterItems(w http.ResponseWriter, r *http.Request) {
	// получение формы
	minPrice := r.FormValue("min_price")
	maxPrice := r.FormValue("max_price")
	minRating := r.FormValue("min_rating")
	maxRating := r.FormValue("max_rating")

	// Execute SQL query with form values
	items, err := db.Query("SELECT * FROM items WHERE price >= ? AND price <= ? AND rating >= ? AND rating <= ?", minPrice, maxPrice, minRating, maxRating)
	if err != nil {
		// обработка ошибки
	}

	// отображение элементов
	// ...
}

func commentItem(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values
	itemID := r.FormValue("item_id")
	commentText := r.FormValue("comment_text")

	// Execute SQL query to insert comment
	_, err := db.Exec("INSERT INTO comments (item_id, comment_text) VALUES (?, ?)", itemID, commentText)
	if err != nil {
		// Handle error
	}

	// Redirect to item detail page
	http.Redirect(w, r, "/item?id="+itemID, http.StatusFound)
}
