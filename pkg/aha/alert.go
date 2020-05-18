package aha

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/anynines/a9s-aha-cli/pkg/notify"

	log "github.com/sirupsen/logrus"
)

type Content struct {
	Idea []struct {
		Title            string `json:"name"`
		Updated_at       string `json:"updated_at"`
		Reference_number string `json:"reference_num"`
		Workflow_status  struct {
			Name string `json:"name"`
		} `json:"workflow_status"`
	} `json:"ideas"`
}

type StaleIdea struct {
	Title            string
	Reference_number string
	Url              string
}

func Stale(notifyFlag bool, verboseFlag bool) error {
	var status Content
	var ideas []StaleIdea

	if verboseFlag {
		log.SetLevel(log.DebugLevel)
	}

	err := CheckEnvVariables(notifyFlag)
	if err != nil {
		return err
	}

	log.Debug("Querying aha.io view...")
	err = fetchData(&status)
	if err != nil {
		return err
	}
	log.Debug("Queried aha.io view.")

	for i, _ := range status.Idea {
		idea := ReturnIdeaIfStale(status, i)
		if (StaleIdea{}) != idea {
			ideas = append(ideas, idea)
		}
	}

	return printReport(notifyFlag, ideas)
}

func CheckEnvVariables(notifyFlag bool) error {
	keys := []string{"AHA_USERNAME", "AHA_PASSWORD"}
	if notifyFlag {
		keys = append(keys, "SLACK_URL")
	}

	for _, key := range keys {
		if len(os.Getenv(key)) < 1 {
			return fmt.Errorf("You MUST set the environment variable %s", key)
		}
	}

	return nil
}

func ReturnIdeaIfStale(status Content, i int) StaleIdea {
	if status.Idea[i].Workflow_status.Name != "Needs review" {
		return StaleIdea{}
	}

	if !CheckDate(string(status.Idea[i].Updated_at)) {
		return StaleIdea{}
	}

	url := fmt.Sprintf("https://anynines.aha.io/api/v1/ideas/%s?fields=updated_at,name", status.Idea[i].Reference_number)
	return StaleIdea{status.Idea[i].Title, status.Idea[i].Reference_number, url}
}

func CheckDate(date string) bool {
	date_layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(date_layout, date)
	two_weeks_ago := ((time.Now().Add(-336 * time.Hour)).Unix())

	if t.Unix() < two_weeks_ago {
		return true
	}

	return false
}

func printReport(notifyFlag bool, ideas []StaleIdea) error {
	message := "Ideas: \n"
	for i, _ := range ideas {
		if notifyFlag {
			message += fmt.Sprintf(" <%s|%s> \n", "https://anynines.aha.io/ideas/ideas/"+ideas[i].Reference_number, ideas[i].Title)
		} else {
			message += fmt.Sprintf("%s (%s)\n", ideas[i].Title, "https://anynines.aha.io/ideas/ideas/"+ideas[i].Reference_number)
		}
	}
	message += "\nare in 'Needs review' state and haven't been updated since two weeks!"
	fmt.Println(message)

	if notifyFlag {
		log.Debug("Sending alert to Slack...")
		err := notify.Send(message)
		if err != nil {
			return err
		}
		log.Debug("Sent alert to Slack.")
	}

	return nil
}

func fetchData(status *Content) error {
	url := "https://anynines.aha.io/api/v1/ideas?fields=updated_at,name,workflow_status,reference_num"
	body, err := HttpGet(url, os.Getenv("AHA_USERNAME"), os.Getenv("AHA_PASSWORD"))
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(body), &status)
}
