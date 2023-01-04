package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"strings"

	//"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
)

func main() {
    flag.Parse()

    if flag.NArg() < 2 {
        fmt.Println("Error: incorrect number of arguments")
        return
    }

    username := flag.Arg(0)
    content := strings.Join(flag.Args()[1:], " ")

	fmt.Println(username)
	fmt.Println(content)

	var token = os.Getenv("DISCORD_TOKEN")
	s, err := session.New(token)
	if err != nil { log.Fatalln("Session failed:", err) }

	sent := false
	channels, err := s.PrivateChannels();
	for i := 0; i < len(channels); i++ {
		if channels[i].DMRecipients[0].Username == username {
			s.SendMessage(channels[i].ID, content, nil)
			sent = true
			break
		}
	}
	if sent {
		fmt.Printf("Sent '"+content+"' to "+username+"\n")
	} else {
		fmt.Println("Failed to send")
	}

	if err := s.Open(); err != nil { log.Fatalln("Failed to connect:", err) }
	defer s.Close()
}