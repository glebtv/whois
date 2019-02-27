package whois

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	whois_parser "github.com/likexian/whois-parser-go"
	"github.com/weppos/publicsuffix-go/net/publicsuffix"
)

const (
	WHOIS_PORT = "43"
)

type TLD struct {
	Group   string `json:"_group"`
	Type    string `json:"_type"`
	Adapter string `json:"adapter"`
	Host    string `json:"host"`
}

var TLDs map[string]TLD

type Result struct {
	Query       string
	Icann       bool
	TLD         string
	Server      TLD
	Raw         string
	Result      whois_parser.WhoisInfo
	Error       error
	Nameservers []string
	Created     time.Time
	Expires     time.Time
}

func init() {
	err := json.Unmarshal(data, &TLDs)
	if err != nil {
		panic(err)
	}
}

func Whois(domain string) *Result {
	result := Result{Query: domain}
	publicSuffix, icann := publicsuffix.PublicSuffix(domain)
	result.Icann = icann
	result.TLD = publicSuffix
	//if !icann {
	//return nil
	//}
	server, ok := TLDs[publicSuffix]
	if !ok {
		result.Error = errors.New("whois server not found for this TLD")
		return &result
	}
	if server.Host == "" {
		result.Error = errors.New("whois server not available for this TLD")
		return &result
	}
	result.Server = server
	//spew.Dump(server)
	raw, err := Query(domain, server.Host)
	if err != nil {
		result.Error = err
		return &result
	}
	result.Raw = raw
	//log.Println(raw)

	parsed, err := whois_parser.Parse(result.Raw)
	if err != nil {
		result.Error = err
		return &result
	}
	result.Result = parsed

	result.Nameservers = strings.Split(parsed.Registrar.NameServers, ",")

	if parsed.Registrar.CreatedDate != "" {
		t, err := dateparse.ParseAny(parsed.Registrar.CreatedDate)
		if err == nil {
			result.Created = t
		}
	}

	if parsed.Registrar.ExpirationDate != "" {
		t, err := dateparse.ParseAny(parsed.Registrar.ExpirationDate)
		if err == nil {
			result.Expires = t
		}
	}

	return &result
}

func Query(domain string, server string) (result string, err error) {
	conn, e := net.DialTimeout("tcp", net.JoinHostPort(server, WHOIS_PORT), time.Second*30)
	if e != nil {
		err = e
		return
	}

	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	buffer, e := ioutil.ReadAll(conn)
	if e != nil {
		err = e
		return
	}

	result = string(buffer)

	return
}
