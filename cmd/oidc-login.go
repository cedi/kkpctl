package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/oauth2"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var (
	conf         oauth2.Config
	ctx          context.Context
	kkpServer    string
	kkpCloudName string
)

var oidcLoginCommand = &cobra.Command{
	Use:     "oidc-login [cloudName]",
	Short:   "Uses OIDC to login to your KKP Cloud",
	Example: "kkpctl oidc-login imke",
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"login"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			kkpCloudName = args[0]
		} else {
			kkpCloudName = Config.Context.CloudName
		}

		kkpCloud, err := Config.Cloud.Get(kkpCloudName)
		if err != nil {
			return err
		}

		kkpServer = kkpCloud.URL

		ctx = context.Background()
		conf = oauth2.Config{
			ClientID:     kkpCloud.ClientID,
			ClientSecret: kkpCloud.ClientSecret,
			Scopes:       []string{"openid", "email", "profile"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  fmt.Sprintf("%s/dex/auth", kkpServer),
				TokenURL: fmt.Sprintf("%s/dex/token", kkpServer),
			},
			// my own callback URL
			RedirectURL: "http://localhost:8000",
		}

		// add transport for self-signed certificate to context
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
		sslcli := &http.Client{Transport: tr}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

		fmt.Println(color.CyanString("You will now be taken to your browser for authentication"))
		time.Sleep(1 * time.Second)
		open.Run(url)
		time.Sleep(1 * time.Second)
		fmt.Printf("Authentication URL: %s\n", url)

		http.HandleFunc("/", callbackHandler)

		return http.ListenAndServe(":8000", nil)
	},
}

func init() {
	rootCmd.AddCommand(oidcLoginCommand)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)
	accessToken, err := parseToken(queryParts)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println(color.CyanString("Authentication successful"))

	// show succes page
	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the kkpctl CLI</p>"
	fmt.Fprint(w, msg)

	saveToken(accessToken)
	os.Exit(0)
}

func parseToken(queryParts url.Values) (string, error) {
	codes, ok := queryParts["code"]
	if !ok || len(codes) == 0 {
		return "", fmt.Errorf("failed to find code in request response")
	}

	// Use the authorization code that is pushed to the redirect URL
	code := codes[0]

	// Exchange will do the handshake to retrieve the initial access token.
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// The HTTP Client returned by conf.Client will refresh the token as necessary.
	client := conf.Client(ctx, token)

	resp, err := client.Get(kkpServer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	return token.AccessToken, nil
}

func saveToken(accessToken string) {
	kkpCloud, err := Config.Cloud.Get(kkpCloudName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kkpCloud.Bearer = accessToken
	Config.Cloud.Set(kkpCloudName, kkpCloud)
	Config.Save()
}
