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

	resp.ID = fmt.Sprintf("vote%d", len(rsa.ballots)+1)

	var ballot rad.Ballot = rad.Ballot{
		Rule:     req.Rule,
		Profile:  make(cs.Profile, 0),
		Options:  make([][]int, 0),
		VotersId: make([]string, 0),
		NbAlts:   req.NbAlts,
		Deadline: deadline,
		Result:   rad.ResultResponse{Winner: -1, Ranking: make([]int, 0)}}

	var orderedAlts = make([]cs.Alternative, req.NbAlts)

	for i := range req.NbAlts {
		orderedAlts[i] = cs.Alternative(i + 1)
	}
	rand.Shuffle(req.NbAlts, func(i, j int) { orderedAlts[i], orderedAlts[j] = orderedAlts[j], orderedAlts[i] })

	tieBreak := cs.TieBreakFactory(orderedAlts)

	switch req.Rule {
	case "majority":
		ballot.SCF.FuncNoOption = cs.SCFFactory(cs.MajoritySCF, tieBreak)
		ballot.SWF.FuncNoOption = cs.SWFFactory(cs.MajoritySWF, tieBreak)
	case "borda":
		ballot.SCF.FuncNoOption = cs.SCFFactory(cs.BordaSCF, tieBreak)
		ballot.SWF.FuncNoOption = cs.SWFFactory(cs.BordaSWF, tieBreak)
	case "condorcet":
		ballot.SCF.FuncNoOption = cs.SCFFactory(cs.CondorcetWinner, tieBreak)
	case "approval":
		//DEBUG if it crashes then it might be due to this
		ballot.SCF.FuncOneOption = cs.SCFOptionFactory(cs.ApprovalSCF, tieBreak)
		ballot.SWF.FuncOneOption = cs.SWFOptionFactory(cs.ApprovalSWF, tieBreak)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("'%s' n'est pas une méthode implémentée", req.Rule)
		w.Write([]byte(msg))
		return
	}

	rsa.ballots[resp.ID] = ballot

	//DEBUG - Si on a un pb c'est parce qu'on a lancé une goroutine sur un mutex locked
	// TODO - lancer une goroutine qui handle le ballot (ballotHandler pour le nom ?)

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
