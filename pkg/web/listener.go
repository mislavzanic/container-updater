package web

import (
	"encoding/json"
	"log"
	"net/http"

	t "github.com/mislavzanic/container-updater/pkg/types"
	c "github.com/mislavzanic/container-updater/pkg/container"
)

type Listener struct {
	handler func(http.ResponseWriter, *http.Request)
}

func NewListener(client c.Client) Listener {
	handler := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var p t.Payload
		if err := decoder.Decode(&p); err != nil {
			log.Fatal(err)
		}

		id, err := client.CreateUpdateContainer(c.NewContainer(p));
		if err != nil {
			log.Fatalf("Error updating image: %s", err)
		}
		log.Printf("Started %s", id)

	}

	return Listener{
		handler: handler,
	}
}

func (l Listener) Run() {
	http.HandleFunc("/", l.handler)
	http.ListenAndServe(":8081", nil)
}
