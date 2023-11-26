package main

import (
	"fmt"
	"ia04/agt"
	"os"
)

func main() {
	// First parser layer
	argc := len(os.Args)
	if argc < 3 {
		fmt.Fprintln(os.Stderr, "SYNTAX: client VOTER_ID {new_ballot,vote,result} [OPTIONS...]")
		os.Exit(1)
		return
	}
	voterID := os.Args[1]
	client := agt.NewRestClientAgent("http://localhost:8080", voterID)
	argv := os.Args[3:]
	// Call subcommand
	switch os.Args[2] {
	case "new_ballot":
		if len(argv) != 5 {
			fmt.Fprintln(os.Stderr, "SYNTAX: client VOTER_ID new_ballot {majority,borda,stv,approval,copeland} DEADLINE VOTERS ALTS TIE_BREAK")
			os.Exit(1)
		}
		client.CreateBallot(
			argv[0],
			agt.ParseTime(argv[1], "DEADLINE"),
			agt.ParseStringList(argv[2], "VOTERS"),
			agt.ParseInt(argv[3], "ALTS"),
			agt.ParseAlternatives(argv[4], "TIE_BREAK"),
		)
	case "vote":
		var options []int
		if len(argv) == 3 {
			options = agt.ParseIntList(argv[2], "OPTIONS")
		} else if len(argv) != 2 {
			fmt.Fprintln(os.Stderr, "SYNTAX: client VOTER_ID vote BALLOT_ID PREFS [OPTIONS]")
			os.Exit(1)
		}
		client.Vote(argv[0], agt.ParseAlternatives(argv[1], "PREFS"), options)
	case "result":
		if len(argv) != 1 {
			fmt.Fprintln(os.Stderr, "SYNTAX: client [VOTER_ID] result [BALLOT_ID]")
			os.Exit(1)
		}
		client.Result(argv[0])
	}
}
