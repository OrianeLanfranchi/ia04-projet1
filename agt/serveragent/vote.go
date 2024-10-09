package serveragent

import (
	"encoding/json"
	"net/http"
)

func (rsa *ServerAgent) doVote(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	rsa.Lock()
	defer rsa.Unlock()

	w.WriteHeader(http.StatusOK)
	rsa.Lock()
	defer rsa.Unlock()
	serial, _ := json.Marshal(rsa.reqCount)
	w.Write(serial)
}
