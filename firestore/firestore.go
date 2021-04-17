package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/okabe-yuya/tagos/constants"
	"github.com/okabe-yuya/tagos/types"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GetClient() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config { ProjectID: constants.PROJECT_NAME }
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

func GetTagos(mode string, data *types.TagosRecord, ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	switch mode {
	case "year":
		return yearTagos(data, ctx, client)
	case "month":
		return monthTagos(data, ctx, client)
	case "date":
		return dateTagos(data, ctx, client)
	}
	return emptyTagos(), nil
}

func PostTagos(data *types.TagosRecord, ctx context.Context, client *firestore.Client) error {
	col := collectionName()
	if _, _, err := client.Collection(col).Add(ctx, map[string]interface{}{
		"sender": data.Sender,
		"receiver": data.Receiver,
		"year": data.Year,
		"month": data.Month,
		"date": data.Date,
		"created_at": data.CreatedAt,
	}); err != nil {
		log.Fatalf("Failed from postTacos: %v", err)
		return err
	}
	return nil
}


// private functions
func collectionName() string {
	mode := os.Getenv("EXEC_MODE")
	if mode == "test" {
		return "tagos_test"
	}
	return "tagos"
}

func emptyTagos() []map[string]interface{} {
	return make([]map[string]interface{}, 0)
}

func tagosIterator(iter *firestore.DocumentIterator) ([]map[string]interface{}, error) {
	resp := emptyTagos()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed from getTagos: %v", err)
			return emptyTagos(), err
		}
		resp = append(resp, doc.Data())
	}
	return resp, nil
}

func yearTagos(data *types.TagosRecord, ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	col := collectionName()
	iter := client.Collection(col).Where("year", "==", data.Year).Documents(ctx)
	resp, err := tagosIterator(iter)
	return resp, err
}

func monthTagos(data *types.TagosRecord, ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	col := collectionName()
	iter := client.Collection(col).Where("year", "==", data.Year).Where("month", "==", data.Month).Documents(ctx)
	resp, err := tagosIterator(iter)
	return resp, err
}

func dateTagos(data *types.TagosRecord, ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	col := collectionName()
	iter := client.Collection(col).Where("year", "==", data.Year).Where("month", "==", data.Month).Where("date", "==", data.Month).Documents(ctx)
	resp, err := tagosIterator(iter)
	return resp, err
}
