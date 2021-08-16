package command

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/tbuchaillot/discord-bot-dnd/session"
)

func SpellHandler(sess *session.Session, s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().Unix())

	poder := rand.Intn(20) + 1
	control := rand.Intn(20) + 1

	msg := "**" + sess.Active[m.Author.ID].Nombre + " tiro una habilidad!**"
	s.ChannelMessageSend(m.ChannelID, msg)
	time.Sleep(time.Duration(rand.Intn(3) * 500 * int(time.Millisecond)))
	msg = ":fire: **Poder:**" + fmt.Sprint(poder) + "   :magic_wand:" + " **Control:**" + fmt.Sprint(control) + "."
	s.ChannelMessageSend(m.ChannelID, msg)

	pj := sess.Active[m.Author.ID]
	if pj.Poder > 0 && pj.Control > 0 {
		if poder < pj.Poder && control < pj.Control {
			msg = "**ENTRA EL PODER** :mage: "
		} else {
			msg = "**NO ENTRA EL PODER** :face_with_spiral_eyes:"
			msg += "\n **" + pj.Nombre + "** tiene **" + fmt.Sprint(pj.Poder) + " y " + fmt.Sprint(pj.Control) + "** de **poder y control** pero salio **" + fmt.Sprint(poder) + " y " + fmt.Sprint(control) + " **"
		}
		s.ChannelMessageSend(m.ChannelID, msg)
	}

}
