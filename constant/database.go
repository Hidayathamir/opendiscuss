package constant

type VoteOption int

// should be sync with db table vote_option
const (
	VoteOptionThumbsUp   VoteOption = 1 // id table vote option row thumbs up
	VoteOptionThumbsDown VoteOption = 2 // id table vote option row thumbs down
	VoteOptionCancel     VoteOption = 3 // id table vote option row cancel
)
