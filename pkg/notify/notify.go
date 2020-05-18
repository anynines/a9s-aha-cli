package notify

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

const slackUsername = "a9s-aha-cli"

func Send(msg_text string) error {
	attachment := slack.Attachment{
		Text:       msg_text,
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Username:    slackUsername,
		Attachments: []slack.Attachment{attachment},
	}

	return slack.PostWebhook(os.Getenv("SLACK_URL"), &msg)
}
