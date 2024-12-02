package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func createRandomArray(n int) []int {
    arr := make([]int, n)
    for i := range arr {
        arr[i] = rand.Intn(1000)
    }
    return arr
}

func ParallelQuickSort(arr []int, wg *sync.WaitGroup) {
    defer wg.Done()
    if len(arr) <= 1 {
        return
    }

    pivot := partition(arr)
    var subWg sync.WaitGroup

    subWg.Add(2)
    go func() {
        ParallelQuickSort(arr[:pivot], &subWg)
    }()
    go func() {
        ParallelQuickSort(arr[pivot+1:], &subWg)
    }()

    subWg.Wait()
}

func partition(arr []int) int {
    pivot := arr[len(arr)-1]
    i := -1
    for j := 0; j < len(arr)-1; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
    return i + 1
}

func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    pivot := partition(arr)
    QuickSort(arr[:pivot])
    QuickSort(arr[pivot+1:])
}

func main() {
    runtime.GOMAXPROCS(6)

    n := 999
    rand.Seed(69)

    fmt.Printf("Criando um vetor com %d elementos\n", n)
    arr := createRandomArray(n)
	fmt.Print("Vetor criado!\n")

	// fmt.Println("Array antes da ordenação:", arr)

    start := time.Now()

    var wg sync.WaitGroup
    wg.Add(1)
    ParallelQuickSort(arr, &wg)
    wg.Wait()

    elapsed := time.Since(start)

    fmt.Printf("Tempo de execução do Quicksort paralelo: %s\n", elapsed)

	// fmt.Println("Array após a ordenação:", arr)
}
