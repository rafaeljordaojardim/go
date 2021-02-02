package main

import (
	"fmt"
	"time"
)

func main() {
	canalDatabase1 := getDatabaseInformation(3)
	canalDatabase2 := getDatabaseInformation(2)
	fmt.Println("Pulling database to get something")

	for i := 0; i < 10; i++ {
		select {
		case canal1 := <-canalDatabase1:
			fmt.Println("DO SOMETHING DATA 1", canal1)
		case canal2 := <-canalDatabase2:
			fmt.Println("DO SOMETHING DATA 2", canal2)
		}
	}

	fmt.Println("PROGRAM FINISHED")

	for i := 0; i < 10; i++ {
		select {
		case canal1 := <-canalDatabase1:
			fmt.Println("DO SOMETHING DATA 1", canal1)
		case canal2 := <-canalDatabase2:
			fmt.Println("DO SOMETHING DATA 2", canal2)
		}
	}
}

func getDatabaseInformation(wait int) chan string {
	canal := make(chan string)
	// basicamente estou retornando um canal direto que fica pegando algo no banco
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(wait))
			canal <- fmt.Sprintf("Database data")
		}
	}()

	return canal
}
