package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/walkure/stdslack/util"
)

var channels = flag.String("channels", "", "channel ids split with comma")
var comment = flag.String("comment", "", "comments attached to posted file")
var raw = flag.Bool("raw", false, "send data as raw")
var filename = flag.String("filename", "stdin.txt", "filename recognized by slack")

func main() {
	flag.Parse()

	bytesInput, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("cannot read stdin: %w\n", err)
		return
	}

	token := os.Getenv("SLACK_TOKEN")

	if token == "" {
		fmt.Fprintf(os.Stderr, "env:SLACK_TOKEN required.\n")
		os.Stdout.Write(bytesInput)
		return
	}

	client := slack.New(token)

	if *channels == "" {
		fmt.Fprintf(os.Stderr, "channels argument is mandatory\n")
		os.Stdout.Write(bytesInput)
		return
	}

	var input io.Reader

	if *raw {
		input = bytes.NewReader(bytesInput)
	} else {
		input, err = util.ToUTF8(bytesInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failure to convert input: %v\n", err)
			os.Stdout.Write(bytesInput)
			return
		}
	}

	err = uploadContent(context.Background(), client, input, strings.Split(*channels, ","), *comment, *filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failure to send slack: %v\n", err)
		os.Stdout.Write(bytesInput)
		return
	}

}

// uploadContent upload content
func uploadContent(ctx context.Context, api *slack.Client, content io.Reader, channels []string, comment, filename string) error {

	params := slack.FileUploadParameters{
		Channels:       channels,
		InitialComment: comment,
		Filename:       filename,
		Reader:         content,
	}

	_, err := api.UploadFileContext(ctx, params)

	if err != nil {
		return fmt.Errorf("failure to upload file: %w", err)
	}

	return nil
}
