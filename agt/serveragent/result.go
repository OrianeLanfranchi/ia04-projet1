package serveragent

import (
	"net/http"
)

func (rsa *ServerAgent) doResult(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	rsa.Lock()
	defer rsa.Unlock()
}
