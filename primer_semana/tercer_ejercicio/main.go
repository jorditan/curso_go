package main

import (
	correo "ejercicio_saludo_go/tercer_ejercicio/correo"
)

func main() {
	correo.EnviarMailCompleto("Damian", "Aumenyto yoga", "maldito milei")
	correo.EnviarMailConDestinatario("Matias")
	correo.EnviarMailConDestinatarioYAsunto("Andrea", "Pensión habilitada")
}
