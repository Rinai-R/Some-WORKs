package main

import "sync"

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	//done就是add传入-1
	wg.Done()

	wg.Wait()
}
