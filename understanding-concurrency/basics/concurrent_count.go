package basics

import (
	"fmt"
	"sync"
	"time"
)

// infinetily printin index along with same input
func count(input string) {
	for i := 0; true; i++ {
		fmt.Println(i, input)
	}
}

// infinetily printin index along with same input
func count2(input string) {
	halfSecond := time.Millisecond * 500
	for i := 0; i <= 5; i++ {
		fmt.Println(i, input)
		time.Sleep(halfSecond)
	}
}

// using string type channel 'c' to communicate (here it is sending)
func count3(input string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- input // blocking, waiting till reciever is ready to receive.
		fmt.Println(i, input)
	}
}

// using string type channel 'c' to communicate (here it is sending)
// here, it waits for half second too.
func count4(input string, c chan string) {
	halfSecond := time.Millisecond * 500
	for i := 0; i <= 5; i++ {
		c <- input // blocking, waiting till reciever is ready to receive.
		fmt.Println(i, input)
		time.Sleep(halfSecond)
	}
}

// using string type channel 'c' to communicate (here it is sending)
// here, it waits for half second and closes chanel once finished.
func count5(input string, c chan string) {
	halfSecond := time.Millisecond * 500
	for i := 0; i <= 5; i++ {
		c <- input // blocking, waiting till reciever is ready to receive.
		fmt.Println(i, input)
		time.Sleep(halfSecond)
	}

	close(c)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// will only print sheep with count
func func1() {
	count("sheep")
}

// will only print sheep with count, the count("fist") never get executed
func func2() {
	count("sheep")
	go count("fish")
}

// will not print anything as by the time it start execution of count func, the func3() completes.
func func3() {
	go count("sheep")
	go count("fish")
}

// this is blocking, will go-routine completes, here it will run infinitely
func func4() {
	var wg sync.WaitGroup
	wg.Add(1) // imlies , I am waiting for '1' go-routine to complete

	go func() {
		count("sheep")
		wg.Done()
	}() // anonymous function, coz, the wait-group related changes are not responsibility of go-routine
	wg.Wait()
}

// this is blocking, will go-routine completes, here it will run 5 times while printing some info.
func func5() {
	var wg sync.WaitGroup
	wg.Add(1) // imlies , I am waiting for '1' go-routine to complete

	go func() {
		count2("sheep")
		wg.Done()
	}() // anonymous function, coz, the wait-group related changes are not responsibility of go-routine
	wg.Wait()
}

//using channel to send/recieve some info, channel work like a pipe,
//Also, this will get hte info in msg and then terminates,
//May process the some next steps like on receiver OR sender end
// blocking nature of channels helps in synchronising the code.
func func6() {
	c := make(chan string)
	go count3("sheep", c)

	msg := <-c       // it is receiving values and assigning it to msg, it is blocking, will receive till it get some data
	fmt.Println(msg) // it may get executed after some statement from the go-routine, btw
}

//channels help in synchroinization, This will give a error as the count3 would be done with sending data, but the receiving end still waits, creating a dedlock
func func7() {
	c := make(chan string)
	go count4("sheep", c)

	for {
		msg := <-c       // it is receiving values and assigning it to msg, it is blocking, will receive till it get some data
		fmt.Println(msg) // it may get executed after some statement from the go-routine, btw
	}
}

// resolving the issue of deadlock, by closing channel at sender side.
// never close channel at receiver, coz you may not know if sender like to send some data
// this func will still go-on infinitely though
func func8() {
	c := make(chan string)
	go count5("sheep", c)

	for {
		msg := <-c       // it is receiving values and assigning it to msg, it is blocking, will receive till it get some data
		fmt.Println(msg) // it may get executed after some statement from the go-routine, btw
	}
}

// resolving the issue of deadlock, by closing channel at sender side.
// never close channel at receiver, coz you may not know if sender like to send some data
// this func will still go-on stop once channel closes though
func func9() {
	c := make(chan string)
	go count5("sheep", c)

	for {
		msg, open := <-c // here we also take the status of the channel, like if it is open ot not ?
		if open == false {
			break
		}
		fmt.Println(msg) // it may get executed after some statement from the go-routine, btw
	}
}

// resolving the issue of deadlock, by closing channel at sender side.
// never close channel at receiver, coz you may not know if sender like to send some data
// this func will still go-on stop once channel closes though
func func9_2() {
	c := make(chan string)
	go count5("sheep", c)

	for msg := range c { // syntactic sugur to receive and use the msg value, by using range over the channel
		fmt.Println(msg) // it may get executed after some statement from the go-routine, btw
	}
}

// naively, we may think, this will work, but it will not, rather it will cause dead-lock as channels only communicate across go-routines
func plainDeadlock() {
	c := make(chan string)

	c <- "sheep" // send waits till receiver is there

	msg := <-c       // waits till sender is there
	fmt.Println(msg) // never get executed
}

func func10() {
	c := make(chan string, 2) // buffered channel of capacity of 2, can be non-blockong till buffer is filled.
	c <- "hello1"
	c <- "hello2" // can put till 2 things before we have to wait for reciever to exist some-place.

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

func func11() {
	c := make(chan string, 2) // buffered channel of capacity of 2, can be non-blockong till buffer is filled.
	c <- "hello1"
	c <- "hello2" // can put till 2 things before we have to wait for eciever to exist some-place.
	c <- "hello3" // wil block, thus create a dead-lock, as channel is gonna be full.

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

// even though the first message arrive early and it can now receive more messages from first channel, it waits till 2nd channel sends
// 2nd go-routine execution kinda slow down the first go-routine`s usage/processing.
func plainSynchronization() {
	c1 := make(chan string)
	c2 := make(chan string)
	halfSecond := time.Millisecond * 500
	twoSeconds := time.Second * 2
	go func() {
		for {
			c1 <- "Every 500Ms"
			time.Sleep(halfSecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2Sec"
			time.Sleep(twoSeconds)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}

// using Select statement to get away from the slowing down part of first go-routine
func func13() {
	c1 := make(chan string)
	c2 := make(chan string)
	halfSecond := time.Millisecond * 500
	twoSeconds := time.Second * 2
	go func() {
		for {
			c1 <- "Every 500Ms"
			time.Sleep(halfSecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2Sec"
			time.Sleep(twoSeconds)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

//get the job from 'jobs' channel and return it in 'result' channel
func worker(jobs chan int, results chan int) {
	for n := range jobs {
		results <- fib(n)
	}
}

//single worker based usage
func func14() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results) // making it a worker/processor

	for i := 0; i < 100; i++ {
		jobs <- i % 17 // sending 100 integers on jobs channel
	}

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// multiple worker based usage
func func15() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results) // making it a worker/processor
	go worker(jobs, results) // making it a worker/processor
	go worker(jobs, results) // making it a worker/processor
	go worker(jobs, results) // making it a worker/processor

	for i := 0; i < 100; i++ {
		jobs <- i // sending 100 integers on jobs channel
	}

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

//playing with closing channels, Closing always gives a message!!!
func func16() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c) // it prints zero value of the type-of channel
	// prints 0

	//BTW, closing gives 2 values : Zero-value of type, false (telling theere will be no more chanel)
	c2 := make(chan string)
	close(c2)
	zeroValue, isMore := <-c2 // zero value for a string is ""
	fmt.Println("correct Zero value : ", zeroValue == "")
	fmt.Println("channel is closed, thus this is `false`?, isMore : ", isMore)

}

//playing with closing channels, Closing always gives a message!!!,
//what if I try to get the closing message multiple times ? : it returns zero again!!!
func func17() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c) // it prints zero value of the type-of channel
	fmt.Println(<-c)
}

func func18() {

}

func Start() {
	//func1()
	//func2()
	//func3()

	//With Waitgroup
	//func4()
	//func5()

	//with channels
	//func6()
	//func7()
	//func8()
	//func9()
	//func9_2()
	//plainDeadlock()

	//using buffered channels
	//func10()
	//func11()

	//synchronization
	//plainSynchronization()
	//func13()

	// worker pools
	//func14()
	//func15()

	//playing with closing channels
	//func16()
	func17()
}
