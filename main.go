package main



import (

    "fmt"
    "sync"
)

//test 123

//test 123



//test 123
//test 123
//test 123
//test 123
//test 123



var ch1 chan int = make(chan int,1)  //声明并 初始化channel 变量





var ch2 chan int = make(chan int,1)  //声明并初始化channel变量


var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main() {

	wg := &sync.WaitGroup{}
	
	wg.Add(1)
	
	go func () {
	    
	defer wg.Done()
		getChan(0)<-1
		
		getChan(1)<-2
	}()
	wg.Wait()
	select {

	case <-getChan(0):

        fmt.Println("1th case is selected.")

	case <-getChan(1):

        fmt.Println("2th case is selected.")

    default:

    fmt.Println("default!.")

    }

}

func getNumber(i int) int {
    
    
    

    fmt.Printf("numbers[%d]\n", i)

    return numbers[i]

}

func getChan(i int) chan int {

    fmt.Printf("chs[%d]\n", i)
    

    return chs[i]

}
