package http_client

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIService struct {
	client HTTPClient
}

func (s *APIService) FetchStatus(url string) (int, error) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
