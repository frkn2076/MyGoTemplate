package helper

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	// "html/template"
	// "bytes"
	"strings"

	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/cache"
  
	gomail "gopkg.in/mail.v2"
  )

  func SendMail(to string, key string, language string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "FindYourSoulMatee@gmail.com")
  
	// Set E-Mail receivers
	m.SetHeader("To",to)

	htmlFile, err := ioutil.ReadFile("helper/mailTemplate.html")
	if err != nil {
		logger.ErrorLog("An error occured while reading mailTemplate.html file - SendMail - emailSender.go ", err.Error())
	}

	html := string(htmlFile)

	mailHeader := cache.Get("MAILHEADER" + language)
	mailText := cache.Get("MAILTEXT" + language)
	mailSubject := cache.Get("MAILSUBJECT" + language)

	html = strings.Replace(html, "MAILKEY", key, -1)
	html = strings.Replace(html, "MAILHEADER", mailHeader, -1)
	html = strings.Replace(html, "MAILTEXT", mailText, -1)

	// Set E-Mail subject
	m.SetHeader("Subject", mailSubject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", html)
  
	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "FindYourSoulMatee@gmail.com", "fbr01994")
  
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
	  fmt.Println(err)
	  panic(err)
	}
  }