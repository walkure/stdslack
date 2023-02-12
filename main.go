package main

import (
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

	token := os.Getenv("SLACK_TOKEN")

	if token == "" {
		fmt.Printf("env:SLACK_TOKEN required.")
		return
	}

	client := slack.New(token)

	if *channels == "" {
		fmt.Printf("channels argument is mandatory")
		return
	}

	input := io.Reader(os.Stdin)
	var err error

	if !*raw {
		input, err = util.ToUTF8(os.Stdin)
		if err != nil {
			fmt.Printf("failure to convert input: %v", err)
			return
		}
	}

	err = uploadContent(context.Background(), client, input, strings.Split(*channels, ","), *comment, *filename)
	if err != nil {
		fmt.Printf("failure to send slack: %v", err)
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
