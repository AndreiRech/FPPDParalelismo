// por Fernando Dotti - PUCRS -
// aqui utilizamos como tarefa "cpu intensiva" calcular um valor
// primo com um determinado número de casas.
// em "casas" define-se tamanhos de 3, 6, 9, 12, 15, 18 casas
// em seguida, mede-se o tempo para gerar um valor primo com cada número de casas,
// com a função timeToGenPrime(...)
// a geração de um número primo é,
//    sorteia um valor com o numero de casas,
//    verifica se é primo
// repete até achar um primo
// diferentes execuções, iniciando com diferentes seeds, terão diferentes tempos.

// o trabalho "genPrime(tamanho)" representa uma computação intensiva,
// que é mais custosa quanto o tamanho do primo a ser gerado.

// suponha que voce tem que gerar um conjunto de N valores primos.
// calcule o speedup de uma solução paralela com P núcleos processadores.
// faca uma análise de speedup para os diferentes tamanhos de valores primos.
// Exemplos para
//   N: 50;  ou mais
//   P:  conforme seu hardware {2, 4, 6, 8, 10, 12}
//   tamanhos dos valores, coforme exemplificado no programa abaixo.

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	casas := [6]int{
		999,
		999999,
		999999999,
		999999999999,
		999999999999999,
		999999999999999999,
	}

	processCounts := []int{1, 2, 4, 6, 8, 10, 12, 14, 16}
	rand.Seed(69)

	N := 10

	for _, procs := range processCounts {
		fmt.Printf("Testando com %d processadores:\n", procs)
		runtime.GOMAXPROCS(procs)

		for _, tam := range casas {
			start := time.Now()

			timeToGenPrime(N, tam, procs)

			fmt.Printf("Tamanho: %d | Tempo: %v\n", tam, time.Since(start))
		}

		fmt.Println()
	}
}

func genPrime(tam int) {
	notPrimo := true
	v := 0
	for notPrimo {
		v = rand.Intn(tam)
		notPrimo = !isPrime(v)
	}
	//fmt.Printf("Primo: %d\n", v)
}

func timeToGenPrime(N, tam, procs int) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, procs)

	for i := 0; i < N; i++ {
		wg.Add(1)
		ch <- struct{}{}

		go func() {
			defer wg.Done()
			genPrime(tam)
			<-ch
		}()
	}

	wg.Wait()
}

// Is p prime?``
func isPrime(p int) bool {
	if p%2 == 0 {
		return false
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}
