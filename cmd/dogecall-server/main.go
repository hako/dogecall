/*
	dogecall-server.go
	porting my small dogecall.py program to Go.
*/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/hako/dogecall/Godeps/_workspace/src/github.com/docopt/docopt-go"
	"github.com/hako/dogecall/Godeps/_workspace/src/github.com/joho/godotenv"
)

var usage string = `dogecall-server - SERVER FOR DOGECALL.

Usage:
 dogecall [-s]
 dogecall [-h | --help] [-v | --version]

Options:
  -s                  wow lanch server 2 serve u.
  -h, --help          show this help message and exit.
  -v, --version       the versions lol.
`
var VERSION = "0.5"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	args, _ := docopt.Parse(usage, nil, true, VERSION, false)

	if os.Getenv("GO_ENV") != "production" {

		err := godotenv.Load()
		if err != nil {
			fmt.Println("stap, no .env file. wont start.")
			os.Exit(1)
		}
	}

	if args["-s"] != false {
		if os.Getenv("PORT") == "" {
			fmt.Println("stap, doge port not set in env.")
			os.Exit(1)
		} else {
			port := os.Getenv("PORT")
			ServeTwiML(port)
		}
	}
}

//	Serves TwiML on the server.
func ServeTwiML(port string) {
	//	Response TwiML page based on the digits on the phone.
	ResponseHandler := func(rw http.ResponseWriter, req *http.Request) {
		var twiml string

		digit := req.FormValue("Digits")

		if digit == "1" {
			twiml = `
		<Response>
			<Say>You've chosen: "Pet, Doge.". The doge, is pleased.</Say>
		</Response>
		`
		} else if digit == "2" {
			twiml = `
		<Response>
			<Say>You've chosen: "Snuggle, Doge.". The doge, is pleased.</Say>
		</Response>
		`
		} else if digit == "3" {
			twiml = `
		<Response>
		  		<Say>You've chosen: "Feed, Doge.". The doge, is pleased.</Say>
		</Response>
		`
		} else if digit == "0" {
			twiml = `
		<Response>
			<Gather action="/response"  method="GET" numDigits="1">
		  		<Say>To Pet the Doge Press 1. To Snuggle the Doge Press 2. To Feed the Doge Press 3. To Repeat the options Press 0.</Say>
		  		</Gather>
		    <Redirect/>
		</Response>
		`
		} else {
			twiml = `
		<Response>
		  		<Say>Invalid option.</Say>
		    <Redirect>/doge</Redirect>
		</Response>
		`
		}

		rw.Header().Set("Content-Type", "application/xml")
		rw.Write([]byte(twiml))
	}
	//	DogeMenuHandler TwiML page.
	DogeMenuHandler := func(rw http.ResponseWriter, req *http.Request) {
		var twiml string
		twiml = `
		<Response>
		    <Gather action="/response"  method="GET" numDigits="1">
		        <Say>Hello. You have encountered, a doge. Choose from the following options: To Pet the Doge Press 1. To Snuggle the Doge Press 2. To Feed the Doge Press 3. To Repeat the options Press 0.</Say>
		    </Gather>
		    <Redirect/>
		</Response>
		`
		rw.Header().Set("Content-Type", "application/xml")
		rw.Write([]byte(twiml))
	}

	//	Index page for dogecall TwiML. Nothing fancy yet.
	Index := func(rw http.ResponseWriter, req *http.Request) {
		t := time.Now()
		randomDoge := []string{"wow many servers", "such server", "http lol", "wow", "the http", "dogecall"}
		rand.Seed(t.UnixNano())
		rnd := randomDoge[rand.Intn(len(randomDoge))]
		rw.Write([]byte(rnd))
	}

	//	Setup TWiML http handlers.
	http.HandleFunc("/", Index)
	http.HandleFunc("/doge", DogeMenuHandler)
	http.HandleFunc("/response", ResponseHandler)

	//	Listen and serve.
	fmt.Printf("serving on http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
