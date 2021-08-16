package chars

import "fmt"

type Character struct {
	Nombre    string `json:"nombre"`
	Atributos struct {
		Fuerza   int `json:"fuerza"`
		Agilidad int `json:"agilidad"`
	} `json:"atributos"`
	Percepcion int `json:"percepcion"`
	Poder      int `json:"poder"`
	Control    int `json:"control"`
}

func (c *Character) Prettify() string {
	charStr := "===> **" + c.Nombre + "** <===" +
		"\n**Atributos**:" +
		"\n\t Fuerza:" + fmt.Sprint(c.Atributos.Fuerza) +
		"\n\t Agilidad:" + fmt.Sprint(c.Atributos.Agilidad) +
		"\n **Percepcion**:" + fmt.Sprint(c.Percepcion) +
		"\n **Poder**:" + fmt.Sprint(c.Poder) +
		"\n **Control**:" + fmt.Sprint(c.Control)

	return charStr
}

const Example = `{
    "nombre":"Nuevo Personaje",
    "atributos":{
        "fuerza":20,
        "agilidad":10
    },
    "percepcion":20,
    "poder":15,
    "control":3
}`
