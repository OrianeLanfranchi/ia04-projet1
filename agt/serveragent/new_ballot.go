package serveragent

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
	cs "github.com/OrianeLanfranchi/ia04-projet1/comsoc"
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

	var ballot rad.Ballot = rad.Ballot{
		Profile:  make(cs.Profile, 0),
		Options:  make([][]int, 0),
		VotersId: make([]string, 0),
		NbAlts:   req.NbAlts,
		Deadline: deadline,
		Result:   rad.ResultResponse{}}

	var orderedAlts = make([]cs.Alternative, req.NbAlts)

	for i := range req.NbAlts {
		orderedAlts[i] = cs.Alternative(i + 1)
	}
	rand.Shuffle(req.NbAlts, func(i, j int) { orderedAlts[i], orderedAlts[j] = orderedAlts[j], orderedAlts[i] })

	tieBreak := cs.TieBreakFactory(orderedAlts)

	var status = http.StatusOK

	switch req.Rule {
	case "majority":
		ballot.SCF = cs.SCFFactory(cs.MajoritySCF, tieBreak)
	case "borda":
		ballot.SCF = cs.SCFFactory(cs.BordaSCF, tieBreak)
	case "condorcet":
		ballot.SCF = cs.SCFFactory(cs.CondorcetWinner, tieBreak)
	case "approval":
		//TODO : SCFFactory with options
		//ballot.SCF = cs.SCFOptionFactory(cs.ApprovalSCF, tieBreak)
		status = http.StatusNotImplemented
	default:
		status = http.StatusNotImplemented
	}

	rsa.ballots[resp.ID] = ballot

	//DEBUG - Si on a un pb c'est parce qu'on a lancé une goroutine sur un mutex locked
	// TODO - lancer une goroutine qui handle le ballot (ballotHandler pour le nom ?)

	w.WriteHeader(status)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
