package command

import (
	"github.com/tbuchaillot/discord-bot-dnd/chars"
)

const STARTSESSIONCMD = "!empezar"
const STARTSESSIONCMD_HELP = "Como se usa? \n **!empezar nombre_partida**"

const STOPSESSIONCMD = "!terminar"

const ROLLCMD = "!roll"
const ROLLCMD_HELP = "Como se usa? \n **!rollX** Y \nDonde X=numero de caras Y=numero de dados"

const SPELLCMD = "!spell"
const SPELLCMD_HELP = "Como se usa? !spell"

const CREATECHAR = "!crear-pj"
const CREATECHAR_HELP = "Como se usa? \n **!crear-pj **\n ```" + chars.Example + "```"

const GETCHARCMD = "!personaje"

const HELLOCMD = "!hola"

const HELPCMD_1 = "!help"
const HELPCMD_2 = "!ayuda"
