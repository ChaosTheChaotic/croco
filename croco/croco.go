package croco

import (
	"bytes"
	"io"
	"os"

	"github.com/schollz/croc/v10/src/croc"
	"github.com/schollz/croc/v10/src/models"
)

func Send(msg string, code string) error {
	opts := croc.Options{
		SharedSecret:   code,
		IsSender:       true,
		RelayAddress:   models.DEFAULT_RELAY,
		RelayPassword:  models.DEFAULT_PASSPHRASE,
		RelayPorts:     []string{"9009", "9010", "9011", "9012", "9013"},
		NoPrompt:       true,
		DisableLocal:   false,
		SendingText:    true,
		Curve: "p256",
	}

	sender, err := croc.New(opts)
	if err != nil {
		return err
	}

	// Create a temporary file with the text content
	tmpFile, err := os.CreateTemp("", "croc-text-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	// Write the message to the temporary file
	if _, err := tmpFile.Write([]byte(msg)); err != nil {
		return err
	}
	tmpFile.Close()

	// Get file information for the temporary file
	filesInfo, emptyFolders, totalNumberFolders, err := croc.GetFilesInfo([]string{tmpFile.Name()}, false, false, nil)
	if err != nil {
		return err
	}

	return sender.Send(filesInfo, emptyFolders, totalNumberFolders)
}

func Recv(code string) (string, error) {
	opts := croc.Options{
		SharedSecret:   code,
		IsSender:       false,
		RelayAddress:   models.DEFAULT_RELAY,
		RelayPassword:  models.DEFAULT_PASSPHRASE,
		RelayPorts:     []string{"9009", "9010", "9011", "9012", "9013"},
		NoPrompt:       true,
		DisableLocal:   false,
		Curve: "p256",
	}

	recipient, err := croc.New(opts)
	if err != nil {
		return "", err
	}

	// Create a pipe to capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the receive operation
	err = recipient.Receive()

	// Close the pipe and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Return the captured text and any error
	return buf.String(), err
}
