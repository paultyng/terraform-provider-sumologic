package sumologic

import (
	"encoding/json"
	"fmt"
)

func (s *Client) CreateScanBudget(scanBudget ScanBudget) (string, error) {
	urlWithoutParams := "v1/budgets"

	data, err := s.Post(urlWithoutParams, scanBudget)
	if err != nil {
		return "", err
	}

	var createdScanBudget ScanBudget

	err = json.Unmarshal(data, &createdScanBudget)
	if err != nil {
		return "", err
	}

	return createdScanBudget.ID, nil
}

func (s *Client) GetScanBudget(id string) (*ScanBudget, error) {
	urlWithoutParams := "v1/budgets/%s"
	paramString := ""
	sprintfArgs := []interface{}{}
	sprintfArgs = append(sprintfArgs, id)

	urlWithParams := fmt.Sprintf(urlWithoutParams+paramString, sprintfArgs...)

	data, _, err := s.Get(urlWithParams)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	var scanBudget ScanBudget

	err = json.Unmarshal(data, &scanBudget)
	if err != nil {
		return nil, err
	}

	return &scanBudget, nil

}

func (s *Client) DeleteScanBudget(id string) error {
	urlWithoutParams := "v1/budgets/%s"
	paramString := ""
	sprintfArgs := []interface{}{}
	sprintfArgs = append(sprintfArgs, id)

	urlWithParams := fmt.Sprintf(urlWithoutParams+paramString, sprintfArgs...)

	_, err := s.Delete(urlWithParams)

	return err
}

func (s *Client) UpdateScanBudget(scanBudget ScanBudget) error {
	urlWithoutParams := "v1/budgets/%s"
	paramString := ""
	sprintfArgs := []interface{}{}
	sprintfArgs = append(sprintfArgs, scanBudget.ID)

	urlWithParams := fmt.Sprintf(urlWithoutParams+paramString, sprintfArgs...)

	scanBudget.ID = ""

	_, err := s.Put(urlWithParams, scanBudget)

	return err

}

type ScanBudget struct {
	ID             string                 `json:"id,omitempty"`
	OrgID          string                 `json:"orgId"`
	Name           string                 `json:"name"`
	Capacity       int                    `json:"capacity"`
	Unit           string                 `json:"unit"`
	BudgetType     string                 `json:"budgetType,omitempty"`
	Window         string                 `json:"window"`
	Grouping       string                 `json:"applicableOn"`
	GroupingEntity string                 `json:"groupBy"`
	Action         string                 `json:"action"`
	Scope          map[string]interface{} `json:"scope"`
	Status         string                 `json:"status"`
}
