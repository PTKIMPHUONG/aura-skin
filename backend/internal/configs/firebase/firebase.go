package configs

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	FirebaseStorageBucketName = "auraskin-913fa.appspot.com"
)

func InitializeFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("./internal/configs/firebase/auraskin-913fa-firebase-adminsdk-gt6i2-5bf6c1bda2.json")

	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		StorageBucket: FirebaseStorageBucketName,
	}, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	return app
}
