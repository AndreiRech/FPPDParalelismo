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

func ParallelMergeSort(arr []int, wg *sync.WaitGroup) {
    defer wg.Done()
    if len(arr) <= 1 {
        return
    }

    mid := len(arr) / 2

    var subWg sync.WaitGroup
    subWg.Add(2)

    go func() {
        ParallelMergeSort(arr[:mid], &subWg)
    }()

    go func() {
        ParallelMergeSort(arr[mid:], &subWg)
    }()

    subWg.Wait()

    merge(arr, mid)
}

func merge(arr []int, mid int) {
    left := append([]int(nil), arr[:mid]...)
    right := append([]int(nil), arr[mid:]...)

    i, j, k := 0, 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }
    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}

func main() {
    runtime.GOMAXPROCS(16)

    n := 10000000 

    fmt.Printf("Criando um vetor com %d elementos\n", n)
    arr := createRandomArray(n)
	fmt.Print("Vetor criado!\n")

	// fmt.Println("Array antes da ordenação:", arr)

    start := time.Now()

    var wg sync.WaitGroup
    wg.Add(1)
    ParallelMergeSort(arr, &wg)
    wg.Wait()

    elapsed := time.Since(start)
	
    fmt.Printf("Tempo de execução do Merge Sort paralelo: %s\n", elapsed)

	//fmt.Println("Array após a ordenação:", arr)
}
