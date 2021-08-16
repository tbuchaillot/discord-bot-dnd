package session

import (
	"errors"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/tbuchaillot/discord-bot-dnd/chars"
	"github.com/tbuchaillot/discord-bot-dnd/store"
)

type Session struct {
	Partida       string
	Active        map[string]chars.Character
	StartTime     time.Time
	LastStartTime time.Time

	activeMu sync.Mutex
	store    *store.Storage
	once     sync.Once
}

func NewSession() *Session {
	session := &Session{}
	session.activeMu = sync.Mutex{}
	session.Active = make(map[string]chars.Character)
	session.store = &store.Storage{}
	session.LastStartTime = time.Now()
	session.StartTime = time.Now()

	session.store.Init()

	session.store.Db.AutoMigrate(&Session{})

	return session
}

func (s *Session) LoadSession(name string) error {
	sess := &Session{}
	result := s.store.Db.Last(sess, "partida = ?", name)

	if result.Error != nil {
		log.Error("Error getting session from db:", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("Session not found")
	}

	sess.store = s.store
	*s = *sess
	return nil
}

func (s *Session) CreateSession() error {
	result := s.store.Db.Create(s)
	if result.Error != nil {
		log.Error("Error saving session to db:", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("session not saved. Error")
	}
	return nil
}

func (s *Session) StoreSession() error {
	if s.store.Db == nil {
		log.Error("la db es nil :O")
		return errors.New("Error usando la base de datos")
	}
	result := s.store.Db.Model(&Session{}).Where("partida = ?", s.Partida).Save(s)
	if result.Error != nil {
		log.Error("Error updating session to db:", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("session not updated. Error")
	}
	return nil
}

func (s *Session) IsValid(discordSession *discordgo.Session, m *discordgo.MessageCreate) bool {
	if s.Partida == "" || s.store == nil {
		discordSession.ChannelMessageSend(m.ChannelID, ":warning: No estas jugando ninguna partida. Prueba con **!ayuda** o **!help** ")
		return false
	}

	if _, ok := s.Active[m.Author.ID]; !ok {
		discordSession.ChannelMessageSend(m.ChannelID, ":warning: No tienes ningun personaje en esta partida. Prueba con **!ayuda** o **!help**")
		return false
	}
	return true
}
