package config

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

type TLDs struct {
	Domains []string
}

var cacheTld *cache.Cache

func init() {
	cacheTld = cache.New(7*24*time.Hour, 24*time.Hour)
	GetTlds()
}

func CheckTld(str string) bool {
	val := getTldsValue()
	for _, tld := range val.Domains {
		if strings.Contains(str, "."+strings.ToLower(tld)) {
			log.Print(str, ", \".\"+", strings.ToLower(tld))
			return false
		}
	}
	return true
}

func getTldsValue() TLDs {
	val, _ := cacheTld.Get("tlds")
	return val.(TLDs)
}

func GetTlds() {
	resp, err := http.Get("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		log.Fatal("Fatal get tlds: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Fatal read response's body: ", err)
		return
	}
	arrTlds := strings.Split(string(body), "\n")
	i := len(arrTlds) - 2
	arrTlds = arrTlds[:0+copy(arrTlds[0:], arrTlds[0+1:])]
	arrTlds = arrTlds[:i+copy(arrTlds[i:], arrTlds[i+1:])]

	struTld := TLDs{
		Domains: arrTlds,
	}
	cacheTld.Set("tlds", struTld, cache.DefaultExpiration)
}
