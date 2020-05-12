package helperGit

import (
	"context"
	"github.com/google/go-github/v31/github"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)


type TypeGetRepository struct {
	Valid bool
	Error error
	Data *github.Repository
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepository "gearboxworks" "docker-template" }}
func (me *TypeLogin) GetRepository(owner interface{}, repo interface{}) *TypeGetRepository {
	var ret TypeGetRepository

	for range only.Once {
		op := helperTypes.ReflectString(owner)
		if op == nil {
			break
		}

		rp := helperTypes.ReflectString(repo)
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

	return &ret
}

func (me *TypeGetRepository) GetName() helperTypes.TypeGenericStringArray {
	var ret helperTypes.TypeGenericStringArray

	for range only.Once {
		ret.Array = append(ret.Array, *me.Data.Name)
		ret.Valid = true
	}

	return ret
}

func (me *TypeGetRepository) GetFullName() helperTypes.TypeGenericStringArray {
	var ret helperTypes.TypeGenericStringArray

	for range only.Once {
		ret.Array = append(ret.Array, *me.Data.FullName)
		ret.Valid = true
	}

	return ret
}

func (me *TypeGetRepository) GetUrl() helperTypes.TypeGenericStringArray {
	var ret helperTypes.TypeGenericStringArray

	for range only.Once {
		ret.Array = append(ret.Array, *me.Data.URL)
		ret.Valid = true
	}

	return ret
}

// Usage: {{ $branch := GetHeadBranch }}
func (me *TypeGetRepository) GetHeadBranch() helperTypes.TypeGenericString {
	var ret helperTypes.TypeGenericString

	ret.String = me.Data.GetDefaultBranch()

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

func (me *TypeGetRepository) GetCurrentCommitFromRepository() helperTypes.TypeGenericString {
	var ret helperTypes.TypeGenericString

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

func (me *TypeGetRepository) GetLatestTagFromRepository() helperTypes.TypeGenericString {
	var ret helperTypes.TypeGenericString

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

type TypeGetRepositories struct {
	Valid bool
	Error error
	Data []*github.Repository
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
func (me *TypeLogin) GetRepositories(owner interface{}) *TypeGetRepositories {
	var ret TypeGetRepositories

	for range only.Once {
		op := helperTypes.ReflectString(owner)
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

	return &ret
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
//		{{ $names := $repos.GetNames }}
func (me *TypeGetRepositories) GetNames() helperTypes.TypeGenericStringArray {
	var ret helperTypes.TypeGenericStringArray

	for range only.Once {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.Name)
		}
		ret.Valid = true
	}

	return ret
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
//		{{ $names := $repos.GetFullNames }}
func (me *TypeGetRepositories) GetFullNames() helperTypes.TypeGenericStringArray {
	var ret helperTypes.TypeGenericStringArray

	for range only.Once {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.FullName)
		}
		ret.Valid = true
	}

	return ret
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
//		{{ $urls := $repos.GetUrls }}
func (me *TypeGetRepositories) GetUrls() helperTypes.TypeGenericStringArray {
	var ret helperTypes.TypeGenericStringArray

	for range only.Once {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.URL)
		}
		ret.Valid = true
	}

	return ret
}
