package p

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"

	"github.com/okabe-yuya/tagos/aggregate"
	"github.com/okabe-yuya/tagos/firestore"
	"github.com/okabe-yuya/tagos/types"
	"github.com/slack-go/slack"
)

func GetTagosHttpServer(w http.ResponseWriter, r *http.Request) {
	args, err := slack.SlashCommandParse(r)
  if err != nil {
		log.Fatalf("Failed from GetTagosHttpServer: %v", err)
    w.WriteHeader(http.StatusBadRequest)
		return
  }

	resp := aggregate.ToResponse(getTagos(args.Text))
	params := &slack.Msg{Text: resp, ResponseType: "in_channel"}
	b, err := json.Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func PostTagosHttpServer(w http.ResponseWriter, r *http.Request) {
	args, err := slack.SlashCommandParse(r)
  if err != nil {
		log.Fatalf("Failed from PostTagosHttpServer: %v", err)
    w.WriteHeader(http.StatusBadRequest)
		return
  }

	slice := strings.Split(args.Text, " ")
	if len(slice) > 1 {
		if err := postTagos(args.UserName, slice[0]); err != nil {
			log.Fatalf("Failed from PostTagosHttpServer: %v", err)
			w.Write([]byte("おや！何か手違いがありました！ もう一度試してください😭"))
			w.WriteHeader(http.StatusBadRequest)
		} else {
			response := fmt.Sprintf("%vさんから%vさんに🌮が届きました！\n「%v」", args.UserName, slice[0], slice[1])
			params := &slack.Msg{Text: response, ResponseType: "in_channel"}
			b, err := json.Marshal(params)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			w.WriteHeader(http.StatusCreated)
		}
		return
	}
	w.Write([]byte("おや！何か手違いがありました！ もう一度試してください😭"))
	w.WriteHeader(http.StatusBadRequest)
}

func getTagos(mode string) []map[string]interface{} {
	now := time.Now().Add(time.Hour * 9)
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

func postTagos(sender, receiver string) error {
	ctx, client, err := firestore.GetClient()
	if err != nil {
		log.Fatalf("Failed from postTagos: %v", err)
    return err
	}
	now := time.Now().Add(time.Hour * 9)
	data := types.TagosRecordInit(sender, receiver, now.Year(), int(now.Month()), now.Day())
	if err := firestore.PostTagos(data, ctx, client); err != nil {
		log.Fatalf("Failed from postTagos: %v", err)
    return err
	}
	return nil
}