package forum

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestformatUrl(t *testing.T) {
	fmt.Printf("Testing URL formatting...")
	_, err := formatUrl(1566069)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestShopThread(t *testing.T) {
	fmt.Printf("Testing shop thread...")
	_, err := Retrieve(1566069)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestInvalidThread(t *testing.T) {
	fmt.Printf("Testing invalid thread...")
	_, err := Retrieve(0)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("OK")
}

func TestRandomThreads(t *testing.T) {
	fmt.Printf("Testing 10 random threads...")
	rand.Seed(time.Now().Unix())
	errCh := make(chan error, 10)
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		id := rand.Intn(1600000) // max thread id
		go func() {
			defer wg.Done()
			_, err := Retrieve(id)
			errCh <- err
		}()
	}
	wg.Wait()
	for i := 1; i <= 10; i++ {
		err := <-errCh
		if err != nil {
			t.FailNow()
		}
	}
	fmt.Println("OK")
}
