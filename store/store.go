package store

// /home/ec2-user/
import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

const DB_URL = "/home/ec2-user/db"

func (s *Storage) Init() {
	dialect := sqlite.Open(DB_URL)

	db, err := gorm.Open(dialect, &gorm.Config{
		AutoEmbedd:  true,
		UseJSONTags: true,
	})
	if err != nil {
		log.Error(err)
		return
	}
	s.Db = db
}
