package p

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/okabe-yuya/tagos/aggregate"
	"github.com/okabe-yuya/tagos/firestore"
	"github.com/okabe-yuya/tagos/types"
)


func TagosHttpServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fetched := aggregate.AggregateReceiver(getTagos(r))
		resp, err := json.Marshal(types.GetRespInit(fetched))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	case "POST":
		if err := postTagos(r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTagos(r *http.Request) []map[string]interface{} {
	mode := r.URL.Query().Get("mode")
	now := time.Now()
	empty := make([]map[string]interface{}, 0)
	data := types.TagosRecordInit("", "", now.Year(), int(now.Month()), now.Day())
	ctx, client, err := firestore.GetClient()
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed from getTagos: %v", err)
		return empty
	}

	resp, err := firestore.GetTagos(mode, data, ctx, client)
	if err != nil {
		log.Fatalf("Failed from getTagos: %v", err)
		return empty
	}
	return resp
}

func postTagos(r *http.Request) error {
	var body *types.Body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    log.Fatalf("Failed from postTagos: %v", err)
    return err
  }

	ctx, client, err := firestore.GetClient()
	if err != nil {
		log.Fatalf("Failed from postTagos: %v", err)
    return err
	}
	now := time.Now()
	data := types.TagosRecordInit(body.Sender, body.Reveiver, now.Year(), int(now.Month()), now.Day())
	if err := firestore.PostTagos(data, ctx, client); err != nil {
		log.Fatalf("Failed from postTagos: %v", err)
    return err
	}
	return nil
}