// +build e2e

package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ekozlova94/parking/internal/parking"
	"github.com/ekozlova94/parking/pkg/forms"
	"github.com/stretchr/testify/require"
)

const host = "http://localhost:8000/v1"

func TestSubscription(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		t.Parallel()
		// act
		body, code, err := post(host+"/subscription", &forms.Request{AutoNumber: 2})
		// assert
		require.NoError(t, err)
		require.Equal(t, 200, code)
		resp, errDecode := decodeJson(body)
		require.NoError(t, errDecode)
		require.Equal(t, true, resp.Success)
	})
	t.Run("conflict", func(t *testing.T) {
		// arrange
		t.Parallel()
		_, _, _ = post(host+"/subscription", &forms.Request{AutoNumber: 1})
		// act
		body, code, err := post(host+"/subscription", &forms.Request{AutoNumber: 1})
		// assert
		require.NoError(t, err)
		require.Equal(t, 409, code)
		resp, errDecode := decodeString(body)
		require.NoError(t, errDecode)
		require.Equal(t, parking.ErrConflict.Error(), resp)
	})
}

func post(host string, req *forms.Request) (body io.ReadCloser, code int, err error) {
	r, err := json.Marshal(req)
	if err != nil {
		return
	}
	resp, err := http.Post(host, "application/json", strings.NewReader(string(r)))
	if err != nil {
		return
	}
	return resp.Body, resp.StatusCode, nil
}

func decodeJson(body io.ReadCloser) (*forms.Response, error) {
	decoder := json.NewDecoder(body)
	var result forms.Response
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeString(body io.ReadCloser) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	if err != nil {
		return "", err
	}
	s := buf.String()
	return strings.TrimSpace(s), nil
}
