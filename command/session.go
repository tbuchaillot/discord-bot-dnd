package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/tbuchaillot/discord-bot-dnd/session"
)

func StartSessionHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	argCmd := strings.Split(cmd, " ")
	if len(argCmd) > 1 {
		err := sess.LoadSession(argCmd[1])
		if err != nil && (err.Error() == "Session not found" || err.Error() == "record not found") {
			sess.Partida = argCmd[1]
			errStore := sess.CreateSession()
			if errStore != nil {
				s.ChannelMessageSend(m.ChannelID, ":x: Error creando partida:"+errStore.Error())
				return
			}
			sess.LastStartTime = time.Now()
			s.ChannelMessageSend(m.ChannelID, ":white_check_mark: Partida **"+sess.Partida+"** creada y cargada correctamente!")
			return
		} else if err != nil {
			s.ChannelMessageSend(m.ChannelID, ":x: Error cargando partida:"+err.Error())
			return
		}

		sess.LastStartTime = time.Now()

		s.ChannelMessageSend(m.ChannelID, ":white_check_mark: Partida **"+sess.Partida+"** cargada correctamente!")
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, ":warning: "+STARTSESSIONCMD_HELP)
		return
	}

}

func EndSessionHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {

	errStore := sess.StoreSession()
	if errStore != nil {
		s.ChannelMessageSend(m.ChannelID, ":2: Error guardando la partida:"+errStore.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, ":white_check_mark: Partida **"+sess.Partida+"** guardada correctamente!")
	diff := time.Now().Sub(sess.LastStartTime)
	msg := ":information_source: Estuvieron viciando:**" + fmt.Sprintf("%v", diff) + "**"
	s.ChannelMessageSend(m.ChannelID, msg)
	diff = time.Now().Sub(sess.StartTime)
	msg = ":information_source: Total de vicio en esta partida:**" + fmt.Sprintf("%v", diff) + "**"
	s.ChannelMessageSend(m.ChannelID, msg)

}
