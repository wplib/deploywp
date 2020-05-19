package helperGitHub

import (
	"context"
	"github.com/google/go-github/v31/github"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"reflect"
	"strings"
)


// Usage: {{ array := GitHubGetOrganization "gearboxworks" }}
func HelperGitHubGetOrganization(i interface{}) []string {
	var sa []string

	for range only.Once {
		var err error

		v := reflect.ValueOf(i)
		if v.Kind() != reflect.String {
			break
		}

		var orgs []*github.Organization
		//orgs, err = fetchOrganizations(v.Output())
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
type TypeLogin struct {
	User *github.User
	Client *github.Client

	Valid bool
	State *ux.State
}

func HelperGitHubLogin(username interface{}, password interface{}, twofactor interface{}) *TypeLogin {
	var auth TypeLogin

	for range only.Once {
		usernameString := ""
		if u := helperTypes.ReflectString(username); u != nil {
			usernameString = *u
		} else {
			usernameString = ""
		}
		if usernameString == "" {
			usernameString = helperSystem.HelperUserPrompt("GitHub username: ")
		}


		passwordString := ""
		if p := helperTypes.ReflectString(password); p != nil {
			passwordString = *p
		} else {
			passwordString = ""
		}
		if passwordString == "" {
			passwordString = helperSystem.HelperUserPromptHidden("GitHub password: ")
		}


		twofactorString := ""
		if f := helperTypes.ReflectString(twofactor); f != nil {
			twofactorString = *f
		} else {
			twofactorString = ""
		}


		tp := github.BasicAuthTransport{
			Username: strings.TrimSpace(usernameString),
			Password: strings.TrimSpace(passwordString),
		}

		//fmt.Printf("username: %s\tpassword: %s\t 2fa: %s\n", u.Output(), p.Output(), f.Output())

		auth.Client = github.NewClient(tp.Client())
		ctx := context.Background()

		var err error
		auth.User, _, err = auth.Client.Users.Get(ctx, "")
		if _, ok := err.(*github.TwoFactorAuthError); ok {
			// Is this a two-factor auth error? If so, prompt for OTP and try again.
			if twofactorString == "" {
				twofactorString = helperSystem.HelperUserPrompt("GitHub 2FA password: ")
			}

			tp.OTP = strings.TrimSpace(twofactorString)
			auth.User, _, err = auth.Client.Users.Get(ctx, "")
		}

		auth.State.SetError(err)
		if auth.State.IsError() {
			break
		}

		auth.Valid = true

		//fmt.Printf("\n%v\n", github.Stringify(auth))
	}

	return &auth
}


//// Usage: {{ $user := GitHubLogin "gearboxworks" "docker-template" "master" }}
//type TypeGetBranch struct {
//	Valid bool
//	Error error
//	Reference *github.Reference
//}
//func (me TypeLogin) GetBranch(owner interface{}, repo interface{}, reference interface{}) TypeGetBranch {
//	var ret TypeGetBranch
//
//	for range only.Once {
//		op := general.ReflectString(owner)
//		if op == nil {
//			break
//		}
//
//		rp := general.ReflectString(repo)
//		if rp == nil {
//			break
//		}
//
//		rfp := general.ReflectString(reference)
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
//		//fmt.Printf("\n>%s\n", ret.Reference.Output())
//	}
//
//	return ret
//}
//
//
//
//// Usage: {{ $user := GitHubGetRepository "gearboxworks" "docker-template" }}
//type TypeGetRepository struct {
//	Valid bool
//	Error error
//	Data *github.Repository
//}
//func (me TypeLogin) GetRepository(owner interface{}, repo interface{}) TypeGetRepository {
//	var ret TypeGetRepository
//
//	for range only.Once {
//		op := general.ReflectString(owner)
//		if op == nil {
//			break
//		}
//
//		rp := general.ReflectString(repo)
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
//func (me TypeGetRepository) GetName() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		ret.Data = append(ret.Data, *me.Data.Name)
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGetRepository) GetFullName() TypeGenericStringArray {
//	var ret TypeGenericStringArray
//
//	for range only.Once {
//		ret.Data = append(ret.Data, *me.Data.FullName)
//		ret.Valid = true
//	}
//
//	return ret
//}
//func (me TypeGetRepository) GetUrl() TypeGenericStringArray {
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
//type TypeGetRepositories struct {
//	Valid bool
//	Error error
//	Data []*github.Repository
//}
//func (me TypeLogin) GetRepositories(owner interface{}) TypeGetRepositories {
//	var ret TypeGetRepositories
//
//	for range only.Once {
//		op := general.ReflectString(owner)
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
//func (me TypeGetRepositories) GetNames() TypeGenericStringArray {
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
//func (me TypeGetRepositories) GetFullNames() TypeGenericStringArray {
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
//func (me TypeGetRepositories) GetUrls() TypeGenericStringArray {
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
//func (me TypeGetRepository) GetCurrentBranchFromRepository() TypeGenericString {
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
//				ret.Data = branchRef.Name().Output()
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
//func (me TypeGetRepository) GetCurrentCommitFromRepository() TypeGenericString {
//	var ret TypeGenericString
//
//	for range only.Once {
//		headRef, ret.Error = repository.Head()
//		if ret.Error != nil {
//			break
//		}
//
//		ret.Data = headRef.Hash().Output()
//	}
//
//	return ret
//}
//
//func (me TypeGetRepository) GetLatestTagFromRepository() TypeGenericString {
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
//			revision := plumbing.Revision(tagRef.Name().Output())
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
//				ret.Data = tagRef.Name().Output()
//			}
//
//			if commit.Committer.When.After(latestTagCommit.Committer.When) {
//				latestTagCommit = commit
//				ret.Data = tagRef.Name().Output()
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
