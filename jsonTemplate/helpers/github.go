package helpers

import (
	"context"
	"github.com/google/go-github/v31/github"
	"github.com/wplib/deploywp/only"
	"reflect"
	"strings"
)


// Usage: {{ array := GitHubGetOrganization "gearboxworks" }}
func GitHubGetOrganization(i interface{}) []string {
	var sa []string

	for range only.Once {
		var err error

		v := reflect.ValueOf(i)
		if v.Kind() != reflect.String {
			break
		}

		var orgs []*github.Organization
		//orgs, err = fetchOrganizations(v.String())
		orgs, err = fetchOrganizations("")
		if err != nil {
			break
		}

		for _, org := range orgs {
			sa = append(sa, *org.Name)
			//fmt.Printf("%v. %v\n", i+1, org.GetLogin())
		}
	}

	return sa
}

func fetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}


// Usage: {{ $user := GitHubLogin "username" "password" "" }}
type TypeGitHubLogin struct {
	Valid bool
	Error error
	User *github.User
	Client *github.Client
}

func GitHubLogin(username interface{}, password interface{}, twofactor interface{}) TypeGitHubLogin {
	var auth TypeGitHubLogin

	for range only.Once {
		usernameString := ""
		if u := ReflectString(username); u != nil {
			usernameString = *u
		} else {
			usernameString = ""
		}
		if usernameString == "" {
			usernameString = UserPrompt("GitHub username: ")
		}


		passwordString := ""
		if p := ReflectString(password); p != nil {
			passwordString = *p
		} else {
			passwordString = ""
		}
		if passwordString == "" {
			passwordString = UserPromptHidden("GitHub password: ")
		}


		twofactorString := ""
		if f := ReflectString(twofactor); f != nil {
			twofactorString = *f
		} else {
			twofactorString = ""
		}


		tp := github.BasicAuthTransport{
			Username: strings.TrimSpace(usernameString),
			Password: strings.TrimSpace(passwordString),
		}

		//fmt.Printf("username: %s\tpassword: %s\t 2fa: %s\n", u.String(), p.String(), f.String())

		auth.Client = github.NewClient(tp.Client())
		ctx := context.Background()

		auth.User, _, auth.Error = auth.Client.Users.Get(ctx, "")
		if _, ok := auth.Error.(*github.TwoFactorAuthError); ok {
			// Is this a two-factor auth error? If so, prompt for OTP and try again.
			if twofactorString == "" {
				twofactorString = UserPrompt("GitHub 2FA password: ")
			}

			tp.OTP = strings.TrimSpace(twofactorString)
			auth.User, _, auth.Error = auth.Client.Users.Get(ctx, "")
		}

		if auth.Error != nil {
			break
		}

		auth.Valid = true
		//fmt.Printf("\n%v\n", github.Stringify(auth))
	}

	return auth
}


