package instance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

// Config is the structure of the configuration for a single GitHub instance.
type Config struct {
	Organization string      `json:"organization"`
	OAuth        OAuthConfig `json:"oauth"`
}

type OAuthConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	State        string `json:"state"`
}

type Instance interface {
	GetName() string
	GetOrganization() string
	TokenToCookie(token *oauth2.Token) (*http.Cookie, error)
	TokenFromCookie(r *http.Request) (*oauth2.Token, error)
	OAuthLoginURL() string
	OAuthCallback(ctx context.Context, state, code string) (*oauth2.Token, *github.User, error)
	OAuthIsAuthenticated(ctx context.Context, token *oauth2.Token) (*github.User, error)
}

type instance struct {
	name        string
	config      Config
	oauthConfig *oauth2.Config
}

func (i *instance) GetName() string {
	return i.name
}

func (i *instance) GetOrganization() string {
	return i.config.Organization
}

// TokenToCookie returns a cookie for the given oauth token.
func (i *instance) TokenToCookie(token *oauth2.Token) (*http.Cookie, error) {
	cookieValue, err := tokenToBase64(token)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     "kobs-oauth-github-" + i.config.Organization,
		Value:    cookieValue,
		Secure:   false,
		HttpOnly: false,
		Path:     "/",
	}, nil
}

// TokenFromCookie returns the token from the "kobs-oauth-github" cookie in the given request.
func (i *instance) TokenFromCookie(r *http.Request) (*oauth2.Token, error) {
	cookie, err := r.Cookie("kobs-oauth-github-" + i.config.Organization)
	if err != nil {
		return nil, err
	}

	return tokenFromBase64(cookie.Value)
}

func (i *instance) OAuthLoginURL() string {
	return i.oauthConfig.AuthCodeURL(i.config.OAuth.State, oauth2.AccessTypeOnline)
}

func (i *instance) OAuthCallback(ctx context.Context, state, code string) (*oauth2.Token, *github.User, error) {
	if state != i.config.OAuth.State {
		return nil, nil, fmt.Errorf("invalid oauth state, expected '%s', got '%s'", i.config.OAuth.State, state)
	}

	token, err := i.oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, nil, err
	}

	oauthClient := i.oauthConfig.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, nil, err
	}

	return token, user, nil
}

func (i *instance) OAuthIsAuthenticated(ctx context.Context, token *oauth2.Token) (*github.User, error) {
	oauthClient := i.oauthConfig.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}

// New returns a new GitHub instance for the given configuration.
func New(name string, options map[string]any) (Instance, error) {
	var config Config
	err := mapstructure.Decode(options, &config)
	if err != nil {
		return nil, err
	}

	return &instance{
		name:   name,
		config: config,
		oauthConfig: &oauth2.Config{
			ClientID:     config.OAuth.ClientID,
			ClientSecret: config.OAuth.ClientSecret,
			Scopes:       []string{"user", "repo", "notifications", "project"},
			Endpoint:     githuboauth.Endpoint,
		},
	}, nil
}