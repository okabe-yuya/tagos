package firestore

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	// "github.com/n2i-inc/tagos/constants"
)

func getClient() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config { ProjectID: "aso-viva" }
	client, err := func() (*firestore.Client, error) {
		env := os.Getenv("ENV_NAME")
		if env == "production" {
			app, err := firebase.NewApp(ctx, conf)
			if err != nil {
				log.Fatalf("Failed from getClient: %v", err)
				return nil, err
			}
			client, err := app.Firestore(ctx)
			if err != nil {
				log.Fatalf("Failed from getClient: %v", err)
				return nil, err
			}
			return client, nil
		}

		sa := option.WithCredentialsFile("./../serviceAccount.json")
		app, err := firebase.NewApp(ctx, conf, sa)
		if err != nil {
			log.Fatalf("Failed from getClient: %v", err)
			return nil, err
		}
		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalf("Failed from getClient: %v", err)
			return nil, err
		}
		return client, nil
	}()

	if err != nil {
		return nil, nil, err
	}
	return ctx, client, err
}

func getTagos(ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	tagoses := make([]map[string]interface{}, 0)
	iter := client.Collection("tagos").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed from getTagos: %v", err)
			return make([]map[string]interface{}, 0), err
		}
		tagoses = append(tagoses, doc.Data())
	}
	return tagoses, nil
}

func monthTagos(year, month string, ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	tagoses := make([]map[string]interface{}, 0)
	iter := client.Collection("tagos").Doc(year).Collections(ctx)
	for {
		col, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed from getTagos: %v", err)
			return make([]map[string]interface{}, 0), err
		}
		tagoses = append(tagoses, col.Add()
	}
}

func postTagos(ctx context.Context, client *firestore.Client) error {
	if _, _, err := client.Collection("tagos").Doc("2021").Collection("04").Doc("15").Collection("vote").Add(ctx, map[string]interface{}{
		"sender": "A",
		"receiver": "B",
		"created_at": "2021/03/23",
	}); err != nil {
		log.Fatalf("Failed from postTacos: %v", err)
		return err
	}
	return nil
}