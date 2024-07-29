package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Mañana")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Mañana")
	case t.Hour() < 16:
		fmt.Println("Tarde")
	default:
		fmt.Println("Noche")

	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> windows")
	case "linux":
		fmt.Println("Go run -> linux")
	case "darwin":
		fmt.Println("Go run -> Mac OS")
	default:
		fmt.Println("otro OS")
	}

}
