package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }

// Email sender configuration
// NOTE: never put passwords into source code!

const sender = "notifications@example.com"
const password = "correcthoresebatterystaple"
const hostname = "smpt.example.com"
const template = `Warning: you are using %d bytes of storage,%d%% of your quota.`

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))

	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

// CheckQuota checks to see if a user is using 90% or more of their disk allotment
func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 /// 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}

	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
