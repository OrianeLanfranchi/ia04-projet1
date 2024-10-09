package serveragent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	comsoc "github.com/OrianeLanfranchi/ia04-projet1/comsoc"
	rad "gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo"
)

type Ballot struct {
	ballotId string
	profile  comsoc.Profile
	deadline time.Time
	votersId []string
	nbAlts   int
}

func (rsa *ServerAgent) doNewBallot(w http.ResponseWriter, r *http.Request) {
	// mise à jour du nombre de requêtes
	rsa.Lock()
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	var resp rad.Response

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
