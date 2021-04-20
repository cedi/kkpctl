package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
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

	//fmt.Println("BODY: " + string(body))

	r := regexp.MustCompile(`.*\/dex\/auth\/keystone\?req=\w*">`)
	match := r.FindStringSubmatch(string(body))
	fmt.Printf("* MATCH: %#v\n", match)

	requestToken := match[0]
	requestToken = strings.Replace(requestToken, "<form method=\"post\" action=\"", "", -1)
	requestToken = strings.Replace(requestToken, "/dex/auth/keystone?req=", "", -1)
	requestToken = strings.Replace(requestToken, "\">", "", -1)
	requestToken = strings.TrimSpace(requestToken)
	fmt.Printf("* REQUEST_TOKEN: %s\n", requestToken)

	ctx := context.Background()

	// Initialize a provider by specifying dex's issuer URL.
	provider, err := oidc.NewProvider(ctx, kkpURL+"/dex")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     "kubermatic",
		ClientSecret: "ZGFma2Rsc2pmYWRzaGZramhlZmxraHF3ZWtuYXZrbG51aXFybnZpbnJpdW52dm5qbnFyaW92cHFub2lmbnZqa2FybnExNDk1ODM5MTAyLTc4ODkyODc0NS1zZGYK",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oauth2Token, err := oauth2Config.Exchange(ctx, requestToken)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		fmt.Printf("Error: %v", rawIDToken)
		return
	}
	fmt.Printf("ID Token: %s", rawIDToken)
}
