package protonvpn

import (
	"io/ioutil"
	"net/http"
)

const logicalsFileUrl string = "https://api.protonvpn.ch/vpn/logicals"

func GetLogicals(fetchOffline bool) (*LogicalsResponse, error) {
	var data []byte

	if !fetchOffline {
		fetchedOnline, err := fetchLogicals()

		if err != nil {
			return nil, err
		}

		data = fetchedOnline
	} else {
		d, err := embedFS.ReadFile("data/logical.json")

		if err != nil {
			return nil, err
		}

		data = d
	}

	return parseLogicalResponse(data)
}

/**
 * Fetches the Logicals
 */
func fetchLogicals() ([]byte, error) {
	resp, err := http.Get(logicalsFileUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
