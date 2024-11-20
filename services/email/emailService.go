package email

import (
	generalUtilities "SleekSpace/utilities/funcs/general"
	"net/smtp"
)

func emailForVerificationCodeHtml(firstName string, verificationCode string) string {
	return ("<h2>Hi " + firstName + "</h2> <p>Welcome to <span style=\"color: rgb(77, 9, 205);font-size: 14px;font-weight:bold;\">SleekSpace</span>. Please enter this code to activate your account: " + verificationCode + "</p></p> This code will expire in 15 minutes. See you soon.</p>")
}

func SendVerificationCodeEmail(email string, firstName string, verificationCode string) bool {
	html := emailForVerificationCodeHtml(firstName, verificationCode)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: Confirm your email for Account Registration\n" + headers + "\n\n" + html
	auth := smtp.PlainAuth(
		"",
		generalUtilities.GetEnvVariables().Email,
		generalUtilities.GetEnvVariables().EmailPassword,
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		generalUtilities.GetEnvVariables().Email,
		[]string{email},
		[]byte(msg),
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
