package protonvpn

import (
	"encoding/json"
)

func parseLogicalResponse(body []byte) (*LogicalsResponse, error) {
	var logicalsResponse LogicalsResponse

	if err := json.Unmarshal(body, &logicalsResponse); err != nil {
		return nil, err
	}

	return &logicalsResponse, nil
}
