package main

import (
	"fmt"
	"time"
)

//func main() {
//	message := make(chan int, 10)
//
//	for i := 0; i < 10; i++ {
//		message <- i
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//
//	go func(ctx context.Context) {
//		ticker := time.NewTicker(1 * time.Second)
//		for _ = range ticker.C {
//			select {
//			case <-ctx.Done():
//				fmt.Println("child process interrupt...")
//				return
//			default:
//				fmt.Printf("send message:%+v \n", <-message)
//
//			}
//		}
//	}(ctx)
//
//	defer close(message)
//	defer cancel()
//
//	select {
//	case <-ctx.Done():
//		fmt.Println("main get done")
//		time.Sleep(1 * time.Second)
//		fmt.Println("main process exit!")
//	}
//}

func main() {
	message := make(chan int, 10)

	for i := 0; i < 10; i++ {
		message <- i
	}

	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message:%+v \n", <-message)

			}
		}
	}()

	defer close(message)
	close(done)

	select {
	case <-done:
		fmt.Println(<-done)
		fmt.Println("main get done")
		time.Sleep(5 * time.Second)
		fmt.Println("main process exit!")
	}
}
