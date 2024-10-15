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
	Rule     string
	Profile  cs.Profile
	SCF      SCFWrapper
	SWF      SWFWrapper
	Options  [][]int
	VotersId []string
	NbAlts   int
	Deadline time.Time
	Result   ResultResponse
}

// SCF Type
type SCFNoOption func(cs.Profile) (cs.Alternative, error)
type SCFOption func(cs.Profile, []int) (cs.Alternative, error)

type SCFFunction interface {
	Call(profile cs.Profile, options []int) (cs.Alternative, error)
}

type SCFWrapper struct {
	FuncNoOption  SCFNoOption
	FuncOneOption SCFOption
}

// TODO - améliorer la méthode Call pour qu'elle fasse toutes les vérifications
func (w SCFWrapper) Call(profile cs.Profile, options [][]int) (cs.Alternative, error) {
	if w.FuncNoOption != nil {
		return w.FuncNoOption(profile)
	} else if w.FuncOneOption != nil {
		listOptions := make([]int, len(options))
		for i, option := range options {
			listOptions[i] = option[0]
		}
		return w.FuncOneOption(profile, listOptions)
	}
	return cs.Alternative(-1), errors.New("(w SCFWrapper) - Pas de fonction valide")
}

// SWF
type SWFNoOption func(cs.Profile) ([]cs.Alternative, error)
type SWFOption func(cs.Profile, []int) ([]cs.Alternative, error)

type SCWFFunction interface {
	Call(profile cs.Profile, options []int) ([]cs.Alternative, error)
}

type SWFWrapper struct {
	FuncNoOption  SWFNoOption
	FuncOneOption SWFOption
}

// TODO - améliorer la méthode Call pour qu'elle fasse toutes les vérifications
func (w SWFWrapper) Call(profile cs.Profile, options [][]int) ([]cs.Alternative, error) {
	if w.FuncNoOption != nil {
		return w.FuncNoOption(profile)
	} else if w.FuncOneOption != nil {
		listOptions := make([]int, len(options))
		for i, option := range options {
			listOptions[i] = option[0]
		}
		return w.FuncOneOption(profile, listOptions)
	}
	return nil, errors.New("(w SWFWrapper) - Pas de fonction valide")
}
