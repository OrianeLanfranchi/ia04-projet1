package serveragent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	rad "github.com/OrianeLanfranchi/ia04-projet1/agt"
)

func ExecuteTestVoteRequest(req rad.VoteRequest, handler *http.ServeMux) *httptest.ResponseRecorder {
	data, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/vote", bytes.NewBuffer(data))

	handler.ServeHTTP(w, request)

	return w
}

func TestDoVote(t *testing.T) {
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

	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoVotePastDeadline(t *testing.T) {
	_, handler := LaunchTestServer()
	deadline := time.Now().Add(time.Second * 2)

	reqBallot := rad.BallotRequest{
		Rule:     `borda`,
		Deadline: deadline.Format(time.UnixDate),
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	ExecuteTestNewBallotRequest(reqBallot, handler)

	time.Sleep(time.Second * 4)
	reqVote := rad.VoteRequest{
		AgentId: "1",
		VoteId:  "vote1",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}

	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusServiceUnavailable {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoVoteVotesTwice(t *testing.T) {
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
	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusForbidden {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoVoteUnregisteredVoter(t *testing.T) {
	_, handler := LaunchTestServer()
	deadline := time.Now().Add(time.Hour * 2)

	reqBallot := rad.BallotRequest{
		Rule:     `borda`,
		Deadline: deadline.Format(time.UnixDate),
		VotersId: []string{"1", "2", "3"},
		NbAlts:   2}

	ExecuteTestNewBallotRequest(reqBallot, handler)

	reqVote := rad.VoteRequest{
		AgentId: "4",
		VoteId:  "vote1",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}

	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusForbidden {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoVoteBallotDoesntExist(t *testing.T) {
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
		VoteId:  "vote2",
		Prefs:   []int{1, 2},
		Option:  []int{},
	}

	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusNotImplemented {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoVoteBallotToManyPrefs(t *testing.T) {
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
		Prefs:   []int{1, 2, 3},
		Option:  []int{},
	}

	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}

func TestDoVoteBallotToBadPrefs(t *testing.T) {
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
		Prefs:   []int{1, 1},
		Option:  []int{},
	}

	w := ExecuteTestVoteRequest(reqVote, handler)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("server status code %d", w.Result().StatusCode)
		t.Errorf("%s", w.Body)
	}
}
