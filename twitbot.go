package main

import (
	"fmt"
	"github.com/agonopol/twitbot/lib"
	"sync"
)

type Msg struct {
}

func (m *Msg) Execute() {
	fmt.Println("Hello, 世界")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	task := new(Msg)
	go lib.Schedule(task, 1000000000, 1000000000000)()

	wg.Wait()

}
