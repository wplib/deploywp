package helpers

import (
	"context"
	"fmt"
	"github.com/google/go-github/v31/github"
	"github.com/src/golang.org/x/oauth2"
	"github.com/wplib/deploywp/only"
	"reflect"
	"strings"
)


func GitHubGetBranch(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			return true
		default:
			return false
	}
}


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
	Data *github.User
}
func GitHubLogin(username interface{}, password interface{}, twofactor interface{}) TypeGitHubLogin {
	var user TypeGitHubLogin

	for range only.Once {
		usernameString := ""
		if u := ReflectString(username); u != nil {
			usernameString = *u
		} else {
			usernameString = UserPrompt("GitHub username: ")
		}

		passwordString := ""
		if p := ReflectString(password); p != nil {
			passwordString = *p
		} else {
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

		client := github.NewClient(tp.Client())
		ctx := context.Background()

		user.Data, _, user.Error = client.Users.Get(ctx, "")
		if _, ok := user.Error.(*github.TwoFactorAuthError); ok {
			// Is this a two-factor auth error? If so, prompt for OTP and try again.
			if twofactorString == "" {
				twofactorString = UserPrompt("GitHub 2FA password: ")
			}

			tp.OTP = strings.TrimSpace(twofactorString)
			user.Data, _, user.Error = client.Users.Get(ctx, "")
		}

		if user.Error != nil {
			break
		}

		user.Valid = true
		fmt.Printf("\n%v\n", github.Stringify(user))
	}

	return user
}


// Usage: {{ $user := GitHubGetRepository TypeGitHubLogin.Data }}
type TypeGitHubGetRepository struct {
	Valid bool
	Error error
	Data *github.Repository
}
func GitHubGetRepository(auth interface{}, owner interface{}, repo interface{}) TypeGitHubGetRepository {
	var r TypeGitHubGetRepository

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

		tokenService := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: "<!-- Your API Keys -->"},
		)
		tokenClient := oauth2.NewClient(ctx, tokenService)

		client := github.NewClient(tokenClient)

		r.Data, _, r.Error = client.Repositories.Get(ctx, *op, *rp)

		if r.Error != nil {
			break
		}

		//fmt.Printf("%+v\n", pack)
		r.Valid = true
	}

	return r
}