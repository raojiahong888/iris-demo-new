package main

import (
	"fmt"
	"iris-demo-new/unit_test"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestStart(t *testing.T)  {
	newApp()
	t.Run("BootStart", unit_test.BootStart)
	t.Run("TimeoutContextTest", unit_test.TestTimeoutContext)
	t.Run("CancelContextTest", unit_test.TestCancelContext)
	t.Run("TestCancelContextV2", unit_test.TestCancelContextV2)
	t.Run("TestChannel", unit_test.TestChannel)
	t.Run("TestOrder", unit_test.TestOrder)
	t.Run("TestSubmitOrder", unit_test.TestSubmitOrder)
}

// buffered channel, can first write, then get. or first get, then write.
func TestBufferedChan(t *testing.T)  {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	var wg sync.WaitGroup
	n := 20
	type MapData struct {
		key        string
		value      string
	}
	counter := struct {
		sync.RWMutex
		accept chan MapData
		len int64
		m map[string]string
	}{
		accept: make(chan MapData, n),
		m: make(map[string]string),
	}

	defer close(counter.accept)

	// write channel
	for i:=0; i<n; i++ {
		fmt.Println("write channel i:", i)
		s := strconv.Itoa(i)
		counter.accept <- MapData{
			key:   "jiahong"+s,
			value: "v"+s,
		}
	}

	// get channel, get map
	for i:=0; i<n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("get channel i:", i)
			counter.Lock()
			defer counter.Unlock()
			defer wg.Done()
			for  {
				select {
				case accept := <-counter.accept:
					fmt.Println("recipient only sent channel a:", accept)
					if v,ok := counter.m[accept.key]; ok {
						fmt.Printf("counter %s, v is %s \n", accept.key, v)
					} else {
						counter.m[accept.key] = accept.value
						atomic.AddInt64(&counter.len, 1)
					}
					return
				}
			}
		}(i)
	}



	//for i:=0; i<n; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		fmt.Println("write map i:", i)
	//		counter.accept <- i
	//		//counter.receive = counter.accept
	//		s := strconv.Itoa(i)
	//		counter.m[s] = s
	//		atomic.AddInt64(&counter.len, 1)
	//		return
	//	}(i)
	//}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(counter.len)

	return

}

// non-buffered channel, only can first get, then write.
func TestNonBufferedChan(t *testing.T)  {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	var wg sync.WaitGroup
	n := 20
	type MapData struct {
		key        string
		value      string
	}
	counter := struct {
		sync.RWMutex
		accept chan MapData
		len int64
		m map[string]string
	}{
		accept: make(chan MapData),
		m: make(map[string]string),
	}

	defer close(counter.accept)

	// get map
	for i:=0; i<n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("get map i:", i)
			counter.Lock()
			defer counter.Unlock()
			defer wg.Done()
			for  {
				select {
				case accept := <-counter.accept:
					fmt.Println("recipient only sent channel a:", accept)
					if v,ok := counter.m[accept.key]; ok {
						fmt.Printf("counter %s, v is %s \n", accept.key, v)
					} else {
						counter.m[accept.key] = accept.value
						atomic.AddInt64(&counter.len, 1)
					}
					return
				}
			}
		}(i)
	}

	// write map
	for i:=0; i<n; i++ {
		fmt.Println("write map i:", i)
		s := strconv.Itoa(i)
		counter.accept <- MapData{
			key:   "jiahong"+s,
			value: "v"+s,
		}
	}

	//for i:=0; i<n; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		fmt.Println("write map i:", i)
	//		counter.accept <- i
	//		//counter.receive = counter.accept
	//		s := strconv.Itoa(i)
	//		counter.m[s] = s
	//		atomic.AddInt64(&counter.len, 1)
	//		return
	//	}(i)
	//}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(counter.len)

	return

}

func TestNilChan(t *testing.T)  {
	var done <-chan struct{}
	closeChan := make(chan struct{})
	done = closeChan
	close(closeChan)
	for {
		select {
		case closeData := <-done:
			fmt.Println("receive chan", closeData)
			return
		}
	}
}

func TestMap(t *testing.T)  {
	m := make(map[string]string)
	if _,ok := m["name"]; !ok {
		fmt.Println("this map key didnt exist")
	}
}
