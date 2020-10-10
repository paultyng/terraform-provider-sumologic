package sumologic

import (
	"encoding/json"
	"fmt"
)

type UniversalSource struct {
	ID        int             `json:"id,omitempty"`
	Type      string          `json:"sourceType"`
	Config    json.RawMessage `json:"config"`
	SchemaRef SchemaReference `json:"schemaRef"`
}

type SchemaReference struct {
	Type    string `json:"type"`
	Version string `json:"version,omitempty"`
}

func (s *Client) CreateUniversalSource(source UniversalSource, collectorID int) (int, error) {

	type UniversalSourceMessage struct {
		Source UniversalSource `json:"source"`
	}

	request := UniversalSourceMessage{
		Source: source,
	}

	urlPath := fmt.Sprintf("v1/collectors/%d/sources", collectorID)

	body, err := s.Post(urlPath, request)

	if err != nil {
		return -1, err
	}

	var response UniversalSourceMessage
	err = json.Unmarshal(body, &response)

	if err != nil {
		return -1, err
	}

	return response.Source.ID, nil
}

func (s *Client) GetUniversalSource(collectorID, sourceID int) (*UniversalSource, error) {
	urlPath := fmt.Sprintf("v1/collectors/%d/sources/%d", collectorID, sourceID)
	body, _, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}

	if body == nil {
		return nil, nil
	}

	type UniversalSourceResponse struct {
		Source UniversalSource `json:"source"`
	}

	var response UniversalSourceResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response.Source, nil
}

func (s *Client) UpdateUniversalSource(source UniversalSource, collectorID int) error {
	url := fmt.Sprintf("v1/collectors/%d/sources/%d", collectorID, source.ID)

	type UniversalSourceMessage struct {
		Source UniversalSource `json:"source"`
	}

	request := UniversalSourceMessage{
		Source: source,
	}

	_, err := s.Put(url, request)

	return err
}
