package serveragent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
)

func (rsa *ServerAgent) doNewBallot(w http.ResponseWriter, r *http.Request) {
	// mise à jour du nombre de requêtes
	rsa.Lock()
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := decodeRequest[rad.BallotRequest](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// vérirification de la date
	deadline, errDate := time.Parse(time.UnixDate, req.Deadline)

	if errDate != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' n'est pas au format UnixDate", req.Deadline)
		w.Write([]byte(msg))
		return
	}

	if deadline.Before(time.Now()) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' est une date antérieure à la date présente", req.Deadline)
		w.Write([]byte(msg))
		return
	}

	// traitement de la requête
	var resp rad.BallotResponse

	resp.ID = fmt.Sprintf("ballot%d", len(rsa.ballots)+1)
	rsa.ballots = append(rsa.ballots, resp)

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
