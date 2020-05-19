# SMTP for offce365

Package to easily handle office 365 authentication and message


## Example code

```go
package main

import (
	"net/smtp"
	"github.com/jhw0604/smtp/office365"
)

func main() {
	auth := office365.StartTLSAuth("account@office365.com", "P@ssword")
	msg := office365.NewMessage(
		"this is title",
		"from@office365.com",
		[]string{"to1@office365.com", "to2@office365.com"},
		[]string{"cc1@office365.com"},
		nil,
		"text/html",
		"email with html tag<br/><table><tr><td>col1</td><td>col2</td></tr><tr><td>a1</td><td>b1</td></tr><tr><td>a2</td><td>b2</td></tr></table>",
    )
    
	err := smtp.SendMail("smtp.office365.com:587", auth, msg.From(), msg.To(), msg.Byte())
	if err != nil {
		panic(err)
	}
}
```