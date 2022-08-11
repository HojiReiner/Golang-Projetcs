package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Println("\nInsert Domain:") ; scanner.Scan(); fmt.Println("\nInsert Domain:"){
		CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %+v \n", err)
	}
}

func CheckDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//Get MX Records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	} else {
		hasMX = len(mxRecords) > 0
	}
	
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println("\nDomain ->", domain)
	fmt.Println("Has MX? ->", hasMX)
	fmt.Println("Has SPF? ->", hasSPF)
	if hasSPF {
		fmt.Println("SPF Record ->", spfRecord)
	}
	fmt.Println("Has DMARC? ->", hasDMARC)
	if hasDMARC {
		fmt.Println("DMARC Record ->", dmarcRecord)
	}
}
