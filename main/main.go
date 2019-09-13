package main

import (
	"crypto/tls"
	"fmt"
	"github.com/hasinthaindrajee/awesomeProject/authentication"
	"github.com/hasinthaindrajee/awesomeProject/config"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting application..")
	serverConfigs := config.ReadConfig()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	oauthConfig := authentication.GetOpenIDProviderConfigs(serverConfigs)
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received..")
		var code = r.URL.Query().Get("code")
		fmt.Fprintf(w, "authorization code, %q", code)
		token, error := oauthConfig.Exchange(oauth2.NoContext, code)
		if error != nil {
			fmt.Println("error occurred", error)
		} else {
			fmt.Println(token.Extra("id_token"))
		}
		os.Exit(0)
	})

	fmt.Println(oauthConfig.AuthCodeURL(""))
	authentication.OpenBrowser(oauthConfig.AuthCodeURL(""))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
