# GoMail - golang Mailing Wrapper

## Usage

```golang
package main

import (
	. "./mailer"
)

func main() {
	mailer := New(
		"some-SMTP",
		587,
		"some-email",
		"some-pw",
	)

	mailer.SendMail([]string{"some-email-again"}, "Test", "Test Message")

}

```