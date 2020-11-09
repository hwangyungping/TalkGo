package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup // number of working goroutines
var results = make(chan int64, 100)
var routineNum int = 10
var runningNum int = 100
var percent float64 = 0.95

func main() {
	fmt.Printf("Hello, Please input Test URL, default baidu.com \n")
	input, err := ReadFrom(os.Stdin, 20)
	if err != nil {
		fmt.Printf("Read input fail.")
		return
	}

	url := string(input)
	if len(url) < 5 {
		url = "http://baidu.com"
	} else {
		url = "http://" + url
	}

	fmt.Print("Input RoutineNum(10) and RunningMax(100), using Space.")
	fmt.Scanln(&routineNum, &runningNum)
	fmt.Printf("Input RoutineNum: %d RunningMax: %d\n", routineNum, runningNum)

	//利用channel自动阻塞的性质来控制当前运行的goroutine的总数量
	ch := make(chan int, routineNum)

	for i := 0; i < runningNum; i++ {
		//每当创建goroutine的时候就向channel中放入一个数据，如果里面已经有10个数据了，就会
		//阻塞，由此我们将同时运行的goroutine的总数控制在<=10个的范围内
		ch <- 1
		wg.Add(1)
		go GetUrl(url, ch)
		//fmt.Printf("index: %d,goroutine Num: %d \n", i, runtime.NumGoroutine())
	}

	// 等待所有goroutine结束
	wg.Wait()
	close(results)

	var sum int64 = 0
	index := 1

	var sortValues []int
	for val := range results {
		sum += val
		sortValues = append(sortValues, (int)(val))
		//fmt.Printf("Index %d, %d ms\n", index, val)
		index++
	}
	num := len(sortValues)

	fmt.Printf("Total %d sum = %v ms, while mean=%.3v ms.\n", num, sum, float64(sum)/float64(num))

	sort.Ints(sortValues)
	sortsum := 0
	sortnum := (int)(float64(num) * percent)
	for i, val := range sortValues {
		if i < sortnum {
			sortsum += val
		}
		//fmt.Printf("Sort Index %d, %d ms\n", i, val)
	}

	fmt.Printf("Sort Total %d sortsum = %v ms, while mean=%.3v ms.\n", sortnum, sortsum, float64(sortsum)/float64(sortnum))

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 1 {
		return p[:n-1], nil
	}
	return p, err
}

func GetUrl(url string, ch chan int) (int64, error) {
	defer wg.Done()

	start := time.Now()
	_, err := http.Get(url)
	elapsed := time.Since(start)
	<-ch
	results <- elapsed.Milliseconds()
	if err != nil {
		//fmt.Printf("GET %s fail %s : %d ms .\n", url, err.Error(), elapsed.Milliseconds())
		return elapsed.Milliseconds(), err
	} else {
		//fmt.Printf("GET %s Success : %d ms \n", url, elapsed.Milliseconds())
		return elapsed.Milliseconds(), nil
	}

}
