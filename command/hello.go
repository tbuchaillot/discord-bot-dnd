package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tbuchaillot/discord-bot-dnd/session"
)

func HelloHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {

	if sess.Partida == "" {
		s.ChannelMessageSend(m.ChannelID, ":warning: No estas jugando ninguna partida. Puedes crear partida usando: \n "+STARTSESSIONCMD_HELP)
		return
	}

	if pj, ok := sess.Active[m.Author.ID]; !ok {
		s.ChannelMessageSend(m.ChannelID, ":warning: No tienes ningun personaje en esta partida. Puedes crear un personaje usando: \n "+CREATECHAR_HELP)
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, "Hola **"+pj.Nombre+"**! \n")
	}
}
