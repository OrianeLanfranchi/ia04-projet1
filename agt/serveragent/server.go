package serveragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	rad "gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo"
)

type ServerAgent struct {
	sync.Mutex
	id       string
	reqCount int
	addr     string
}

func NewServerAgent(addr string) *ServerAgent {
	return &ServerAgent{id: addr, addr: addr}
}

// Test de la méthode
func (rsa *ServerAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

func (*ServerAgent) decodeRequest(r *http.Request) (req rad.Request, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *ServerAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("POST /new_ballot", rsa.doNewBallot)
	mux.HandleFunc("POST /vote", rsa.doVote)
	mux.HandleFunc("POST /result", rsa.doResult)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
