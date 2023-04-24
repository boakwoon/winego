package app

import "github.com/jinzhu/gorm"

type Fav struct { // структура избранного
	gorm.Model
	UserId    uint  `json:"user_id"`
	ProductId int64 `json:"product_id"`
}

func rateItem(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values
	itemID := r.FormValue("item_id")
	rating := r.FormValue("rating")

	// Execute SQL query to update rating
	_, err := db.Exec("UPDATE items SET rating = ? WHERE id = ?", rating, itemID)
	if err != nil {
		// Handle error
	}
}
