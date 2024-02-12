package correo

import (
	"fmt"
)

func enviarCorreo(destinatario string, asunto string, mensaje string) {
	if mensaje == "" {
		mensaje = "-"
	}

	if asunto == "" {
		asunto = "-"
	}

	if destinatario == "" {
		destinatario = "-"
	}

	fmt.Printf("enviando correo...\ndestinatario: %s\nasunto: %s\nmensaje: %s", destinatario, asunto, mensaje)
}

func EnviarMailConDestinatario(destinatario string) {
	enviarCorreo(destinatario, "", "")
}

func EnviarMailConDestinatarioYAsunto(destinatario string, asunto string) {
	enviarCorreo(destinatario, asunto, "")
}

func EnviarMailCompleto(destinanatrio string, asunto string, mensaje string) {
	enviarCorreo(destinanatrio, asunto, mensaje)
}
