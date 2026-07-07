package main
import "fmt"

func main(){

	ch := make(chan string)
	
	go func() {
		ch <- "Hello from goroutine!"
	}()
	
	message := <-ch
	fmt.Println(message)
}