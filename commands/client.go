package commands

import (
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func newClient(cmd *cobra.Command) (*github.Client, error) {
	token := getFlagString(cmd, "token")
	if token != "" {
		return newOauthClient(token), nil
	}

	// 2. --user[:--password] command line arguments
	username := getFlagString(cmd, "username")
	if username != "" {
		return newBasicClient(
			username, 
			getFlagString(cmd, "password"),
			getFlagString(cmd, "otp"),
			), nil
	}

	// 3. GITHUB_OAUTH_TOKEN env variable
	if token := os.Getenv("GITHUB_OAUTH_TOKEN"); token != "" {
		return newOauthClient(token), nil
	}

	// 4. GITHUB_USERNAME, GITHUB_PASSWORD, GITHUB_OTP env variables
	if username := os.Getenv("GITHUB_USERNAME"); username != "" {
		return newBasicClient(
			username,
			os.Getenv("GITHUB_PASSWORD"),
			os.Getenv("GITHUB_OTP"),
		), nil
	}

	// 5. No authentication
	return github.NewClient(nil), nil
}

func newOauthClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

func newBasicClient(user, password, otp string) *github.Client {
	t := github.BasicAuthTransport{
		Username: user,
		Password: password,
		OTP: otp,
	}

	return github.NewClient(t.Client())
}
