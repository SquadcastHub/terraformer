package squadcast

import (
	"io/ioutil"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadcastService struct {
	terraformutils.Service
}

func (s *SquadcastService) generateRequest(uri string) ([]byte, error) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.Args["accessToken"].(string))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
