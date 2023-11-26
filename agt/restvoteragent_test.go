package agt

import (
	"ia04/comsoc"
	"testing"
	"time"
)

func TestRestServerAgent(t *testing.T) {
	assert := comsoc.NewAssert(t)
	server := NewRestServerAgent(":8080")
	server.Start()
	client := NewRestClientAgent("http://localhost:8080", "v1")
	ok := client.WaitAvailable(time.Second * 10)
	assert.True(ok)
	// go run ia04/cmd/client v1 new-ballot majority '2023-11-26T16:27:11+00:00' 'v1,v2,v3' 5 '1,2,3,4,5'
	ballotID, err := client.CreateBallot(
		"majority",
		time.Now().Add(time.Second*2),
		[]string{"v1", "v2", "v3"},
		5,
		[]comsoc.Alternative{1, 2, 3, 4, 5},
	)
	assert.NoError(err)
	// go run ia04/cmd/client v1 vote majority-18c0c24a3245e '4,2,3,1,5'
	err = client.Vote(ballotID, []comsoc.Alternative{4, 2, 3, 1, 5}, []int{})
	assert.NoError(err)

	// wait ballot
	time.Sleep(time.Second * 3)

	// go run ia04/cmd/client v1 result majority-18c0c24a3245e
	winner, err := client.Result(ballotID)
	assert.NoError(err)
	assert.Equal(winner, 4)
}
