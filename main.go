package main

import (
	"fmt"
	"log"
	"os"
	"time"

	modbus "github.com/wz2b/modbus"
)

func main() {

	go func() {
		testModbus()
	}()

	// go func() {
	// 	for {
	// 		testHttp()

	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()

	select {}

}

func testModbus() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Modbus Recovered in f", r)
		}
	}()
	handler := modbus.NewTCPClientHandler("10.214.25.151:502")
	handler.Timeout = 10 * time.Second
	handler.IdleTimeout = 60 * time.Minute
	handler.SlaveId = 247
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)

	// Connect manually so that multiple requests are handled in one connection session
	// err := handler.Connect()
	// if err != nil {
	// 	fmt.Println(time.Now())
	// 	fmt.Println(err)
	// }
	// defer handler.Close()

	client := modbus.NewClient(handler)

	for {
		start := time.Now()
		_, err := client.ReadHoldingRegisters(34, 2)
		if err != nil {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(time.Now())
			fmt.Println(err)
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		}
		elapsed := time.Since(start)

		fmt.Println(time.Now(), "   ", elapsed.Abs().String())

		time.Sleep(time.Second * 5)
	}

}

// func testHttp() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered in f", r)
// 		}
// 	}()

// 	resp, err := http.Get("http://connection-api/connections/api/v1")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// We Read the response body on the line below.
// 	_, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }
