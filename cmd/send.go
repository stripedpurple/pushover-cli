/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/viruscmd/pushover-cli/pkg"
	"github.com/gregdel/pushover"
	"github.com/spf13/cobra"
	"log"
	"time"
	"fmt"
)

// sendCmd represents the send command
var (
	Title       string
	Message     string
	Priority    int
	URL         string
	URLTitle    string
	Timestamp   int64
	Retry       time.Duration
	Expire      time.Duration
	DeviceName  string
	CallbackURL string
	Sound       string
	HTML        bool
	sendCmd     = &cobra.Command{
		Use:   "send",
		Short: "Sends push notification to pushover",
		Long: `Sends push notification to specified or default a pushover app. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			sendNotification()
		},
	}
)


func init() {
	// Adds send command to rootCmd
	rootCmd.AddCommand(sendCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	sendCmd.Flags().StringVarP(&Title, "title", "t", "", "Title to appear in push notification")
	sendCmd.Flags().StringVarP(&Message, "message", "m", "", "The message to appear in the body of the push notification")
	sendCmd.Flags().IntVarP(&Priority, "priority", "p", 0, "Priority of push notification. Default value is 0, and must be an integer between -2 and 2")
	sendCmd.Flags().StringVarP(&URL, "url", "u", "", "URL to be delivered with push notification")
	sendCmd.Flags().StringVarP(&URLTitle, "url-title", "U", "", "URL to be delivered with push notification")
	sendCmd.Flags().Int64VarP(&Timestamp, "timestamp", "T", _now(), "Priority of push notification. Default value is 0, and must be an integer between -2 and 2")
	sendCmd.Flags().DurationVarP(&Retry, "retry", "r", 0, "Priority of push notification. Default value is 0, and must be an integer between -2 and 2")
	sendCmd.Flags().DurationVarP(&Expire, "expire", "e", 0, "Priority of push notification. Default value is 0, and must be an integer between -2 and 2")
	sendCmd.Flags().StringVarP(&CallbackURL, "callback-url", "c", "", "URL to be delivered with push notification")
	sendCmd.Flags().StringVarP(&DeviceName, "device", "d", "", "the name of the device to send the push notification to")
	sendCmd.Flags().StringVarP(&Sound, "sound", "s", "", "Sets notification sound")
	sendCmd.Flags().BoolVarP(&HTML, "html", "H", true, "Tell pushover that the message body is html")
}

func constructNotification() pushover.Message {
	msg := pushover.Message{
		Message: Message,
	}

	if &Title != nil {
		msg.Title = Title
	}

	if &Priority != nil {
		msg.Priority = Priority
	}

	if &URL != nil {
		msg.URL = URL
	}

	if &URLTitle != nil {
		msg.URLTitle = URLTitle
	}

	if &Timestamp != nil {
		msg.Timestamp = Timestamp
	}

	if &Retry != nil {
		msg.Retry = Retry
	}

	if &Expire != nil {
		msg.Expire = Expire
	}

	if &DeviceName != nil {
		msg.DeviceName = DeviceName
	}

	if &Timestamp != nil {
		msg.CallbackURL = CallbackURL
	}

	if &Sound != nil {
		msg.Sound = Sound
	}

	if &HTML != nil {
		msg.HTML = HTML
	}

	return msg
}

func sendNotification() {
	cfg, err := pkg.ReadConfigFile()

	// Create a new pushover app with a token
	app := pushover.New(cfg.ApplicationToken)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.UserToken)

	msg := constructNotification()
	response, err := app.SendMessage(&msg, recipient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func _now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
