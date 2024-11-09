package restclientagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
)

type RestClientAgent struct {
	id      string
	url     string
	ballot  string
	nbAlts  int
	prefs   []int
	options []int
}

func NewRestClientAgent(id string, url string, ballot string, nbalt int, prefs []int, options []int) *RestClientAgent {
	return &RestClientAgent{id, url, ballot, nbalt, prefs, options}
}

func (rca *RestClientAgent) doRequest() (err error) {
	req := rad.VoteRequest{
		AgentId: rca.id,
		VoteId:  rca.ballot,
		Prefs:   rca.prefs,
		Option:  rca.options,
	}

	// sérialisation de la requête
	url := rca.url + "/vote"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	return
}

func (rca *RestClientAgent) Start() {
	log.Printf("démarrage de %s", rca.id)

	//par défaut
	rca.options = append(rca.options, rca.nbAlts)

	//Création aléatoire du profil de préférences à partir du nombre d'alternatives
	var orderedAlts = make([]int, rca.nbAlts)
	for i := range rca.nbAlts {
		orderedAlts[i] = (i + 1)
	}
	rand.Shuffle(rca.nbAlts, func(i, j int) { orderedAlts[i], orderedAlts[j] = orderedAlts[j], orderedAlts[i] })
	rca.prefs = orderedAlts

	err := rca.doRequest()

	if err != nil {
		log.Fatal(rca.id, "error:", err.Error())
	} else {
		log.Printf("[%s] voted for %s\n", rca.id, rca.ballot)
	}
}
