package main

import "github.com/your-username/your-wine-selling-website/recommender"

func main() {
	// Create a new recommender system
	r := &recommender.CollaborativeFiltering{}

	// Train the recommender system on past purchases
	r.Train(pastPurchases)

	// Get recommendations for a user
	recommendations, err := r.GetRecommendations(userID)
	if err != nil {
		panic(err)
	}

	// Display the recommendations to the user
	for _, recommendation := range recommendations {
		fmt.Printf("You might like %s - %s\n", recommendation.Name, recommendation.Description)
	}
}
