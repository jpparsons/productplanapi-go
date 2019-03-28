package productplan

const httpHeaderAuthorization = "Authorization"

// Credentials to be used for authenticating requests
type Credentials interface {
	Headers() map[string]string
}

// OAuth token authentication
type oauthTokenCredentials struct {
	oauthToken string
}

// NewOauthTokenCredentials construct Credentials using the OAuth access token.
func NewOauthTokenCredentials(oauthToken string) Credentials {
	return &oauthTokenCredentials{oauthToken: oauthToken}
}

func (c *oauthTokenCredentials) Headers() map[string]string {
	return map[string]string{httpHeaderAuthorization: "Bearer " + c.oauthToken}
}
