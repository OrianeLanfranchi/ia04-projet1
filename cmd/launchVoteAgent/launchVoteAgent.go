package main

import (
	"fmt"

	restclientagent "github.com/OrianeLanfranchi/ia04-projet1/agt/voteragent"
)

func main() {
	ag := restclientagent.NewRestClientAgent("ag_id1", "http://localhost:8080", "vote1", 5, make([]int, 0), make([]int, 0))
	ag.Start()
	fmt.Scanln()
}
