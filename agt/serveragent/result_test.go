package serveragent

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
)

func ExecuteTestResultRequest(req rad.ResultRequest, handler *http.ServeMux) *httptest.ResponseRecorder {
	data, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/result", bytes.NewBuffer(data))

	handler.ServeHTTP(w, request)

	return w
}

func TestDoResult(t *testing.T) {
	_, handler := LaunchTestServer()
	deadline := time.Now().Add(time.Second * 3)

	reqBallot := rad.BallotRequest{
		Rule:     `borda`,
		Deadline: deadline.Format(time.UnixDate),
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	ExecuteTestNewBallotRequest(reqBallot, handler)

	reqVote := rad.VoteRequest{
		AgentId: "1",
		VoteId:  "vote1",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}
	ExecuteTestVoteRequest(reqVote, handler)

	reqVote2 := rad.VoteRequest{
		AgentId: "2",
		VoteId:  "vote1",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}
	ExecuteTestVoteRequest(reqVote2, handler)

	reqVote3 := rad.VoteRequest{
		AgentId: "3",
		VoteId:  "vote1",
		Prefs:   []int{2, 1},
		Option:  []int{},
	}
	ExecuteTestVoteRequest(reqVote3, handler)

	time.Sleep(time.Second * 4)
	reqResult := rad.ResultRequest{
		BallotId: "vote1",
	}
	w := ExecuteTestResultRequest(reqResult, handler)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	} else {
		//else check ballot results

		//check json object integrity
		var respResult rad.ResultResponse
		body, err1 := io.ReadAll(w.Result().Body)
		if err1 != nil {
			t.Error(err1)
		}
		err2 := json.Unmarshal(body, &respResult)
		if err1 != nil {
			t.Error(err2)
		}

		//check vote results
		if respResult.Winner != 1 {
			t.Error("wrong winner")
		}

		if !reflect.DeepEqual(respResult.Ranking, []int{1, 2}) {
			t.Error("wrong ranking")
		}

	}

}

func TestDoResultVoteNotFinished(t *testing.T) {
	_, handler := LaunchTestServer()
	deadline := time.Now().Add(time.Hour * 2)

	reqBallot := rad.BallotRequest{
		Rule:     `borda`,
		Deadline: deadline.Format(time.UnixDate),
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	ExecuteTestNewBallotRequest(reqBallot, handler)

	reqVote := rad.VoteRequest{
		AgentId: "1",
		VoteId:  "vote1",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}
	ExecuteTestVoteRequest(reqVote, handler)

	reqVote2 := rad.VoteRequest{
		AgentId: "2",
		VoteId:  "vote1",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}
	ExecuteTestVoteRequest(reqVote2, handler)

	reqVote3 := rad.VoteRequest{
		AgentId: "3",
		VoteId:  "vote1",
		Prefs:   []int{2, 1},
		Option:  []int{},
	}
	ExecuteTestVoteRequest(reqVote3, handler)

	reqResult := rad.ResultRequest{
		BallotId: "vote1",
	}
	w := ExecuteTestResultRequest(reqResult, handler)

	if w.Result().StatusCode != http.StatusTooEarly {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoResultVoteDoesntExist(t *testing.T) {
	_, handler := LaunchTestServer()

	reqResult := rad.ResultRequest{
		BallotId: "vote1",
	}
	w := ExecuteTestResultRequest(reqResult, handler)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}

}
