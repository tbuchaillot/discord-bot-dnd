module github.com/tbuchaillot/discord-bot-dnd

go 1.15

require (
	github.com/bwmarrin/discordgo v0.23.2
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.7
)
replace gorm.io/gorm => github.com/TykTechnologies/gorm v1.20.7-0.20210409171139-b5c340f85ed0
