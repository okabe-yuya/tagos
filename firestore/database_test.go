package firestore

import (
	// "fmt"
	"fmt"
	"testing"
)

func Test_postTagos(t *testing.T) {
	ctx, client, err := getClient()
	defer client.Close()
	if err != nil {
		t.Errorf("Failed: can't created context / %v", err)
	}

	if err := postTagos(ctx, client); err != nil {
		t.Errorf("Failed: can't post tagos / %v", err)
	}
}

func Test_getTagos(t *testing.T) {
	ctx, client, err := getClient()
	defer client.Close()
	if err != nil {
		t.Errorf("Failed: can't created context / %v", err)
	}

	tagos, err := getTagos(ctx, client)
	if err != nil {
		t.Errorf("Failed: can't post tagos / %v", err)
	}

	fmt.Println(tagos)
}