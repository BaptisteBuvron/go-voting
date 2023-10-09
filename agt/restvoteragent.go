package agt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tp3/comsoc"
)

type RestClientAgent struct {
	id    string
	url   string
	perfs []comsoc.Alternative
}

func NewRestClientAgent(id string, url string, prefs []comsoc.Alternative) *RestClientAgent {
	return &RestClientAgent{id, url, prefs}
}

func parseInt(arg string, name string) int {
	i, err := strconv.Atoi(arg)
	if err == nil {
		fmt.Fprintf(os.Stderr, "%d (%s) is not an int value", arg, name)
		os.Exit(1)
	}
	return i
}

func parseStringList(arg string, name string) []string {
	return strings.Split(arg, ",")

}

func parseIntList(arg string, name string) []int {
	strings := parseStringList(arg, name)
	ints := make([]int, len(strings))
	for i, part := range strings {
		ints[i] = parseInt(part, fmt.Sprintf("%s[%d]", name, i))
	}
	return ints
}

func checkSize(argv []string, size int, syntax string) []int {
	fmt.Fprintf(os.Stderr, "%d (%s) is not an int value")
	os.Exit(1)
}

func request[R any](url string, req any) (R, int) {
	data, _ := json.Marshal(req)
	res, _ := http.Post(url, "application/json", bytes.NewBuffer(data))
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	var resp R
	json.Unmarshal(buf.Bytes(), &resp)
	return resp, res.StatusCode
}

func (rca *RestClientAgent) postNewBallot(argv []string) {
	if len(argv) != 4 {
		return
	}
	req := RequestNewBallot{
		Rule:         argv[0],
		Deadline:     argv[1],
		Voters:       parseStringList(argv[2], "voters"),
		Alternatives: parseInt(argv[3], "alternatives"),
	}
	res, code := request[ResponseNewBallot](rca.url+"/new_ballot", req)
	fmt.Printf("%s %d", res, code)
}

func (rca *RestClientAgent) postVote(agentId string, voteId string, prefs []int, options []int) {
	res, code := request[ResponseMessage](rca.url+"/vote", ResponseNewBallot{})
	fmt.Printf("%s %d", res, code)
}

func (rca *RestClientAgent) postResult(ballotId string) {
	res, code := request[ResponseResult](rca.url+"/result", ResponseResult{})
	fmt.Printf("%s %d", res, code)
}

func (rca *RestClientAgent) Start() {
	argc := len(os.Args[1:])
	if argc < 1 {
		log.Printf("SYNTAX: client new-ballot, vote, result", rca.id)
		os.Exit(1)
		return
	}

	log.Printf("démarrage de %s", rca.id)
	res, err := rca.doRequest()

	if err != nil {
		log.Fatal(rca.id, "error:", err.Error())
	} else {
		log.Printf("[%s] %d %s %d = %d\n", rca.id, rca.arg1, rca.operator, rca.arg2, res)
	}
}
