package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
	arr = []int{22, 8, 3, 31, 4, 2, 42, 1, 16, 6, 11, 25, 9, 18, 10, 12, 88, 14, 7, 15}
	
)

var result = make([]int, len(arr))

func enumSortMain() {
	sortedArray := make(chan []int)
	chunkSize := len(arr) / 10
	go enumSort (arr, chunkSize,sortedArray)
	r := <- sortedArray
	fmt.Println(r)
}

func enumSort(arr []int, subsetSize int, sortedArray chan []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr) / subsetSize)
	for i := 0; i < len(arr); i += subsetSize {
		
		end := i + subsetSize

		if end > len(arr) {
			end = len(arr)
		}
		go enumSortLogic(arr, arr[i:end], &wg)
		
	}
	wg.Wait()
	sortedArray <- result
}

func enumSortLogic(arr []int, unsortedArray []int, wg *sync.WaitGroup) (finalResult []int){
	
	rank := make([]float64, len(unsortedArray))
	for i := 0; i < len(unsortedArray); i++ {
	var x float64= 1
		for j := 0; j < len(arr); j++ {
			if arr[j] < unsortedArray[i] {
				x +=1
			}
		}
	rank[i] = x;
	}
	mutex.Lock()
	for r, rankValue := range rank {
		result[int(rankValue) - 1] = unsortedArray[r]
	}
	mutex.Unlock()
	wg.Done()
	return result
}

func mergeSort(array []int) {
	fmt.Printf("%v\n",  RunMultiMergesort(array))

}

func RunMultiMergesort(data []int) []int {
	bufferdChannel := make(chan struct{}, 4)
	return MultiMergeSort(data, bufferdChannel)
}

func MultiMergeSort(data []int, bufferdChannel chan struct{}) []int {
	if len(data) < 2 {
		return data
	}

	middleIndex := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var leftdata []int
	var rightdata []int

	select {
	case bufferdChannel <- struct{}{}:
		go func() {
			leftdata = MultiMergeSort(data[:middleIndex], bufferdChannel)
			<-bufferdChannel
			wg.Done()
		}()
	default:
		leftdata = SingleMergeSort(data[:middleIndex])
		wg.Done()
	}

	select {
	case bufferdChannel <- struct{}{}:
		go func() {
			rightdata = MultiMergeSort(data[middleIndex:], bufferdChannel)
			<-bufferdChannel
			wg.Done()
		}()
	default:
		rightdata = SingleMergeSort(data[middleIndex:])
		wg.Done()
	}

	wg.Wait()
	return Merge(leftdata, rightdata)
}

func SingleMergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	middleIndex := len(data) / 2
	return Merge(SingleMergeSort(data[:middleIndex]), SingleMergeSort(data[middleIndex:]))
}

func Merge(leftdata []int, rightdata []int) (result []int) {
	result = make([]int, len(leftdata)+len(rightdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(leftdata):
			result[i] = rightdata[ridx]
			ridx++
		case ridx >= len(rightdata):
			result[i] = leftdata[lidx]
			lidx++
		case leftdata[lidx] < rightdata[ridx]:
			result[i] = leftdata[lidx]
			lidx++
		default:
			result[i] = rightdata[ridx]
			ridx++
		}
	}
	return
}



func main() {
	unsorted_array := []int{22, 8, 3, 31, 4, 2, 42, 1, 16, 6, 11, 25, 9, 18, 10, 12, 88, 14, 7, 15}
	runtime.GOMAXPROCS(1)
	fmt.Println("The number of threads on this machine are = ", runtime.GOMAXPROCS(0))

	start_EnumSort := time.Now()
	enumSortMain()
	fmt.Println("Enum Sort took : ", time.Since(start_EnumSort))

	start_MergeSort := time.Now()
	mergeSort(unsorted_array)
	fmt.Println("Merge Sort took : ", time.Since(start_MergeSort))

	runtime.GOMAXPROCS(8)

	fmt.Println("The number of threads on this machine are = ", runtime.GOMAXPROCS(0))

	start_EnumSort_afterCoreChange := time.Now()
	enumSortMain()
	fmt.Println("After increase of cores Enum Sort took : ", time.Since(start_EnumSort_afterCoreChange))

	start_MergeSort_afterCoreChange := time.Now()
	mergeSort(unsorted_array)
	fmt.Println("After increase of cores Merge Sort took : ", time.Since(start_MergeSort_afterCoreChange))

}
