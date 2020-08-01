package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {

}

type GetUserDataSetsResponse struct {
	Type         string        `json:"type"`
	Username     string        `json:"username"`
	DataSetSpecs []DataSetSpec `json:"dataSetSpecs"`
	Error        string        `json:"error"`
}

func GetUserDataSets(username string) ([]DataSetSpec, error) {

	BaseURL := "http://localhost:8080"
	req := struct {
		Username    string `json:"username"`
		Token       string `json:"token"`
		RequestType string `json:"type"`
	}{Username: username, RequestType: "getUserDataSets", Token: GetUserInfo().Token}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("imposssible marshalling error: %s", err.Error())
	}

	r := bytes.NewReader(reqBytes)
	resp, err := http.Post(BaseURL, "application/json", r)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request execution failed: status code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %d", resp.StatusCode)
	}

	response := GetUserDataSetsResponse{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if response.Error != "" {
		return nil, fmt.Errorf("authentication failed to get data sets: %s", response.Error)
	}
	return response.DataSetSpecs, nil
}
