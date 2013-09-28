package octokat

type AccessToken struct {
	Token     string `json:"access_token,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}

type AccessTokenParams struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

func CreateAccessToken(options *Options) (accessToken *AccessToken, err error) {
	client := NewClient()
	client.BaseURL = GitHubURL

	if options == nil {
		options = &Options{}
	}

	if options.Headers == nil {
		options.Headers = Headers{}
	}

	options.Headers["Accept"] = "application/json"
	err = client.jsonPost("login/oauth/access_token", options, &accessToken)
	return
}
