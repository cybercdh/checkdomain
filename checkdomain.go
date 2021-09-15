package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	flag.Parse()

	API_KEY, present := os.LookupEnv("MASHAPE_KEY")
	if !present {
		log.Fatalln("API Key not found")
	}

	// read user input
	var domains io.Reader
	domains = os.Stdin

	arg_domain := flag.Arg(0)
	if arg_domain != "" {
		domains = strings.NewReader(arg_domain)
	}

	sc := bufio.NewScanner(domains)

	for sc.Scan() {

		c2 := sc.Text()

		resp, err := http.Get("https://domainr.p.rapidapi.com/v2/status?mashape-key=" + API_KEY + "&domain=" + c2)

		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// only output domains that are not active
		sb := string(body)
		if strings.Contains(sb, "inactive") {
			log.Printf(sb)
		}

	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

}
