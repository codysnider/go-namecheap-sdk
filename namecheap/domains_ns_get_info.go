package namecheap

import (

	"encoding/xml"
	"fmt"

)

type DomainsNSGetInfoResponse struct {
	XMLName xml.Name `xml:"ApiResponse"`
	Errors  []struct {
		Message string `xml:",chardata"`
		Number  string `xml:"Number,attr"`
	} `xml:"Errors>Error"`
	CommandResponse *DomainsNSGetInfoCommandResponse `xml:"CommandResponse"`
}

type DomainsNSGetInfoCommandResponse struct {
	DomainDNSGetHostsResult *DomainNSInfoResult `xml:"DomainNSInfoResult"`
}

type DomainNSInfoResult struct {
	Domain     *string `xml:"Domain,attr"`
	Nameserver *string `xml:"Nameserver,attr"`
	IP         *string `xml:"IP,attr"`
}

// GetInfo retrieves NS settings for the requested domain.
//
// Namecheap doc: https://www.namecheap.com/support/api/methods/domains-ns/getinfo/
func (dds *DomainsNSService) GetInfo(nameserver string) (*DomainsNSGetInfoCommandResponse, error) {
	var response DomainsNSGetInfoResponse

	parsedDomain, err := ParseDomain(nameserver)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"Command":    "namecheap.domains.ns.getInfo",
		"Nameserver": nameserver,
		"SLD":        parsedDomain.SLD,
		"TLD":        parsedDomain.TLD,
	}

	_, err = dds.client.DoXML(params, &response)
	if err != nil {
		return nil, err
	}
	if len(response.Errors) > 0 {
		apiErr := response.Errors[0]
		return nil, fmt.Errorf("%s (%s)", apiErr.Message, apiErr.Number)
	}
	
	return response.CommandResponse, nil
}
