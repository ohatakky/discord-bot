package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BOT Running...")

	// ctrl + c
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	mention := member.Mention()
	_, err = s.ChannelMessageSend(os.Getenv("CHANNEL_ID"), mention)
	if err != nil {
		log.Fatal(err)
		return
	}
}
