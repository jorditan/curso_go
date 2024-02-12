package main_test

import (
	"main"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaludarConNombreVacio(t *testing.T) {
	saludo := main.Saludar("")
	assert.Equal(t, "Hola extraño", saludo)
}

func TestSaludarConNombreCompleto(t *testing.T) {
	saludo := main.Saludar("Matias")
	assert.Equal(t, "Hola matias", saludo)
}
