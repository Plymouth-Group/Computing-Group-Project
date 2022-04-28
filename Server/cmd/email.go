// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Email Sent

package main

import (
	"fmt"
  "log"
  "github.com/mailjet/mailjet-apiv3-go"
)

func Sent_Email(from_email string, from_name string, to_email string, to_name string, mail_subject string, mail_content string) {
	mailjetClient := mailjet.NewMailjetClient("4a265e4d06e373221647d92829bf2ef5", "02f4046503fbc58a3f4289a68815bc71")

	messagesInfo := []mailjet.InfoMessagesV31 {
    {
      From: &mailjet.RecipientV31{
        Email: from_email,
        Name: from_name,
      },
      To: &mailjet.RecipientsV31{
        mailjet.RecipientV31 {
          Email: to_email,
          Name: to_name,
        },
      },
      Subject: mail_subject,
      TextPart: mail_content,
    },
  }

	messages := mailjet.MessagesV31{Info: messagesInfo }
  _, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
    log.Fatal(err)
  }

	fmt.Println("> Email sent successfull: " + to_email)
}