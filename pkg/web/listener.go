package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	c "github.com/mislavzanic/container-updater/pkg/container"
	t "github.com/mislavzanic/container-updater/pkg/types"
)

type Listener struct {
	handler func(http.ResponseWriter, *http.Request)
}

func NewListener(client c.Client) Listener {
	secret := ""
	if val, present := os.LookupEnv("PAYLOAD_SECRET_FILE"); present {
		bytes, err := os.ReadFile(val)
		if err != nil {
			log.Fatalf("Error reading secret: %s", err)
		}
		secret = string(bytes)
	}
	
	handler := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var p t.Payload
		if err := decoder.Decode(&p); err != nil {
			log.Printf("Incorrect payload: %s", err)
			return
		}

		if strings.TrimSuffix(p.Secret, "\n") != strings.TrimSuffix(secret, "\n") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = "Unauthorized"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
			return
		}

		switch p.RequestType {
		case t.CreateUpdate: 
			id, err := client.CreateUpdateContainer(c.NewContainer(p));
			if err != nil {
				log.Fatalf("Error updating image: %s", err)
			}
			log.Printf("Started %s", id)
		case t.Update: 
			if err := client.UpdateContainer(c.NewContainer(p)); err != nil {
			    log.Fatalf("Error updating image: %s", err)
			}
		}
	}

	return Listener{
		handler: handler,
	}
}

func (l Listener) Run() {
	http.HandleFunc("/", l.handler)
	http.ListenAndServe(":8081", nil)
}
