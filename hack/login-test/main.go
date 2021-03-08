package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {
	var kkpURL string
	flag.StringVar(&kkpURL, "kkpurl", "", "URL to the KKP Installation")

	var kkpUser string
	flag.StringVar(&kkpUser, "user", "", "Username")

	var kkpPassword string
	flag.StringVar(&kkpPassword, "password", "", "Username")

	flag.Parse()

	fmt.Printf("\n## Get request parameter\n\n")
	requestURL := fmt.Sprintf("%s/dex/auth?response_type=id_token&client_id=kubermatic&redirect_uri=%s&scope=openid%%20email%%20profile%%20groups&nonce=7vfLk-QoigcosbZe79qvxPOX9gqLL3CS",
		kkpURL,
		kkpURL,
	)
	fmt.Println("* Request URL: " + requestURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Print("ERROR: " + err.Error())
		return
	}

	req.Header.Set("accept", "text/text")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Printf("* Response: %v\n", res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

	// fmt.Println("BODY: " + string(body))

	r := regexp.MustCompile(`.*\/dex\/auth\/keystone\?req=\w*">`)
	match := r.FindStringSubmatch(string(body))
	fmt.Printf("* MATCH: %#v\n", match)

	requestToken := match[0]
	requestToken = strings.Replace(requestToken, "<form method=\"post\" action=\"", "", -1)
	requestToken = strings.Replace(requestToken, "/dex/auth/keystone?req=", "", -1)
	requestToken = strings.Replace(requestToken, "\">", "", -1)
	requestToken = strings.TrimSpace(requestToken)
	fmt.Printf("* REQUEST_TOKEN: %s\n", requestToken)

	fmt.Printf("\n## Authenticate\n\n")
	authUrl := fmt.Sprintf("%s/dex/auth/keystone?req=%s",
		kkpURL,
		requestToken,
	)
	fmt.Println("* Auth URL: " + authUrl)

	form := url.Values{}
	form.Add("login", kkpUser)
	form.Add("password", kkpPassword)

	req, err = http.NewRequest("POST", authUrl, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Print("ERROR: " + err.Error())
		return
	}

	req.Header.Set("accept", "text/text")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Printf("* Response: %v\n", res)

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

	// fmt.Println("BODY: " + string(body))

	fmt.Printf("\n## Get Bearer Token\n\n")
	approvalUrl := fmt.Sprintf("%s/dex/approval?req=%s",
		kkpURL,
		requestToken,
	)
	fmt.Println("* Approval URL: " + approvalUrl)

	req, err = http.NewRequest("GET", approvalUrl, nil)
	if err != nil {
		fmt.Print("ERROR: " + err.Error())
		return
	}

	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("referer", authUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("cookie", "autoredirect=true")
	req.Header.Set("authority", kkpURL)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Printf("* Response: %v\n", res)

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

	// fmt.Println("BODY: " + string(body))

	bearer := res.Header.Get("location")
	bearer = strings.Replace(bearer, kkpURL+"/projects#access_token=", "", -1)
	bearer = strings.Replace(bearer, "state=&token_type=bearer", "", -1)
	fmt.Println("Bearer: " + bearer)
}
