package main

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
	"time"
)

type Report struct {
	Date        string
	Domain      string
	Subdomain   string
	Fingerprint string
	Ports       string
	Ipspace string
	Probing string
	Vuln    string
}

var home = "CHANGE FOR USER DIRECTORY"
var domain = os.Args[1]
var path = home + ".osmedeus/workspaces/" + domain + "/"
var fingerprintFile = path + "fingerprint/beautify-" + domain + "-http.txt"
var subdomainFile = path + "subdomain/final-" + domain + ".txt"

var ipspaceFile = path + "ipspace/" + domain + "-ip.txt"
var probingFile = path + "probing/diffhttp-" + domain + ".txt"
var vulnFile = path + "vuln/active/vuln-summary.txt"
var portsFile = path + "portscan/open-ports.txt"

//obtener fecha

var date = time.Now().Format("2006-01-02")

//read text file
func readFile(file string) string {
	var text string
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	return text
}

//use markdown template

func toTemplate(report Report) {
	t, err := template.ParseFiles("template.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f := t.Execute(os.Stdout, report)
	if f != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r, err := os.Create("report_" + domain + ".md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f = t.Execute(r, report)

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./report <domain>")
		os.Exit(1)
	}
	//if file exists
	fingerprint := ""
	if _, err := os.Stat(fingerprintFile); err == nil {
		//read file
		fingerprint = readFile(fingerprintFile)
	} else {
		fingerprint = "No se encontraron datos"
	}

	subdomain := ""
	if _, err := os.Stat(subdomainFile); err == nil {
		//read file
		subdomain = readFile(subdomainFile)
	} else {
		subdomain = "No se encontraron datos"
	}
	ipspace := ""
	if _, err := os.Stat(ipspaceFile); err == nil {
		//read file
		ipspace = readFile(ipspaceFile)
	} else {
		ipspace = "No se encontraron datos"
	}
	probing := ""
	if _, err := os.Stat(probingFile); err == nil {
		//read file
		probing = readFile(probingFile)
	} else {
		probing = "No se encontraron datos"
	}
	vuln := ""
	if _, err := os.Stat(vulnFile); err == nil {
		//read file
		vuln = readFile(vulnFile)
	} else {
		vuln = "No se encontraron datos"
	}
	ports := ""
	if _, err := os.Stat(portsFile); err == nil {
		//read file
		ports = readFile(portsFile)
	} else {
		ports = "No se encontraron datos"
	}

	report := Report{date, domain, subdomain, fingerprint, ports, ipspace, probing, vuln}
	toTemplate(report)

}
