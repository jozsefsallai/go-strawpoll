// Package strawpoll provides an easy way to fetch and create polls from/to
// strawpoll.me.
//
// To fetch a poll:
//   func main() {
// 		poll, err := strawpoll.Get(1)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(poll.Title)
// 	}
// (where `1` is the ID of the poll on strawpoll.me)
//
// To create a poll:
// 	func main() {
// 		poll, err := strawpoll.Create(
// 			"title of the poll",
// 			[]string{"option 1", "option 2", "option 3"}, // at least 2
// 			false, // multi-choice
// 			strawpoll.DupcheckNormal, // duplication checking level
// 			false, // require CAPTCHA
// 		)
//
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		fmt.Println(poll.ID)
// 	}
package strawpoll

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL string = "https://www.strawpoll.me/api/v2"

// DupcheckNormal = IP duplication checking
// DupcheckPermissive = Browser cookie duplication checking
// DupcheckDisabled = No duplication checking
const (
	DupcheckNormal     = "normal"
	DupcheckPermissive = "permissive"
	DupcheckDisabled   = "disabled"
)

func isValidDupcheck(input string) bool {
	switch input {
	case DupcheckNormal:
	case DupcheckPermissive:
	case DupcheckDisabled:
		return true
	}

	return false
}

// Get will return a poll with the specified ID
func Get(id int) (Poll, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/polls/%d", baseURL, id), nil)
	if err != nil {
		return Poll{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Poll{}, err
	}

	if res.StatusCode == 404 {
		return Poll{}, errors.New("the requested poll could not be found")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response Poll
	json.Unmarshal(body, &response)

	return response, nil
}

// Create will create a poll with the specified options
func Create(title string, options []string, multi bool, dupcheck string, captcha bool) (Poll, error) {
	if len(title) == 0 {
		return Poll{}, errors.New("you must provide a valid title")
	}

	if len(options) < 2 {
		return Poll{}, errors.New("you must provide at least 2 options")
	}

	if !isValidDupcheck(dupcheck) {
		dupcheck = DupcheckNormal
	}

	poll := Poll{
		Title:    title,
		Options:  options,
		Multi:    multi,
		Dupcheck: dupcheck,
		Captcha:  captcha,
	}

	jsonPayload, _ := json.Marshal(poll)
	payload := strings.NewReader(string(jsonPayload))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", baseURL, "polls"), payload)

	if err != nil {
		return Poll{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return Poll{}, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response Poll
	json.Unmarshal(body, &response)

	return response, nil
}
