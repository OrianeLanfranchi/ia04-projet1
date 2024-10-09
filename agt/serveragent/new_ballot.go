package serveragent

import (
	"encoding/json"
	"fmt"
	"net/http"

	rad "gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo"
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
	req, err := rsa.decodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	var resp rad.Response

	switch req.Operator {
	case "*":
		resp.Result = req.Args[0] * req.Args[1]
	case "+":
		resp.Result = req.Args[0] + req.Args[1]
	case "-":
		resp.Result = req.Args[0] - req.Args[1]
	case "/":
		if req.Args[1] == 0 {
			msg := fmt.Sprintf("Can't be divided by 0")
			w.Write([]byte(msg))
		} else {
			resp.Result = req.Args[0] / req.Args[1]
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Unkonwn command '%s'", req.Operator)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
