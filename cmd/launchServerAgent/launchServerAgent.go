package main

import (
	"fmt"

	sa "github.com/OrianeLanfranchi/ia04-projet1/agt/serveragent"
)

func main() {
	server := sa.NewServerAgent(":8080")
	server.Start()
	fmt.Scanln()
}
