/*
	dogecall-server.go
	porting my small dogecall.py program to Go.
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/hako/dogecall/Godeps/_workspace/src/github.com/docopt/docopt-go"
	"github.com/hako/dogecall/Godeps/_workspace/src/github.com/joho/godotenv"
)

var logo = `

                                                                        
             .N.                                           lW. .W:      
             :M.                                           oM. .M:      
        ,kxxkKN  .xOO0l   :kOOOo  'kkxkk.  .kOxkd  ;kkxOx  dW  .M;      
       ,N.   kK .Wo  .Wc 0O.  cW.'N,'lkk. .Nc     oX.  dK  dN  .M;      
       cX    xK ;M,   Nc.Mc  .KK lWxc.  . lX      Ok   l0  kX  'M,      
        dKxxOKN  oKxxKo  ;0OOx0k  d0xdxkx  k0dokx 'KkldOW; k0  ,M.      
           .       ..        .Nl     .       ..     ..                  
                         dxxx0d                                         
                                                                        

                      S   E   R   V   E   R
`

var usage = `dogecall-server - SERVER FOR DOGECALL.

Usage:
 dogecall [-s] [-p <port>]
 dogecall [-h | --help] [-v | --version]

Options:
  -p                  port.
  -s                  wow lanch server 2 serve u.
  -h, --help          show this help message and exit.
  -v, --version       the versions lol.
`
var version = "0.5"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	args, _ := docopt.Parse(usage, nil, true, version, false)

	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("warning: no .env file. continuing in development environment...")
		}
	}

	if args["-s"] != false || args["-p"] != false {
		var port string

		if args["<port>"] != nil {
			port = args["<port>"].(string)
		} else if os.Getenv("PORT") != "" {
			port = os.Getenv("PORT")
		} else {
			fmt.Println("stap, doge port not set in env or argument.")
			os.Exit(1)
		}

		serveTwiML(port)
	}
}

//	Serves TwiML on the server.
func serveTwiML(port string) {
	logger := log.New(os.Stdout, "", log.Ldate)

	//	Response TwiML page based on the digits on the phone.
	ResponseHandler := func(rw http.ResponseWriter, req *http.Request) {
		var twiml string

		logger.Println(req.Method, req.Host, req.RequestURI)

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
		logger.Println(req.Method, req.Host, req.RequestURI)
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
		logger.Println(req.Method, req.Host, req.RequestURI)
		t := time.Now()
		randomDoge := []string{"many calls lol", "doge lol", "very gopher", "many requests", "just wow", "wow many servers", "such rand", "such server", "http lol", "wow", "the http", "dogecall"}
		rand.Seed(t.UnixNano())
		rnd := randomDoge[rand.Intn(len(randomDoge))]
		rw.Write([]byte(rnd))
	}

	//	Setup TWiML http handlers.
	http.HandleFunc("/", Index)
	http.HandleFunc("/doge", DogeMenuHandler)
	http.HandleFunc("/response", ResponseHandler)

	//	Listen and serve.
	fmt.Println(logo)
	fmt.Printf("serving on http://localhost:%s lol.\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
