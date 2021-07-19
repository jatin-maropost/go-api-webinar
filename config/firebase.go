package config

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Configure and initialise the firebase admin sdk
func ConfigureFirestore() *firestore.Client  {

	// Load firestore configuration keys
	opt := option.WithCredentialsFile("./firestore-config.json")
	
	// Initialise app
	app, err := firebase.NewApp(context.Background(), nil, opt)

	// Log error if error occurs
	if err != nil {
	  log.Fatal(fmt.Errorf("error initializing app: %v", err))
	}

	// Get a Firestore client.
	client,err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing app: %v", err))
	  }

	  // Intialise the firestore client in global variable
	  return client
}
