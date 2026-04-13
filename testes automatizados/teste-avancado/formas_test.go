package teste_avancado

import "testing"

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{Largura: 10, Altura: 12}
		areaEsperada := float64(120)
		areaREcebida := ret.Area()

		if areaEsperada != areaREcebida {
			//Pode ser usado Fatalf, porém para aqui os testes, com Errorf os testes continuam executando
			t.Errorf("A area recebida %f é diferente da esperada %f", areaREcebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{Raio: 10}
		areaEsperada := float64(314.1592653589793)
		areaREcebida := circ.Area()

		if areaEsperada != areaREcebida {
			t.Errorf("A area recebida %f é diferente da esperada %f", areaREcebida, areaEsperada)
		}
	})
}
