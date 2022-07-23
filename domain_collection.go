package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SearchResponse struct {
	Status      string       `json:"status"`
	FreeDomains []FreeDomain `json:"free_domains"`
}

type FreeDomain struct {
	Currency      string `json:"currency"`
	Domain        string `json:"domain"`
	IsInCart      int    `json:"is_in_cart"`
	PriceCent     string `json:"price_cent"`
	PriceInt      string `json:"price_int"`
	ShowTopDomain string `json:"show_top_domain"`
	Status        string `json:"status"`
	Tld           string `json:"tld"`
	Type          string `json:"type"`
}

type SelectResponse struct {
	Available int `json:"available"`
}

type SetDomainValidTimeType struct {
	Status string `json:"status"`
}

func DomainCollection(domain string) SearchResponse {
	client := new(http.Client)
	body := "domain=" + domain + "&tld="
	req, err := http.NewRequest("POST", "https://my.freenom.com/includes/domains/fn-available.php", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("origin", "https://www.freenom.com")
	//req.Header.Set("user-agent:", "https://www.freenom.com")
	req.Header.Set("Cookie", cookie)
	//header := "content-type:application/x-www-form-urlencoded; charset=UTF-8\ncookie:tz=Asia/Shanghai; JSESSIONID=B614F86592ED44D795323E3B58892A38; _eu=0; _ga=GA1.2.596203605.1657291355; _gid=GA1.2.23931287.1657291355; guid=67da48878e0af4289e85c77853a45ed5; AWSALB=SWZUiOt54lFnas4c2IH0uy+S2lqNRsWkWDWv46jZbIEIA3UcG0IW1nTXwo6RFa3oFwdwe0/Jnvaft6X7smignHUgggV1gqQp+ut6eZxRiVszt3Aro+BzmynN7T+7; AWSALBCORS=SWZUiOt54lFnas4c2IH0uy+S2lqNRsWkWDWv46jZbIEIA3UcG0IW1nTXwo6RFa3oFwdwe0/Jnvaft6X7smignHUgggV1gqQp+ut6eZxRiVszt3Aro+BzmynN7T+7\norigin:https://www.vocabulary.com\nreferer:https://www.vocabulary.com/account/classes/yfNwkgmurRq/manage\nsec-ch-ua:\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"\nsec-ch-ua-mobile:?0\nsec-ch-ua-platform:\"macOS\"\nsec-fetch-dest:empty\nsec-fetch-mode:cors\nsec-fetch-site:same-origin\nuser-agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36\nx-requested-with:XMLHttpRequest\n"
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	content, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	//jsonResp := string(content)
	var searchResp SearchResponse
	err = json.Unmarshal(content, &searchResp)
	if err != nil {
		fmt.Println(err)
	}
	return searchResp
}

func SelectDomain(tid, domain string) SelectResponse {
	client := new(http.Client)
	body := "domain=" + domain + "&tld=" + tid
	req, err := http.NewRequest("POST", "https://my.freenom.com/includes/domains/fn-additional.php", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("origin", "https://www.freenom.com")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
	content, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
	jsonResp := string(content)
	var selectRsp SelectResponse
	err = json.Unmarshal(content, &selectRsp)
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
	log.Println(jsonResp)
	return selectRsp
}

func SetDomainValidTime(domain, period string) SetDomainValidTimeType {
	client := new(http.Client)
	body := "domain=" + domain + "&period=" + period
	req, err := http.NewRequest("POST", "https://my.freenom.com/includes/domains/confdomain-update.php", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("origin", "https://www.freenom.com")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
	content, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
	jsonResp := string(content)
	var rsp SetDomainValidTimeType
	err = json.Unmarshal(content, &rsp)
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
	log.Println(jsonResp)
	return rsp
}
