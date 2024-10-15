package serveragent

import (
	"errors"
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
	BallotRequest | VoteRequest | ResultRequest
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
	SCF      SCFWrapper
	Options  [][]int
	VotersId []string
	NbAlts   int
	Deadline time.Time
	Result   ResultResponse
}

type SCFNoOption func(cs.Profile) (cs.Alternative, error)
type SCFOption func(cs.Profile, []int) (cs.Alternative, error)

type SCFFunction interface {
	Call(profile cs.Profile, options []int) (cs.Alternative, error)
}

type SCFWrapper struct {
	FuncNoOption SCFNoOption
	FuncOption   SCFOption
}

func (w SCFWrapper) Call(profile cs.Profile, options []int) (cs.Alternative, error) {
	if w.FuncNoOption != nil {
		return w.FuncNoOption(profile)
	} else if w.FuncOption != nil {
		return w.FuncOption(profile, options)
	}
	return cs.Alternative(-1), errors.New("no valid function provided")
}
