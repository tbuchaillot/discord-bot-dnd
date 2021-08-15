package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/tbuchaillot/discord-bot-dnd/command"

	"github.com/bwmarrin/discordgo"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	//Creamos una Discord session usando el token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Error("Error creando la session de Discord:", err)
		return
	}

	//Agregamos el handler para cuando se crean mensajes
	dg.AddHandler(messageCreateHandler)

	// seteamos que el bot solo va a interactuar con mensajes
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	if err := dg.Open(); err != nil {
		log.Error("Error abriendo conexion:", err)
		return
	}

	log.Info("Bot corriendo :D. Ctrl+C para stopearlo")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	log.Info("bot stopeado D:")
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Si el mensaje es del bot, lo ignoramos
	if m.Author.ID == s.State.SessionID {
		return
	}

	cmd := m.Content
	if !strings.HasPrefix(cmd, "!") {
		return
	}

	switch {
	case strings.HasPrefix(cmd, command.ROLLCMD):
		command.RollHandler(s, m)
	default:
		log.Info("Command " + cmd + " no implementado")
	}
}
