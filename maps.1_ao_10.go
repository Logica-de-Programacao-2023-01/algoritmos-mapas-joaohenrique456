package main

import (
	"fmt"
	"sort"
	"strings"
)

// q.1

func contarAlgarismos(s string) map[string]int {

	frequenciaAlgarismos := make(map[string]int)

	strings.Fields(s)

	for _, ranS := range strings.Fields(s) {

		frequenciaAlgarismos[ranS] += 1

	}

	return frequenciaAlgarismos

}

func main() {

	s := "Toma"

	fmt.Println(contarAlgarismos(s))

}

// q.2

func unirMapas(m1 map[string]int, m2 map[string]int) map[string]int {

	m3 := make(map[string]int)

	for range1, range2 := range m1 {

		m3[range1] = range2

	}

	for range3, range4 := range m2 {

		m3[range3] = range4

	}

	return m3

}

func main() {

	m1 := map[string]int{

		"macaco": 1,
		"aranha": 4,
		"kiko":   2,
	}

	m2 := map[string]int{

		"cavalo": 2,
		"rato":   9999,
		"1":      2,
	}

	fmt.Println(unirMapas(m1, m2))

}

// q.3

func somaMapa(m map[string]int) int {

	soma := 0

	for _, ranM := range m {

		soma += ranM

	}

	return soma

}

func main() {

	m := map[string]int{

		"a": 423,
		"b": 1212,
		"c": 212,
	}

	fmt.Println(somaMapa(m))

}

// q.4

func encontrarAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		sorted := sortString(palavra)

		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}

func sortString(s string) string {

	chars := strings.Split(s, "")

	sort.Strings(chars)

	return strings.Join(chars, "")
}

func main() {

	palavras := []string{"cara", "lindo", "demais", "sorte", "no", "jogo", "azar", "amor"}

	fmt.Println(encontrarAnagramas(palavras))

}

// q.5

func frequenciaLetras(s string) map[string]int {

	m1 := make(map[string]int)
	chars := strings.Split(s, "")

	for _, ranC := range chars {

		m1[ranC] += 1

	}

	return m1

}

func main() {

	fmt.Println(frequenciaLetras("comprei uma esfirra"))

}

// q.6

func contarPalavras2(s string) map[string]int {

	m1 := make(map[string]int)
	conj := strings.Fields(s)

	for _, ranConj := range conj {

		m1[ranConj] += 1

	}

	return m1

}

func unirMapas2(mapas []map[string]int) map[string]int {

	bigMap := make(map[string]int)

	for _, ranMapas := range mapas {

		for chave, valor := range ranMapas {

			bigMap[chave] += valor

		}

	}

	return bigMap

}

func main() {

	a := "muçilon"
	b := "adulto"
	c := "velho"

	conj := []map[string]int{contarPalavras2(a), contarPalavras2(b), contarPalavras2(c)}

	fmt.Println(unirMapas2(conj))

}

// q.7

func contarLetrasEPalavras(s string) map[string]map[string]int {

	mF := make(map[string]map[string]int)

	palavras := strings.Fields(s)

	chars := []string{}

	for _, ranP := range palavras {

		chars = strings.Split(ranP, "")

		mF[ranP] = make(map[string]int)

		for _, ranC := range chars {

			mF[ranP][ranC]++

		}

	}

	return mF

}

func main() {

	fmt.Println(contarLetrasEPalavras("comprei uma vitamina"))

}

// q.8

func igualarDespesas(m1 map[string]float64) {

	maior := 0.0

	for _, ranM1 := range m1 {

		if ranM1 > maior {

			maior = ranM1

		}

	}

	for nome, quantia := range m1 {

		quantia = maior - quantia

		fmt.Printf("%s deve pagar %.2f reais para  igualar.\n", nome, quantia)

	}

}

func main() {

	m1 := make(map[string]float64)
	m1["Repet"] = 10000
	m1["Jubis"] = 100
	m1["Melo"] = 500

	igualarDespesas(m1)

}

// q.9

func fibonacci(n int) map[int]int {

	fibo := make(map[int]int)

	for i := 0; len(fibo) != n; i++ {

		if i < 2 {

			fibo[i] = i

		} else {

			fibo[i] = fibo[i-2] + fibo[i-1]

		}

	}

	return fibo

}

func main() {

	fmt.Println(fibonacci(20))

}

// q.10

func pares(s []int) map[int]int {

	frequenciaS := make(map[int]int)

	for _, ranS := range s {

		frequenciaS[ranS] += 1

	}

	for numero, frequencia := range frequenciaS {

		if frequencia%2 != 0 {

			delete(frequenciaS, numero)

		} else {

			frequenciaS[numero] /= 2

		}

	}

	return frequenciaS

}

func main() {

	s := []int{0, 2, 3, 1, 2, 1, 2, 5, 5, 3, 4, 4, 4, 4}

	fmt.Println(pares(s))

}
