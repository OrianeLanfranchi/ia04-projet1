package serveragent

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
	cs "github.com/OrianeLanfranchi/ia04-projet1/comsoc"
)

func removeFromSliceString(oldSlice []string, element string) []string {
	newSlice := make([]string, len(oldSlice))

	copy(newSlice, oldSlice)

	for i := range oldSlice {
		if newSlice[i] == element {
			newSlice = append(newSlice[:i], newSlice[i+1:]...)
			return newSlice
		}
	}
	return oldSlice // element was not found
}

func (rsa *ServerAgent) doVote(w http.ResponseWriter, r *http.Request) {
	rsa.Lock()
	defer rsa.Unlock()

	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := decodeRequest[rad.VoteRequest](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	ballot, ok := rsa.ballots[req.VoteId]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' n'existe pas", req.VoteId)
		w.Write([]byte(msg))
		return
	}

	if ballot.Deadline.Before(time.Now()) {
		w.WriteHeader(http.StatusTooEarly)
		msg := fmt.Sprintf("'%s' est terminé ; les votes ne sont plus acceptés", req.VoteId)
		w.Write([]byte(msg))
		return
	}

	if !slices.Contains(ballot.VotersId, req.AgentId) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' ne concerne pas le votant %s, ou bien il a déjà voté", req.VoteId, req.AgentId)
		w.Write([]byte(msg))
		return
	}

	ballot.VotersId = removeFromSliceString(ballot.VotersId, req.AgentId)

	//Vérification sur la taille des préférences (au moins une préférence, pas plus de préférences qu'il n'y a d'alternatives)
	if (len(req.Prefs) == 0) || (len(req.Prefs) > ballot.NbAlts) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' - Les préférences du votant %s ne sont pas bien formattées (trop ou pas assez d'alternatives)", req.VoteId, req.AgentId)
		w.Write([]byte(msg))
		return
	}

	//Vérification que les valeurs des préférences ne sont pas aberrantes
	prefs := make([]cs.Alternative, 0)

	for i := range req.Prefs {
		if (req.Prefs[i] > ballot.NbAlts) || (req.Prefs[i] < 1) {
			fmt.Println(req.Prefs[i])
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("'%s' - Les préférences du votant %s ne sont pas bien formattées", req.VoteId, req.AgentId)
			w.Write([]byte(msg))
			return
		}
		prefs = append(prefs, cs.Alternative(req.Prefs[i]))
	}

	//On vérifie que le profil est bien construit
	if ballot.Rule != "approval" { //Comme approval ne nécessite pas que le profil de préférences soit complet, on discrimine ce cas (c'est pas propre, je sais)
		alternativesTemp := make([]cs.Alternative, ballot.NbAlts)
		for i := range ballot.NbAlts {
			alternativesTemp[i] = cs.Alternative(i + 1)
			fmt.Println(alternativesTemp[i])
		}

		errCheck := cs.CheckProfile(prefs, alternativesTemp)

		fmt.Println(errCheck)

		if errCheck != nil {
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("'%s' - Les préférences du votant %s ne sont pas bien formattées", req.VoteId, req.AgentId)
			w.Write([]byte(msg))
			return
		}
	}

	ballot.Profile = append(ballot.Profile, prefs)

	if req.Option != nil {
		ballot.Options = append(ballot.Options, req.Option)
	}

	// Vérification des options de vote
	if len(req.Option) == 0 && ballot.Rule == "approval" { //C'est tjrs pas propre, je sais
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("'%s' - Les options pour le vote du votant %s ne sont pas bien formattées", req.VoteId, req.AgentId)
		w.Write([]byte(msg))
		return
	}

	rsa.ballots[req.VoteId] = ballot

	w.WriteHeader(http.StatusOK)
	/*serial, _ := json.Marshal(ballot.Result)
	w.Write(serial)*/
}
