package helpers

import (
	"context"
	"github.com/google/go-github/v31/github"
	"github.com/wplib/deploywp/only"
)


// Usage: {{ $user := GitHubGetRepository "gearboxworks" "docker-template" }}
type TypeGitHubGetRepository struct {
	Valid bool
	Error error
	Data *github.Repository
}

func (me TypeGitHubLogin) GetRepository(owner interface{}, repo interface{}) TypeGitHubGetRepository {
	var ret TypeGitHubGetRepository

	for range only.Once {
		op := ReflectString(owner)
		if op == nil {
			break
		}

		rp := ReflectString(repo)
		if rp == nil {
			break
		}

		ctx := context.Background()
		ret.Data, _, ret.Error = me.Client.Repositories.Get(ctx, *op, *rp)

		if ret.Error != nil {
			break
		}

		ret.Valid = true
	}

	return ret
}

func (me TypeGitHubGetRepository) GetName() TypeGenericStringArray {
	var ret TypeGenericStringArray

	for range only.Once {
		ret.Data = append(ret.Data, *me.Data.Name)
		ret.Valid = true
	}

	return ret
}

func (me TypeGitHubGetRepository) GetFullName() TypeGenericStringArray {
	var ret TypeGenericStringArray

	for range only.Once {
		ret.Data = append(ret.Data, *me.Data.FullName)
		ret.Valid = true
	}

	return ret
}

func (me TypeGitHubGetRepository) GetUrl() TypeGenericStringArray {
	var ret TypeGenericStringArray

	for range only.Once {
		ret.Data = append(ret.Data, *me.Data.URL)
		ret.Valid = true
	}

	return ret
}

// Usage: {{ $branch := GetHeadBranch }}
func (me TypeGitHubGetRepository) GetHeadBranch() TypeGenericString {
	var ret TypeGenericString

	ret.Data = me.Data.GetDefaultBranch()

	//for range only.Once {
	//	ret.Data = me.Data.GetDefaultBranch()
	//
	//	branchRefs, ret.Error = me.Data.Branches()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	headRef, ret.Error = repository.Head()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	var currentBranchName string
	//	ret.Error = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
	//		if branchRef.Hash() == headRef.Hash() {
	//			ret.Data = branchRef.Name().String()
	//
	//			return nil
	//		}
	//
	//		return nil
	//	})
	//
	//	if ret.Error != nil {
	//		break
	//	}
	//}

	return ret
}

func (me TypeGitHubGetRepository) GetCurrentCommitFromRepository() TypeGenericString {
	var ret TypeGenericString

	//for range only.Once {
	//	headRef, ret.Error = repository.Head()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	ret.Data = headRef.Hash().String()
	//}

	return ret
}

func (me TypeGitHubGetRepository) GetLatestTagFromRepository() TypeGenericString {
	var ret TypeGenericString

	//for range only.Once {
	//	tagRefs, ret.Error = repository.Tags()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	var latestTagCommit *object.Commit
	//	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
	//		revision := plumbing.Revision(tagRef.Name().String())
	//		tagCommitHash, ret.Error = repository.ResolveRevision(revision)
	//		if ret.Error != nil {
	//			return err
	//		}
	//
	//		commit, ret.Error = repository.CommitObject(*tagCommitHash)
	//		if ret.Error != nil {
	//			return err
	//		}
	//
	//		if latestTagCommit == nil {
	//			latestTagCommit = commit
	//			ret.Data = tagRef.Name().String()
	//		}
	//
	//		if commit.Committer.When.After(latestTagCommit.Committer.When) {
	//			latestTagCommit = commit
	//			ret.Data = tagRef.Name().String()
	//		}
	//
	//		return nil
	//	})
	//
	//	if ret.Error != nil {
	//		break
	//	}
	//}

	return ret
}


//////////////////////////////////////////////////////////////////////

// Usage: {{ $user := GitHubGetRepositories "gearboxworks" }}
type TypeGitHubGetRepositories struct {
	Valid bool
	Error error
	Data []*github.Repository
}
func (me TypeGitHubLogin) GetRepositories(owner interface{}) TypeGitHubGetRepositories {
	var ret TypeGitHubGetRepositories

	for range only.Once {
		op := ReflectString(owner)
		if op == nil {
			break
		}

		ctx := context.Background()
		ret.Data, _, ret.Error = me.Client.Repositories.List(ctx, *op, nil)

		if ret.Error != nil {
			break
		}

		//fmt.Printf("%+v\n", pack)
		ret.Valid = true
	}

	return ret
}
func (me TypeGitHubGetRepositories) GetNames() TypeGenericStringArray {
	var ret TypeGenericStringArray

	for range only.Once {
		for _, v := range me.Data {
			ret.Data = append(ret.Data, *v.Name)
		}
		ret.Valid = true
	}

	return ret
}
func (me TypeGitHubGetRepositories) GetFullNames() TypeGenericStringArray {
	var ret TypeGenericStringArray

	for range only.Once {
		for _, v := range me.Data {
			ret.Data = append(ret.Data, *v.FullName)
		}
		ret.Valid = true
	}

	return ret
}
func (me TypeGitHubGetRepositories) GetUrls() TypeGenericStringArray {
	var ret TypeGenericStringArray

	for range only.Once {
		for _, v := range me.Data {
			ret.Data = append(ret.Data, *v.URL)
		}
		ret.Valid = true
	}

	return ret
}
