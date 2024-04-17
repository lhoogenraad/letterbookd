package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

var results = []int{}

func main(){
	t0 := time.Now()
	for i := range 100000 {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTitak exec time: %v", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Printf("db call %v finished\n", i)
	save(i)
	wg.Done()
}

func save(result int){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}
