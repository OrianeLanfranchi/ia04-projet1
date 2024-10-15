package serveragent

import (
	"time"

	cs "github.com/OrianeLanfranchi/ia04-projet1/comsoc"
)

// Requests
type BallotRequest struct {
	Rule     string `json:"rule"`
	Deadline string `json:"deadline"`
	NbAlts   int    `json:"#alts"`
}

type VoteRequest struct {
	AgentId string `json:"agent-id"`
	VoteId  string `json:"vote-id"`
	Prefs   []int  `json:"prefs"`
	Option  []int  `json:"options"`
}

type ResultRequest struct {
	BallotId string `json:"ballot-id"`
}

type Request interface {
	BallotRequest | VoteRequest
}

// Responses
type BallotResponse struct {
	ID string `json:"ballot-id"`
}

type ResultResponse struct {
	Winner  int   `json:"winner"`
	Ranking []int `json:"ranking"`
}

// Useful types
type Ballot struct {
	Profile  cs.Profile
	Options  [][]int
	VotersId []string
	NbAlts   int
	Deadline time.Time
	Result   ResultResponse
}
