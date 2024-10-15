package serveragent

import (
	"encoding/json"
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

	// traitement de la requête

	if ballot.Result.Winner != -1 {
		w.WriteHeader(http.StatusOK)
		serial, _ := json.Marshal(ballot.Result)
		w.Write(serial)
		return
	}

	// else on le calcule directement
	//TODO : Ranking (pour l'instant je le laisse à 0 parce que flemme)

	//DEBUG - Bien vérifier le formattage des options. Logiquement on ne traite que les cas où il y a 0 ou 1 option. Faire un système scalable si on devait update le serveur pour qu'il prenne en compte des votes avec + d'options
	if ballot.Options == nil && ballot.SCF.FuncNoOption == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		msg := fmt.Sprintf("'%s' est mal formé et aucun résultat ne peut être déduit", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	// C'est pas le truc le plus propre mais ça fonctionne (j'espère)
	//DEBUG - peut-être que ça peut planter ici on sait pas

	winner, errSCF := ballot.SCF.Call(ballot.Profile, ballot.Options)

	if errSCF != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("'%s' ne peut pas être traité", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	ballot.Result.Winner = int(winner)

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(ballot.Result)
	w.Write(serial)
}
