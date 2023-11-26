package agt

import (
	"github.com/BaptisteBuvron/go-voting/comsoc"
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

	// Case 1 : Normal vote
	ballotID, err := client.CreateBallot(
		"majority",
		time.Now().Add(time.Second*2),
		[]string{"v1", "v2", "v3"},
		5,
		[]comsoc.Alternative{1, 2, 3, 4, 5},
	)
	assert.NoError(err)
	err = client.Vote(ballotID, []comsoc.Alternative{4, 2, 3, 1, 5}, []int{})
	assert.NoError(err)
	time.Sleep(time.Second * 3)
	res, err := client.Result(ballotID)
	assert.NoError(err)
	assert.Equal(res.Winner, comsoc.Alternative(4))
	assert.DeepEqual(res.Ranking, []comsoc.Alternative{4})

	// Case 2: No voters
	ballotID, err = client.CreateBallot(
		"majority",
		time.Now().Add(time.Second*2),
		[]string{"v1", "v2", "v3"},
		5,
		[]comsoc.Alternative{2, 1, 3, 4, 5},
	)
	assert.NoError(err)
	time.Sleep(time.Second * 3)
	res, err = client.Result(ballotID)
	assert.NoError(err)
	assert.Equal(res.Winner, comsoc.Alternative(2))
	assert.DeepEqual(res.Ranking, []comsoc.Alternative{2, 1, 3, 4, 5})
}