//// Usage: {{ $user := GitHubLogin "gearboxworks" "docker-template" "master" }}
//type TypeGetBranch struct {
//	Valid bool
//	Error error
//	Reference *github.Reference
//}
//func (me TypeGitHubLogin) GetBranch(owner interface{}, repo interface{}, reference interface{}) TypeGetBranch {
//	var ret TypeGetBranch
//
//	for range only.Once {
//		op := ReflectString(owner)
//		if op == nil {
//			break
//		}
//
//		rp := ReflectString(repo)
//		if rp == nil {
//			break
//		}
//
//		rfp := ReflectString(reference)
//		if rfp == nil {
//			break
//		}
//		if *rfp == "" {
//			*rfp = "master"
//		}
//
//		var ctx = context.Background()
//		ret.Reference, _, ret.Error = me.Client.Git.GetRef(ctx, *op, *rp, "refs/heads/" + *rfp)
//		if me.Error != nil {
//			break
//		}
//
//		if ret.Reference == nil {
//			break
//		}
//
//		ret.Valid = true
//		//fmt.Printf("\n>%s\n", ret.Reference.String())
//	}
//
//	return ret
//}
//
//
//
//// Usage: {{ $user := GitHubGetRepository "gearboxworks" "docker-template" }}
//type TypeGitHubGetRepository struct {
//	Valid bool
//	Error error
//	Data *github.Repository
//}
//func (me TypeGitHubLogin) GetRepository(owner interface{}, repo interface{}) TypeGitHubGetRepository {
//	var ret TypeGitHubGetRepository
//
//	for range only.Once {
//		op := ReflectString(owner)
//		if op == nil {
//			break
//		}
//
//		rp := ReflectString(repo)
//		if rp == nil {
//			break
//		}
//
//		ctx := context.Background()
//		ret.Data, _, ret.Error = me.Client.Repositories.Get(ctx, *op, *rp)
//
//		if ret.Error != nil {
//			break
//		}
//
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGitHubGetRepository) GetName() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		ret.Data = append(ret.Data, *me.Data.Name)
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGitHubGetRepository) GetFullName() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		ret.Data = append(ret.Data, *me.Data.FullName)
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGitHubGetRepository) GetUrl() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		ret.Data = append(ret.Data, *me.Data.URL)
//		ret.Valid = true
//	}
//
//	return ret
//}
//
//
//// Usage: {{ $user := GitHubGetRepositories "gearboxworks" }}
//type TypeGitHubGetRepositories struct {
//	Valid bool
//	Error error
//	Data []*github.Repository
//}
//func (me TypeGitHubLogin) GetRepositories(owner interface{}) TypeGitHubGetRepositories {
//	var ret TypeGitHubGetRepositories
//
//	for range only.Once {
//		op := ReflectString(owner)
//		if op == nil {
//			break
//		}
//
//		ctx := context.Background()
//		ret.Data, _, ret.Error = me.Client.Repositories.List(ctx, *op, nil)
//
//		if ret.Error != nil {
//			break
//		}
//
//		//fmt.Printf("%+v\n", pack)
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGitHubGetRepositories) GetNames() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		for _, v := range me.Data {
//			ret.Data = append(ret.Data, *v.Name)
//		}
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGitHubGetRepositories) GetFullNames() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		for _, v := range me.Data {
//			ret.Data = append(ret.Data, *v.FullName)
//		}
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGitHubGetRepositories) GetUrls() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		for _, v := range me.Data {
//			ret.Data = append(ret.Data, *v.URL)
//		}
//		ret.Valid = true
//	}
//
//	return ret
//}
//
//
//// Usage: {{ $user := GetCurrentBranchFromRepository "gearboxworks" "docker-template" }}
//func (me TypeGitHubGetRepository) GetCurrentBranchFromRepository() TypeGenericString {
//	var ret TypeGenericString
//
//	for range only.Once {
//		repo := ret.Data
//
//		branchRefs, ret.Error = repo.Branches()
//		if ret.Error != nil {
//			break
//		}
//
//		headRef, ret.Error = repository.Head()
//		if ret.Error != nil {
//			return "", err
//		}
//
//		var currentBranchName string
//		err = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
//			if branchRef.Hash() == headRef.Hash() {
//				ret.Data = branchRef.Name().String()
//
//				return nil
//			}
//
//			return nil
//		})
//		if ret.Error != nil {
//			return "", err
//		}
//	}
//
//	return ret
//}
//
//func (me TypeGitHubGetRepository) GetCurrentCommitFromRepository() TypeGenericString {
//	var ret TypeGenericString
//
//	for range only.Once {
//		headRef, ret.Error = repository.Head()
//		if ret.Error != nil {
//			break
//		}
//
//		ret.Data = headRef.Hash().String()
//	}
//
//	return ret
//}
//
//func (me TypeGitHubGetRepository) GetLatestTagFromRepository() TypeGenericString {
//	var ret TypeGenericString
//
//	for range only.Once {
//		tagRefs, ret.Error = repository.Tags()
//		if ret.Error != nil {
//			break
//		}
//
//		var latestTagCommit *object.Commit
//		err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
//			revision := plumbing.Revision(tagRef.Name().String())
//			tagCommitHash, ret.Error = repository.ResolveRevision(revision)
//			if ret.Error != nil {
//				return err
//			}
//
//			commit, ret.Error = repository.CommitObject(*tagCommitHash)
//			if ret.Error != nil {
//				return err
//			}
//
//			if latestTagCommit == nil {
//				latestTagCommit = commit
//				ret.Data = tagRef.Name().String()
//			}
//
//			if commit.Committer.When.After(latestTagCommit.Committer.When) {
//				latestTagCommit = commit
//				ret.Data = tagRef.Name().String()
//			}
//
//			return nil
//		})
//
//		if ret.Error != nil {
//			break
//		}
//	}
//
//	return ret
//}
