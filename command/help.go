package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tbuchaillot/discord-bot-dnd/session"
)

func HelpHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "**Comandos disponibles** \n" +
		"**" + HELLOCMD + "**\n" +
		"\n" +
		"**" + STARTSESSIONCMD + "**\n" +
		"**" + STOPSESSIONCMD + "**\n" +
		"\n" +
		"**" + CREATECHAR + "**\n" +
		"**" + GETCHARCMD + "**\n" +
		"\n" +
		"**" + ROLLCMD + "**\n" +
		"**" + SPELLCMD + "**\n"

	s.ChannelMessageSend(m.ChannelID, msg)
}
