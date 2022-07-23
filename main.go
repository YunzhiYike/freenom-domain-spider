package main

import (
	"log"
	"os"
)

var cookie string

func main() {
	ck, _ := os.ReadFile("cookie.ini")
	cookie = string(ck)
	domainList := []string{"cf888", "dz233", "mk22", "dw111"}
	for _, domain := range domainList {
		batchAddDomain(domain)
	}
}

func batchAddDomain(name string) {
	domains := DomainCollection(name)
	for _, domain := range domains.FreeDomains {
		log.Println(domain.Tld, domain.Domain, domain.Type)
		if domain.Type == "FREE" {
			selectRsp := SelectDomain(domain.Tld, domain.Domain)
			dm := domain.Domain + domain.Tld
			if selectRsp.Available != 1 {
				log.Println(dm + " 域名选择失败 ❌ ")
				continue
			}
			res := SetDomainValidTime(dm, "12M")
			if res.Status == "OK" {
				log.Println(dm + " 加入购物车成功 ✅ ")
				continue
			}
			log.Println(dm + " 加入购物车失败 ❌ ")
		}
	}
}
