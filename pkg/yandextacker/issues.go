package yandextacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"codeberg.org/kalsin/UtelBot/pkg/configs"
	log "codeberg.org/kalsin/UtelBot/pkg/logging"
)

// Структура для хранения отправляемых данных в yandex tracker api
type RequestIssue struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Queue       string `json:"queue"`
}

// Структура для хнанения принимаемых данных от yandex tracker api
type ResponseIssue struct {
	Assignee struct {
		Name string `json:"display"`
	} `json:"assignee"`

	Status struct {
		Display string `json:"display"`
	} `json:"status"`

	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func NewRequestIssue() *RequestIssue {
	return &RequestIssue{
		Summary:     "",
		Description: "",
		Queue:       configs.GetTrackerConfig().Queue,
	}
}

func newResponseIssue() *ResponseIssue {
	return &ResponseIssue{}
}

// Метод для создания новой задачи в yandex tracker
func (c *Connection) CreateNewIssue(issue *RequestIssue) (string, error) {
	url := fmt.Sprintf("%s%s", c.Host, "/issues")

	body, err := json.Marshal(issue)
	if err != nil {
		return "", err
	}
	responseBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, url, responseBody)
	if err != nil {
		return "", err
	}

	for key, val := range c.Headers {
		req.Header.Add(key, val)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		err = fmt.Errorf("send issue. Status: %s", resp.Status)
		return "", err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var bodyMap map[string]interface{}

	err = json.Unmarshal(rbody, &bodyMap)
	if err != nil {
		return "", err
	}

	issueID := fmt.Sprintf("%v", bodyMap["id"])

	return issueID, nil
}

// Метод для получения данных задачи по ее идентификатору из yandex tracker
func (c *Connection) GetIssueDataByID(issueID string) (*ResponseIssue, error) {
	url := fmt.Sprintf("%s%s%s", c.Host, "/issues/", issueID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, val := range c.Headers {
		req.Header.Add(key, val)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("get issue by ID. Status: %s", resp.Status)
		log.Warning.Println(err)
		return nil, err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}

	issue := newResponseIssue()

	err = json.Unmarshal(rbody, issue)
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}

	return issue, nil
}

func (c *Connection) AddComment(issueID, text string) error {
	url := fmt.Sprintf("%s/issues/%s/comments?isAddToFollowers=false", c.Host, issueID)

	body, err := json.Marshal(text)
	if err != nil {
		return err
	}
	responseBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, url, responseBody)
	if err != nil {
		return err
	}

	for key, val := range c.Headers {
		req.Header.Add(key, val)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		err = fmt.Errorf("send issue. Status: %s", resp.Status)
		return err
	}

	return nil
}
