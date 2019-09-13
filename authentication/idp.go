package authentication

import (
	"fmt"
	"github.com/hasinthaindrajee/awesomeProject/config"
	"golang.org/x/oauth2"
	"log"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func GetOpenIDProviderConfigs(serverConfig config.Config) oauth2.Config {
	endpoint := oauth2.Endpoint{
		AuthURL:   fmt.Sprintf("%s%s", serverConfig.Idp.Host, serverConfig.Idp.AuthorizeEndpoint),
		TokenURL:  fmt.Sprintf("%s%s", serverConfig.Idp.Host, serverConfig.Idp.TokenEndpoint),
		AuthStyle: oauth2.AuthStyleInHeader}

	config := oauth2.Config{
		Endpoint:     endpoint,
		ClientID:     serverConfig.Idp.ClientId,
		ClientSecret: serverConfig.Idp.ClientSecret,
		RedirectURL:  serverConfig.ServiceProvider.Callback,
		Scopes:       serverConfig.ServiceProvider.Scopes,
	}
	return config
}
