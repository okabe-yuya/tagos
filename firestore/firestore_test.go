package firestore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/okabe-yuya/tagos/types"
	"google.golang.org/api/iterator"
)

func TestMain(m *testing.M) {
	os.Setenv("EXEC_MODE", "test")
	ctx, client, err := GetClient()
	if err != nil {
		log.Fatalf("Failed from TestMain: %v", err)
		return
	}
	deleteCollection(ctx, client)
	client.Close()

	code := m.Run()
	os.Exit(code)
}

func deleteCollection(ctx context.Context, client *firestore.Client) error {
	col := collectionName()
	ref := client.Collection(col)
	for {
		iter := ref.Limit(100).Documents(ctx)
		numDeleted := 0
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed from getTagos: %v", err)
				return err
			}
			batch.Delete(doc.Ref)
			numDeleted++
		}

		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

func Test_GetTagos(t *testing.T) {
	ctx, client, err := GetClient()
	defer client.Close()
	if err != nil {
		t.Errorf("Failed: can't created context / %v", err)
	}

	creater := func (year, month, date int) error {
		data := types.TagosRecordInit("A", "B", year, month, date)
		if err := PostTagos(data, ctx, client); err != nil {
			log.Fatalf("Failed: %v", err)
			return err
		}
		return nil
	}

	expecter := func (mode string, year, month, date, expectHits int) error {
		data := types.TagosRecordInit("A", "B", year, month, date)
		resp, err := GetTagos(mode, data, ctx, client)
		if err != nil {
			log.Fatalf("Failed: %v", err)
			return err
		}

		if len(resp) != expectHits {
			errorText := fmt.Sprintf("Unmatch expectsHits size / expects=%v, response=%v", expectHits, len(resp))
			err := errors.New(errorText)
			return err
		}
		return nil
	}

	creater(2020, 10, 8)
	creater(2020, 10, 9)
	creater(2020, 11, 8)
	creater(2020, 11, 9)
	creater(2021, 10, 8)
	creater(2021, 10, 8)

	// yearで絞り込み
	if err := expecter("year", 2020, 10, 8, 4); err != nil {
		t.Errorf("Failed: GetTacos / %v", err)
	}

	if err := expecter("year", 2021, 10, 8, 2); err != nil {
		t.Errorf("Failed: GetTacos / %v", err)
	}

	// monthで絞り込み
	if err := expecter("month", 2020, 10, 8, 2); err != nil {
		t.Errorf("Failed: GetTacos / %v", err)
	}

	// dateで絞り込み
	if err := expecter("month", 2021, 10, 8, 2); err != nil {
		t.Errorf("Failed: GetTacos / %v", err)
	}
}

func Test_PostTagos(t *testing.T) {
	ctx, client, err := GetClient()
	defer client.Close()
	if err != nil {
		t.Errorf("Failed: can't created context / %v", err)
	}

	data := types.TagosRecordInit("A", "B", 2021, 10, 8)
	if err := PostTagos(data, ctx, client); err != nil {
		t.Errorf("Failed: can't post tagos / %v", err)
	}
}

