** try these blocks of code in place of main **

** Using the below code, you will see random results, which may or may not be correct **

* You would need to understand and witness race condition and its effects on actual data to understand why we use mutexes. *

* Using Wait Groups only

        func main(){ 
            mypost := Post{View: 0}

            var wg sync.WaitGroup
            for i:=0 ; i<= 4; i++ {
                wg.Add(1)
                go mypost.IncrementView(&wg)
            }
            wg.Wait()
            fmt.Println("Views:", mypost.View)
        } 


* Using Only Goroutines

        func main(){
            mypost := Post{View: 0}

            go mypost.IncrementView()
            go mypost.IncrementView()
            go mypost.IncrementView()
            go mypost.IncrementView()

            fmt.Println("Total Views:", mypost.View)
        }


