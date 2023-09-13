package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmbracelet/log"
)

func main() {
	var domain string
	prompt := &survey.Input{
		Message: "enter the domain : ",
	}
	survey.AskOne(prompt, &domain, survey.WithValidator(survey.Required), survey.WithValidator(func(val interface{}) error {
		str := val.(string)
		domainRegex := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
		if !domainRegex.MatchString(str) {
			return fmt.Errorf("please enter a valid domain")
		}
		return nil
	}))

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRcs, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}
	hasMX = len(mxRcs) > 0

	txtRcs, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range txtRcs {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRcs, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range dmarcRcs {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("hasMX: %v\n", hasMX)
	fmt.Printf("hasSPF: %v\n", hasSPF)
	fmt.Printf("spfRecord: %v\n", spfRecord)
	fmt.Printf("hasDMARC: %v\n", hasDMARC)
	fmt.Printf("dmarcRecord: %v\n", dmarcRecord)
}
