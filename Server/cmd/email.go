// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Email Sending

package main

import (
	"fmt"
  "log"

  "github.com/mailjet/mailjet-apiv3-go"
)

// https://dev.mailjet.com/email/guides/getting-started/#send-your-first-email

func main_server_created(sent_email string, sent_name string) {
	email_body := fmt.Sprintf("Congratulations! your codinoc server has been created. Use \"%s\" as the server code for Sign In", glob_sign_up.server_code)

	mailjetClient := mailjet.NewMailjetClient("4a265e4d06e373221647d92829bf2ef5", "02f4046503fbc58a3f4289a68815bc71")

  messagesInfo := []mailjet.InfoMessagesV31 {
    {
      From: &mailjet.RecipientV31{
        Email: "10749846@students.plymouth.ac.uk",
        Name: "Codinoc Developers",
      },
      To: &mailjet.RecipientsV31{
        mailjet.RecipientV31 {
          Email: sent_email,
          Name: sent_name,
        },
      },
      Subject: "Welcome to Codinoc",
      TextPart: email_body,
    },
  }

	messages := mailjet.MessagesV31{Info: messagesInfo }
  res, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Data: %+v\n", res)
}

func password_reset_instruction_sent(server_code string, admin_email string) {
  reset_link := fmt.Sprintf("localhost:8080/reset_new?scode=%s&semail=%s", server_code, admin_email)
	email_body := fmt.Sprintf("Follow this link to reset your server's (%s) administrator account password. %s", server_code, reset_link)

	mailjetClient := mailjet.NewMailjetClient("4a265e4d06e373221647d92829bf2ef5", "02f4046503fbc58a3f4289a68815bc71")

  messagesInfo := []mailjet.InfoMessagesV31 {
    {
      From: &mailjet.RecipientV31{
        Email: "10749846@students.plymouth.ac.uk",
        Name: "Codinoc Developers",
      },
      To: &mailjet.RecipientsV31{
        mailjet.RecipientV31 {
          Email: admin_email,
          Name: server_code,
        },
      },
      Subject: "Password Reset Instructions",
      TextPart: email_body,
    },
  }

	messages := mailjet.MessagesV31{Info: messagesInfo }
  res, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Data: %+v\n", res)
}