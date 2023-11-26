package agt

import (
	"crypto/rand"
	"fmt"
	"ia04/comsoc"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// A ballot is a device used to cast votes in an election and may be found as a piece of paper or a small ball used in voting.
// ref: https://en.wikipedia.org/wiki/Ballot
type BallotAgent struct {
	id               string
	rule             string
	deadline         time.Time
	voters           map[string]bool
	alternativeCount int
	tieBreak         comsoc.TieBreak
	thresholds       []int
	profiles         comsoc.Profile
}

// Create an new BallotAgent by verifying values
func NewBallotAgent(rule string, deadline time.Time, voters []string, alternativeCount int, tieBreak []comsoc.Alternative) (*BallotAgent, error) {
	// Check if the Rule is set
	if rule == "" {
		return nil, comsoc.HTTPError{Code: http.StatusBadRequest, Message: "Rule is not set"}
	}
	// Check if rule is implemented
	switch rule {
	case "majority", "borda", "stv", "approval", "copeland":
	default:
		return nil, comsoc.HTTPErrorf(http.StatusNotImplemented, "Rule %s is not implemented", rule)
	}
	// Check if the Deadline is set
	if deadline.Before(time.Now()) {
		return nil, comsoc.HTTPErrorf(http.StatusBadRequest, "Deadline is already passed")
	}
	// Check if at least one voter is present
	if len(voters) == 0 {
		return nil, comsoc.HTTPErrorf(http.StatusBadRequest, "You must have at lease one voter")
	}
	// Check if the number of alternatives is valid
	if alternativeCount < 2 {
		return nil, comsoc.HTTPErrorf(http.StatusBadRequest, "You must have at lease two alternatives")
	}
	// TieBreak must be valid
	err := comsoc.CheckAlternatives(tieBreak, alternativeCount)
	if err != nil {
		return nil, err
	}
	// generate unique timestamp-based id
	timestamp := time.Now().UnixMilli()
	uniqueNumber, err := rand.Int(rand.Reader, big.NewInt(0xffff))
	if err != nil {
		log.Println(err) // do not show uncontrolled informations
		return nil, comsoc.HTTPErrorf(http.StatusInternalServerError, "Can't generate id")
	}
	id := fmt.Sprintf("%s-%s", rule, strconv.FormatInt(timestamp<<8+uniqueNumber.Int64(), 16))
	// Map voters
	mapVoters := make(map[string]bool)
	for _, voter := range voters {
		mapVoters[voter] = false
	}
	// No error
	return &BallotAgent{id, rule, deadline, mapVoters, alternativeCount, comsoc.TieBreakFactory(tieBreak), nil, nil}, nil

}

// Add a vote to a BallotAgent, can raise if: already voted, not authorized to vote, deadline is over, unnecessary options and bad alternatives
// thread-unsafe
func (b *BallotAgent) Vote(voterId string, alts []comsoc.Alternative, options []int) error {
	// Verify the deadline
	if time.Now().After(b.deadline) {
		return comsoc.HTTPErrorf(http.StatusServiceUnavailable, "Deadline %s is over for %s", b.deadline, b.id)
	}
	// Verify if voter exists in voters and not already vote
	alreadyVote, allowedToVote := b.voters[voterId]
	if !allowedToVote {
		return comsoc.HTTPErrorf(http.StatusForbidden, "Voter %s are not allowed to vote for %s", voterId, b.id)
	}
	if alreadyVote {
		return comsoc.HTTPErrorf(http.StatusForbidden, "Voter %s has already voted for %s", voterId, b.id)
	}
	// Verify if the number of alternatives is correct
	err := comsoc.CheckAlternatives(alts, b.alternativeCount)
	if err != nil {
		return err
	}
	// Check if thresholds is only present for approval
	if b.rule == "approval" {
		if len(b.thresholds) == 1 {
			return comsoc.HTTPErrorf(http.StatusBadRequest, "You must provide only the threshold as option")
		}
		threshold := b.thresholds[0]
		if threshold < 0 && threshold <= b.alternativeCount {
			return comsoc.HTTPErrorf(http.StatusBadRequest, "Invalid threshold %d", threshold)
		} else {
			b.thresholds = append(b.thresholds, threshold)
		}
	} else if len(b.thresholds) != 0 {
		return comsoc.HTTPErrorf(http.StatusBadRequest, "No options available for %s", b.rule)
	}
	// Mark voter as voted and add is vote
	b.voters[voterId] = true
	b.profiles = append(b.profiles, alts)
	// Everything go brrr
	return nil
}

// Get the result of BallotAgent only after deadline
func (b *BallotAgent) result() (comsoc.Alternative, []comsoc.Alternative, error) {
	// Check if vote is over
	if b.deadline.After(time.Now()) {
		return comsoc.Alternative(-1), nil, comsoc.HTTPErrorf(http.StatusServiceUnavailable, "Deadline is not over come back at %v", b.deadline)
	}
	// Get scf
	var scf comsoc.SCF
	switch b.rule {
	case "majority":
		scf = comsoc.MajoritySCF
	case "borda":
		scf = comsoc.BordaSCF
	case "stv":
		scf = comsoc.STV_SCF
	case "approval":
		scf = func(p comsoc.Profile) (bestAlts []comsoc.Alternative, err error) {
			return comsoc.ApprovalSCF(p, b.thresholds)
		}
	case "copeland":
		scf = comsoc.CopelandSCF
	}
	results, err := scf(b.profiles)
	if err != nil {
		return comsoc.Alternative(-1), nil, err
	}
	// apply tie break
	winner, err := b.tieBreak(results)
	// Get result
	return winner, results, err

}
