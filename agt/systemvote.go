package agt

/*func RunSystemDeVoteFromProfile(profile comsoc.Profile, scf func(comsoc.Profile) (comsoc.Alternative, error)) {
	voters := make([]VoterAgentI, 0)
	ch := make(chan []comsoc.Alternative)
	for id, alts := range profile {
		voter := &VoterAgent{Agent: Agent{ID: AgentID(id + 1), Name: fmt.Sprintf("Voteur-%d", id), Prefs: alts}, Channel: ch}
		voters = append(voters, voter)
	}
	ballot := &BallotAgent{Agent{ID: AgentID(0), Name: "Ballot", Prefs: nil}, Channel: ch, VoterCount: len(profile), Scf: scf}
	RunSystemDeVote(voters, ballot)
}

func RunSystemDeVote(voters []VoterAgentI, ballot BallotAgentI) {
	ballot.Start()
	for _, voter := range voters {
		voter.Start()
	}
	time.Sleep(time.Second)
}*/
