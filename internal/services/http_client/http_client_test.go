package http_client

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestFetchStatus(t *testing.T) {
	mockClient := new(MockHTTPClient)

	req, _ := http.NewRequest("GET", "https://example.com", nil)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("OK")),
	}

	mockClient.On("Do", req).Return(resp, nil)

	svc := APIService{client: mockClient}
	status, err := svc.FetchStatus("https://example.com")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	mockClient.AssertExpectations(t)
}
