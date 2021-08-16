package command

import (
	"encoding/json"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tbuchaillot/discord-bot-dnd/chars"
	"github.com/tbuchaillot/discord-bot-dnd/session"
)

func CreateCharHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	msgWithoutPrefix := strings.TrimPrefix(cmd, CREATECHAR)
	msgSlice := strings.Split(msgWithoutPrefix, " ")

	if len(msgSlice) < 2 {
		s.ChannelMessageSend(m.ChannelID, ":x: Error creando personaje. \n "+CREATECHAR_HELP)
		return
	}

	charData := strings.TrimPrefix(cmd, CREATECHAR)
	pj := &chars.Character{}
	err := json.Unmarshal([]byte(charData), &pj)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":x: Error leyendo los datos del personaje.")
		return
	}

	sess.Active[m.Author.ID] = *pj
	sess.StoreSession()

	s.ChannelMessageSend(m.ChannelID, ":white_check_mark: **"+pj.Nombre+"** creado exitosamente.")
	s.ChannelMessageSend(m.ChannelID, pj.Prettify())
}

func GetCharHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	if pj, ok := sess.Active[m.Author.ID]; !ok {
		s.ChannelMessageSend(m.ChannelID, ":warning: No tienes ningun personaje en esta partida. Puedes crear un personaje usando: \n "+CREATECHAR_HELP)
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, pj.Prettify())
	}
}
