package techpalace

import (
	"strings"
)

const welcome = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	msg := welcome + strings.ToUpper(customer)
	return msg
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	msg := stars + "\n" + welcomeMsg + "\n" + stars
	return msg
}

// CleanupMessage cleans up an old marketing message.
// 1st: remove all stars
// 2nd: remove the leading and trailing whitespace from the remaning text
func CleanupMessage(oldMsg string) string {
	/*
		msg := strings.Trim(oldMsg, "*\n")
		msg = strings.TrimLeft(msg, " ")
		msg = strings.TrimRight(msg, " ")
	*/
	msg := strings.ReplaceAll(oldMsg, "*", "")
	msg = strings.TrimSpace(msg)
	return msg
}
