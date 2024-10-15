package serveragent

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

// Responses
type BallotResponse struct {
	ID string `json:"ballot-id"`
}

type ResultResponse struct {
	Winner  int   `json:"winner"`
	Ranking []int `json:"ranking"`
}

type Request interface {
	BallotRequest | VoteRequest
}
