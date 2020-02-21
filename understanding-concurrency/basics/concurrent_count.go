package basics

import (
	"fmt"
	"time"
	"sync"
)

// infinetily printin index along with same input
func count(input string) {
	for i := 0; true; i++ {
		fmt.Println(i, input)
	}
}

// infinetily printin index along with same input
func count2(input string) {
	halfSecond := time.Millisecond*500
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

// will only print sheep with count
func func1(){
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

	msg := <-c // it is receiving values and assigning it to msg, it is blocking, will receive till it get some data
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
		msg, open := <- c // here we also take the status of the channel, like if it is open ot not ?
		if open == false {
			break;
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

	msg := <- c // waits till sender is there
	fmt.Println(msg) // never get executed
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
	plainDeadlock()
}