package command

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tbuchaillot/discord-bot-dnd/session"

	"github.com/bwmarrin/discordgo"
)

//RollHandler !roll command rollea un dado de
func RollHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().Unix())

	valid, dices, faces := getDiceInfo(m.Content)
	if !valid {
		msg := ":x: Error con el commando: " + m.Content + " \n" + ROLLCMD_HELP
		s.ChannelMessageSend(m.ChannelID, msg)
		return
	}
	msg := "**" + sess.Active[m.Author.ID].Nombre + "** tiro **" + fmt.Sprint(dices) + "** dado"
	if dices > 1 {
		msg += "s"
	}
	msg += ":game_die:..."
	s.ChannelMessageSend(m.ChannelID, msg)

	if dices >= 10 {
		msg = ":x: " + m.Author.Username + " esta re loco y quiere tirar demasiados dados..."
		s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	for i := 1; i <= dices; i++ {
		randDado := rand.Intn(faces) + 1

		msg := ":game_die: El dado " + fmt.Sprint(i) + " salio **" + fmt.Sprint(randDado) + "**"

		randomSleep := rand.Intn(i * 500)
		time.Sleep(time.Duration(randomSleep * int(time.Millisecond)))
		s.ChannelMessageSend(m.ChannelID, msg)
	}

}

func getDiceInfo(cmd string) (bool, int, int) {
	var valid bool
	var numFaces, numDices int
	numDices = 1
	msg := cmd
	msgWithoutPrefix := strings.TrimPrefix(msg, ROLLCMD)
	msgSlice := strings.Split(msgWithoutPrefix, " ")
	if len(msgSlice) < 1 {
		log.Error("Intentando usar !roll sin numero")
		valid = false
		return valid, numDices, numFaces
	}

	//Faces
	if len(msgSlice) >= 1 {
		firstElement := msgSlice[0]

		faces, err := strconv.Atoi(firstElement)
		if err != nil {
			log.Error(firstElement + " no es un numero :(")
			valid = false
			return valid, numDices, numFaces
		}
		numFaces = faces
	}

	if len(msgSlice) >= 2 {
		numOfDados := msgSlice[1]

		dices, err := strconv.Atoi(numOfDados)
		if err != nil {
			log.Error(numOfDados + " no es un numero :(")
			valid = false
			return valid, numDices, numFaces
		}

		if dices == 0 {
			dices = 1
		}
		numDices = dices
	}

	valid = true
	return valid, numDices, numFaces
}
