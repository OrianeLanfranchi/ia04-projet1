package serveragent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
	cs "github.com/OrianeLanfranchi/ia04-projet1/comsoc"
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
		w.WriteHeader(http.StatusNotFound)
		msg := fmt.Sprintf("'%s' n'existe pas", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	//fmt.Println(time.Now().UTC())
	//fmt.Println(ballot.Deadline)
	if ballot.Deadline.After(time.Now().UTC()) {
		w.WriteHeader(http.StatusTooEarly)
		msg := fmt.Sprintf("'%s' n'est pas encore terminé", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	// traitement de la requête

	//si on a déjà traité les résultats du ballot :
	if ballot.Result.Winner != -1 {
		w.WriteHeader(http.StatusOK)
		serial, _ := json.Marshal(ballot.Result)
		w.Write(serial)
		return
	}

	//Sinon on calcule les résultats :
	if ballot.Options == nil && ballot.SCF.FuncNoOption == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		msg := fmt.Sprintf("'%s' est mal formé et aucun résultat ne peut être déduit", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	//Calcul du gagnant
	winner, errSCF := ballot.SCF.Call(ballot.Profile, ballot.Options)

	if errSCF != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("'%s' ne peut pas être traité ou bien il n'y a pas de gagnant", req.BallotId)
		w.Write([]byte(msg))
		return
	}

	//Calcul du ranking
	//On vérifie que la règle de vote comporte bien une SWF (c'est pas propre, je sais)
	if ballot.SWF.FuncNoOption != nil || ballot.SWF.FuncOneOption != nil {
		ranking, errSWF := ballot.SWF.Call(ballot.Profile, ballot.Options)

		if errSWF != nil && winner == cs.Alternative(-1) {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("'%s' ne peut pas être traité", req.BallotId)
			w.Write([]byte(msg))
			return
		}

		fmt.Println("ranking :", ranking)
		ballot.Result.Ranking = make([]int, len(ranking))
		for i := range ranking {
			ballot.Result.Ranking[i] = int(ranking[i])
		}
	}

	//Sinon comme Condorcet est la seule règle du projet pour laquelle il n'y a pas de SWF, on pourrait écrire la condition ainsi : (ce qui n'est toujours pas propre)
	/*if ballot.Rule != "condorcet" {
		ranking, errSWF := ballot.SWF.Call(ballot.Profile, ballot.Options)

		if errSWF != nil && winner == cs.Alternative(-1) {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("'%s' ne peut pas être traité", req.BallotId)
			w.Write([]byte(msg))
			return
		}

		fmt.Println("ranking :", ranking)
		ballot.Result.Ranking = make([]int, len(ranking))
		for i := range ranking {
			ballot.Result.Ranking[i] = int(ranking[i])
		}
	}*/

	ballot.Result.Winner = int(winner)
	rsa.ballots[req.BallotId] = ballot

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(ballot.Result)
	w.Write(serial)
}
