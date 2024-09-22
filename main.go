package main

import (
	"flag"
	"fmt"
	// "fmt"

	"log"
	"os"

	"github.com/slack-go/slack"
)

var debug bool

func init() {
    flag.BoolVar(&debug, "debug", false, "Enable debug logging")
    flag.Parse()
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "BOT_TOKEN")
	os.Setenv("CHANNEL_ID", "CHANNEL_ID")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	fileArr := []string{"Apache_Avro_Logo_2023.svg.png", "GitHub_logo_2013.webp", "github-logo-vector.png", "ios-weather.svg", "png-clipart-computer-icons-logo-github-github-logo-logo-computer-program.png"}
	for i := 0; i < len(fileArr); i++ {
		file, _ := os.Open(fileArr[i])
		fileInfo, _ := file.Stat()
		params := slack.UploadFileV2Parameters{
			Channel: os.Getenv("CHANNEL_ID"),
			Filename: fileArr[i],
			FileSize: int(fileInfo.Size()),
			Reader: file,
		}
		uploaded, err := api.UploadFileV2(params)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Name: %s, URL: %s", uploaded.ID, uploaded.Title)

	}
}
