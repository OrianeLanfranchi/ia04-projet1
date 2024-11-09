package serveragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
)

type ServerAgent struct {
	sync.Mutex
	id      string
	addr    string
	ballots map[string]rad.Ballot
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

// do a decoderequest factory
func decodeRequest[Req rad.Request](r *http.Request) (req Req, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *ServerAgent) Start() {

	mux := rsa.SetUpHandlers()

	//init map
	rsa.InitBallots()
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

// création du multiplexer
func (rsa *ServerAgent) SetUpHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /new_ballot", rsa.doNewBallot)
	mux.HandleFunc("POST /vote", rsa.doVote)
	mux.HandleFunc("POST /result", rsa.doResult)
	return mux
}

func (rsa *ServerAgent) InitBallots() {
	rsa.ballots = make(map[string]rad.Ballot)
}
