package main

import "fmt"

func validarSudoku(tabuleiro [][]int) bool {
	// Verificar linhas
	for _, linha := range tabuleiro {
		if !validarGrupo(linha) {
			return false
		}
	}

	// Verificar colunas
	for j := 0; j < 9; j++ {
		coluna := []int{}
		for i := 0; i < 9; i++ {
			coluna = append(coluna, tabuleiro[i][j])
		}
		if !validarGrupo(coluna) {
			return false
		}
	}

	// Verificar regiões 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			regiao := []int{}
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					regiao = append(regiao, tabuleiro[i+x][j+y])
				}
			}
			if !validarGrupo(regiao) {
				return false
			}
		}
	}

	return true
}

func validarGrupo(grupo []int) bool {
	ocorrencias := make(map[int]bool)
	for _, valor := range grupo {
		if valor != 0 && ocorrencias[valor] {
			return false
		}
		ocorrencias[valor] = true
	}
	return true
}

func encontrarValoresIncorretos(tabuleiro [][]int) [][]int {
	valoresIncorretos := [][]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			valor := tabuleiro[i][j]

			if valor == 0 {
				continue
			}

			// Verificar linha
			if contarOcorrencias(tabuleiro[i], valor) > 1 {
				valoresIncorretos = append(valoresIncorretos, []int{i, j})
			}

			// Verificar coluna
			coluna := []int{}
			for x := 0; x < 9; x++ {
				coluna = append(coluna, tabuleiro[x][j])
			}
			if contarOcorrencias(coluna, valor) > 1 {
				valoresIncorretos = append(valoresIncorretos, []int{i, j})
			}

			// Verificar região 3x3
			regiaoI := (i / 3) * 3
			regiaoJ := (j / 3) * 3
			regiao := []int{}
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					regiao = append(regiao, tabuleiro[regiaoI+x][regiaoJ+y])
				}
			}
			if contarOcorrencias(regiao, valor) > 1 {
				valoresIncorretos = append(valoresIncorretos, []int{i, j})
			}
		}
	}

	return valoresIncorretos
}

func contarOcorrencias(slice []int, valor int) int {
	count := 0
	for _, v := range slice {
		if v == valor {
			count++
		}
	}
	return count
}

func exibirCelulasIncorretas(tabuleiro [][]int) {
	valoresIncorretos := encontrarValoresIncorretos(tabuleiro)

	if len(valoresIncorretos) == 0 {
		fmt.Println("O tabuleiro não contém células com valores incorretos.")
	} else {
		fmt.Println("Células com valores incorretos:")
		for _, celula := range valoresIncorretos {
			i, j := celula[0], celula[1]
			fmt.Printf("Célula (%d, %d)\n", i, j)
		}
	}
}

func main() {
	tabuleiro := [][]int{
		{0, 0, 0, 0, 9, 0, 0, 7, 9},
		{4, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 2, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
	}

	//tabuleiro correto
	//tabuleiro := [][]int{
	//	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	//	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	//	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	//	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	//	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	//	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	//	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	//	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	//	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	//}

	exibirCelulasIncorretas(tabuleiro)
}
