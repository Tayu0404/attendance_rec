package main

import (
	"os"
	"fmt"
	"time"
	"log"
	"strings"
	"github.com/bwmarrin/discordgo"

	"github.com/Tayu0404/attendance_rec/discord_bot/modules"
)

var (
	attendance = "k!chikoku"
	absence = "k!mitei"
	tardiness = "!tardiness"
	help = "k!help"
)

//db connect
db, _ := sqlx.Connect("mysql",
	"attendance_rec:@tcp(db:3306)/attendance_rec_db")

func main() {
	var (
		Token = os.Getenv("BOT_TOKEN")
		stopBot = make(chan bool)
	)
	fmt.Println(Token)
	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error Logging in")
		fmt.Println(err)
	}
	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listening...")
	<-stopBot
	return
}

func onMessageCreate(s *discordgo.Session,m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ",err)
		return
	}
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	switch {
		case strings.HasPrefix(m.Content, fmt.Sprintf(attendance)):
			sendMessage(s, c, "test")
		case strings.HasPrefix(m.Content, fmt.Sprintf(absence)):
			sendMessage(s, c, "遅刻")
		case strings.HasPrefix(m.Content, fmt.Sprintf(tardiness)):
			sendMessage(s, c, "未定")
		case strings.HasPrefix(m.Content, fmt.Sprintf(help)):
			sendMessage(s, c, "help予定")
		}
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
