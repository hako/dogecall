/*
	dogecall.go
	porting my small dogecall.py program to Go.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/hako/dogecall/Godeps/_workspace/src/github.com/docopt/docopt-go"
)

var usage string = `dogecall - ENCOUNTER A DOGE THROUGH A PHONE CALL.

Usage:
 dogecall [-n <number>]
 dogecall [-h | --help] [-v | --version]

Options:
  -n [PHONE NUMBER]   such fone numer (with area codez) eg. +44 = 44
  -h, --help          show this help message and exit.
  -v, --version       the versions lol.

Example: dogecall -n [PHONE NUMBER]`

type DogeCallRC struct {
	Accountsid  string `json:"account_sid"`
	TwAuthtoken string `json:"tw_authtoken"`
	TwNumber    string `json:"tw_number"`
	URL         string `json:"url"`
}

var VERSION = "0.5"
var TWILIO_API_BASE = "https://api.twilio.com/2010-04-01"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	args, _ := docopt.Parse(usage, nil, true, VERSION, false)

	// Load .dogecallrc
	dcConfig, err := loadDogeCallRC()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args["-n"] != nil {

		phoneNumber := args["-n"].(string)
		twilioNumber := dcConfig.TwNumber

		// Check both phone numbers before calling.
		check1 := CheckNumber(phoneNumber)
		if check1 != true {
			fmt.Printf("wow, such number \"%s\" not valid, lol.", phoneNumber)
			os.Exit(1)
		}

		check2 := CheckNumber(twilioNumber)
		if check2 != true {
			fmt.Printf("wow, such number in config \"%s\" is not valid, lol.", twilioNumber)
			os.Exit(1)
		}
		Call(twilioNumber, phoneNumber)
	}
}

//	Checks the phone number.
func CheckNumber(phoneNumber string) bool {
	phoneRegex := `[(]?\d{3}[)]?\s?-?\s?\d{3}\s?-?\s?\d{4}`
	result, _ := regexp.MatchString(phoneRegex, phoneNumber)
	return result
}

//	Load .dogecallrc.
func loadDogeCallRC() (DogeCallRC, error) {
	var data DogeCallRC

	usr, _ := user.Current()
	rc, err := ioutil.ReadFile(usr.HomeDir + "/" + ".dogecallrc")
	if err != nil {
		fmt.Println("creating dogecallrc...")
		createDogeCallRC(usr.HomeDir)
		return data, errors.New("please configure .dogecallrc 2 use dogecall pls.")
	} else {
		err := json.Unmarshal(rc, &data)
		if err != nil {
			return data, errors.New("Error decoding from JSON.")
		}
	}

	return data, nil
}

// Create .dogecallrc if it does not exist already.
func createDogeCallRC(dir string) {
	var data DogeCallRC

	data.Accountsid = ""
	data.TwNumber = ""
	data.URL = "http://dc.hakobaito.co.uk/doge" // Default URL, feel free to change.
	data.TwAuthtoken = ""

	rcdata, _ := json.MarshalIndent(data, "", "  ")
	ioutil.WriteFile(dir+"/"+".dogecallrc", rcdata, 0644)
}

// Make a call to the person.
func Call(twilioNumber string, recievingNumber string) {
	dcConfig, _ := loadDogeCallRC()

	accountSid := dcConfig.Accountsid
	authToken := dcConfig.TwAuthtoken
	callURL := dcConfig.URL
	TWILIO_CALL_URL := TWILIO_API_BASE + "/Accounts/" + accountSid + "/Calls.json"

	// Prepare POST values.
	// Remember, the 'To' number must be a 'verified number' when using a twilio trial account.
	postVals := url.Values{}
	postVals.Set("To", recievingNumber)
	postVals.Set("From", twilioNumber)
	postVals.Set("Record", "false")
	postVals.Set("Url", callURL)

	data := *strings.NewReader(postVals.Encode())

	fmt.Println("pls wait...")

	// Create a HTTP client, prepare and send the request.
	client := &http.Client{}
	req, _ := http.NewRequest("POST", TWILIO_CALL_URL, &data)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(req)
	if res.StatusCode != 201 {
		fmt.Printf("such number \"%s\" so wrong, check again pls. sad doge.", recievingNumber)
	} else {
		fmt.Printf("wow. a wild shibe in 3 secs %s, pick up pls.", recievingNumber)
	}
}
