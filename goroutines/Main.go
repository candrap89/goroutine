package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {
	//wg.Add(1) --> one because only have 1 go routines
	wg.Add(2)
	go struct_example()
	// to make delay
	go func() {
		fmt.Println("halo ganteng")
		wg.Done()
	}()
	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
}

var wg = sync.WaitGroup{}

type Doctor struct {
	number      int      //inisializer
	actorName   string   //inisializer
	compnanions []string //inisializer
}

func struct_example() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Candra Ganteng",
		compnanions: []string{
			"farida", "zafran", "sarwiyah",
		},
	}
	fmt.Println(aDoctor.actorName)
	wg.Done()
}
