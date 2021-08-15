package command

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
)

const ROLLCMD = "!roll"

// !roll command rollea un dado de
func rollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().Unix())

	diceFaces := getDiceFaces(m.Content)
	if diceFaces == 0 {
		s.ChannelMessageSend(m.ChannelID, "Error con el comando:"+m.Content)
		return
	}

	randomNum := rand.Intn(diceFaces)

	msg := m.Author.Username + " rolleo " + fmt.Sprint(randomNum)

	s.ChannelMessageSend(m.ChannelID, msg)
	return
}

func getDiceFaces(cmd string) int {

	msg := cmd
	msgWithoutPrefix := strings.TrimPrefix(msg, ROLLCMD)

	msgSlice := strings.Split(msgWithoutPrefix, " ")
	if len(msgSlice) < 1 {
		log.Error("Intentando usar !roll sin numero")
		return 0
	}

	firstElement := msgSlice[0]

	facesInt, err := strconv.Atoi(firstElement)
	if err != nil {
		log.Error("No es entero :(")
		return 0
	}

	return facesInt
}
