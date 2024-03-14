package mocking

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Information struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetInfo(userId int) (*Information, error) {
	response, err := http.Get(fmt.Sprintf("https://yourapitomock.io?user=%d", userId))
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code received")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	info := &Information{}
	err = json.Unmarshal(body, info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func AddUser(information *Information) (int, error) {
	bodyJson, err := json.Marshal(information)
	if err != nil {
		return 0, err
	}

	response, err := http.Post("https://yourapitomock.io/user", "application/json", bytes.NewReader(bodyJson))
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	responseJson := map[string]int{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return 0, err
	}

	return responseJson["id"], nil
}

func UpdateUser(id int, information *Information) error {
	bodyJson, err := json.Marshal(information)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://yourapitomock.io/user/%d", id), bytes.NewReader(bodyJson))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		return errors.New("invalid status code received")
	}

	return nil
}
