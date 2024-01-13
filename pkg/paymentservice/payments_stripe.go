package paymentservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/sub"
)

var stripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")

func createSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request
	var requestParams map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(requestParams["email"].(string)),
	}
	newCustomer, err := customer.New(customerParams)
	if err != nil {
		http.Error(w, "Failed to create customer", http.StatusInternalServerError)
		return
	}

	// Create a subscription for the customer
	subscriptionParams := &stripe.SubscriptionParams{
		Customer: stripe.String(newCustomer.ID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(requestParams["priceId"].(string)),
			},
		},
	}
	newSubscription, err := sub.New(subscriptionParams)
	if err != nil {
		http.Error(w, "Failed to create subscription", http.StatusInternalServerError)
		return
	}

	// Retrun subscription details as JSON response
	response := map[string]interface{}{
		"customerId": newCustomer.ID,
		"subscriptionId": newSubscription.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}