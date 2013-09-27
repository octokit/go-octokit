package octokat

type AccessToken struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

func CreateAccessToken(params *Params) (*AccessToken, error) {
	if err := params.Require("client_id", "client_secret", "code"); err != nil {
		return nil, err
	}

	client := NewClient()
	client.BaseURL = GitHubURL
	var accessToken AccessToken
	headers := map[string]string{"Accept": "application/json"}
	options := Options{Headers: headers}
	if err := client.jsonPost("login/oauth/access_token", &options, params, &accessToken); err != nil {
		return nil, err
	}

	return &accessToken, nil
}
