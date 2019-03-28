package productplan

import (
	"fmt"
	"reflect"
	"testing"
)

func testCredentials(t *testing.T, credentials Credentials, headers map[string]string) {

	if want, got := headers, credentials.Headers(); !reflect.DeepEqual(want, got) {
		t.Errorf("Header %v, want %v", got, want)
	}
}

func TestOauthTokenCredentialsHttpHeader(t *testing.T) {
	oauthToken := "oauth-token"
	credentials := NewOauthTokenCredentials(oauthToken)
	expectedHeaderValue := fmt.Sprintf("Bearer %v", oauthToken)
	testCredentials(t, credentials, map[string]string{httpHeaderAuthorization: expectedHeaderValue})
}
