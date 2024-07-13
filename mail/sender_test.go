package mail

import (
	"testing"

	"github.com/Ali-Gorgani/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping send gmail test in short mode")
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)
	
	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Test email"
	content := `
	<h1>Hello Gmail</h1>
	<p>This is a test email.</p>
	`
	to := []string{"onepiece199922@gmail.com"}
	attachFiles := []string{"../doc/swagger/simple_bank.swagger.json"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}