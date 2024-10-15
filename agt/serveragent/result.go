package serveragent

import (
	"fmt"
	"net/http"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
	//cs "github.com/OrianeLanfranchi/ia04-projet1/comsoc"
)

func (rsa *ServerAgent) doResult(w http.ResponseWriter, r *http.Request) {
	rsa.Lock()
	defer rsa.Unlock()

	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := decodeRequest[rad.ResultRequest](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	ballot, ok := rsa.ballots[req.BallotId]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' n'existe pas", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	if ballot.Deadline.Before(time.Now()) {
		w.WriteHeader(http.StatusTooEarly)
		msg := fmt.Sprintf("'%s' n'est pas encore terminé", req.BallotId)
		w.Write([]byte(msg))
		return
	}

}
