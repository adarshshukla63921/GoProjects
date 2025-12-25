package main

import (
	"fmt"
	"sync"
)

// Mutex locks a resource so that only 1 goroutine can access it at a time.
// Help avoid race conditions, and ensure data integrity.

// Mutex -> Mutual Exclusion {sync.Mutex}

// Race condition -> Multiple goroutines access shared data concurrently, leading to unpredictable results.


type Post struct {
	View int
	mu sync.Mutex
}

// we wrtie unlock in defer, incase something breaks, it will stil unlock when the function exits //
// synax -> defer func() {/* code */}() //

func (p *Post) IncrementView(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // unlock -> other goroutines can access now
		wg.Done()
	}()
	// you could also write this below
	// defer p.mu.Unlock()
	// defer wg.Done()
	p.mu.Lock() // lock -> other goroutines wait till unlock
	p.View += 1
}

func main(){
	var wg sync.WaitGroup

	// Note that mutexes should never be copied as they contain internal state.
	// the below code creates mutex directly into the struct itself
	// myPost := Post{View : 0} // also works, more idomatic

	myPost := Post{View : 0, mu : sync.Mutex{}}

	for i:=1;i<=1000;i++{
		wg.Add(1)
		go myPost.IncrementView(&wg)
	}

	// suspends the main function till all goroutines complete.
	wg.Wait()

	// obersve the final view count
	fmt.Println("Total Views:", myPost.View)
}