package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		//Para se ter a substituição do valores dentro de uma string se faz assim.
		t.Errorf("resultado da soma é inválido: Resultado %d. Esperado %d", total, 30)
	}
}
