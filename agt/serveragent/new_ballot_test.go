package serveragent

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
)

func LaunchTestServer() (server *ServerAgent, handler *http.ServeMux) {
	server = NewServerAgent(":8080")
	server.InitBallots()
	handler = server.SetUpHandlers()
	return
}

func ExecuteTestNewBallotRequest(req rad.BallotRequest, handler *http.ServeMux) *httptest.ResponseRecorder {
	data, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/new_ballot", bytes.NewBuffer(data))

	handler.ServeHTTP(w, request)

	return w
}

func TestDoNewBallot(t *testing.T) {
	_, handler := LaunchTestServer()

	req := rad.BallotRequest{
		Rule:     `borda`,
		Deadline: "Tue Nov 10 23:00:00 UTC 2024",
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	w := ExecuteTestNewBallotRequest(req, handler)
	resp := w.Result()

	//if server error :
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("server status code %d", resp.StatusCode)
		t.Errorf("%s", w.Body)
	} else {
		//else check ballot creation
		var respBallot rad.BallotResponse
		body, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			t.Error(err1)
		}
		err2 := json.Unmarshal(body, &respBallot)
		if err1 != nil {
			t.Error(err2)
		}
		if strings.HasPrefix("vote", respBallot.ID) {
			t.Error("wrong response")
		}
	}
}

func TestDoNewBallotDeadlineBefore(t *testing.T) {
	_, handler := LaunchTestServer()

	req := rad.BallotRequest{
		Rule:     `borda`,
		Deadline: "Tue Nov 10 23:00:00 UTC 2023",
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	w := ExecuteTestNewBallotRequest(req, handler)
	resp := w.Result()
	//if server error :
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("server status code : %d, expected server status : Bad Request", resp.StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoNewBallotRuleNotImplemented(t *testing.T) {
	_, handler := LaunchTestServer()

	req := rad.BallotRequest{
		Rule:     `singlePeak`,
		Deadline: "Tue Nov 10 23:00:00 UTC 2024",
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	w := ExecuteTestNewBallotRequest(req, handler)
	resp := w.Result()
	//if server error :
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("server status code : %d, expected server status : Not Implemented", resp.StatusCode)
		t.Errorf("%s", w.Body)
	}
}
