package git

type Commit struct {
	Hash string
}

func NewCommit(hash string) *Commit {
	return &Commit{
		Hash: hash,
	}
}
