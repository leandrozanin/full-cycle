package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma esperado é 30, resultado %d", total)
	}
}

func TestSoma2(t *testing.T) {
	total := Soma(20, 20)

	if total != 20 {
		t.Errorf("Resultado da soma esperado é 40, resultado %d", total)
	}
}
