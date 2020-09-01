package setting

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type app struct {
	Title    string
	MinDay   int
}
var AppSetting = &app{}

type server struct {
	RunMode      string
	Port         string
}
var ServerSetting = &server{}

type trello struct {
	AppKey     string
	Token      string
	BoardID    string
	ListNumber int
}
var TrelloSetting = &trello{}

var GoogleTokenJson = &oauth2.Token{}

type googleDrive struct {
	FolderId	string
}
var GoogleDriveSetting = &googleDrive{}

type googleCredentialsJsonParent struct {
	Installed googleCredentialsJsonChild `json:"installed"`
}
var googleCredentialsJsonParentSetting = &googleCredentialsJsonParent{}

type googleCredentialsJsonChild struct {
	ClientID				string `json:"client_id"`
	ClientSecret			string `json:"client_secret"`
	ProjectId				string `json:"project_id"`
	AuthURI					string `json:"auth_uri"`
	TokenURI				string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	RedirectURIs			[]string `json:"redirect_uris"`
}
var googleCredentialsJsonChildSetting = &googleCredentialsJsonChild{}

var GoogleCredentialsByte []byte

func Setup() {
	log.Println("Start getting environment variables...")

	mapTo("APP", AppSetting)
	mapTo("SERVER", ServerSetting)
	mapTo("TRELLO", TrelloSetting)
	mapTo("GOOGLE_TOKEN_JSON", GoogleTokenJson)
	mapTo("GOOGLE_DRIVE", GoogleDriveSetting)
	mapTo("GOOGLE_CREDENTIALS_JSON", googleCredentialsJsonChildSetting)

	googleCredentialsJsonParentSetting.Installed = *googleCredentialsJsonChildSetting

	GoogleCredentialsByte = initGoogleCredentialsByte()

	if os.Getenv("PORT") != "" {
		ServerSetting.Port = os.Getenv("PORT")
		log.Println("PORT automatic defined by service...")
	}

	log.Printf("Using PORT: %v\n", ServerSetting.Port)

	printEnvVars()
}

func mapTo(section string, v interface{}) {
	log.Printf("Reading envvars: %v\n", section)
	err := envconfig.Process(section, v)
	if err != nil {
		log.Fatalf("envconfig.MapTo err: %v\n", err.Error())
	}
}

func initGoogleCredentialsByte() []byte {
	out, err := json.Marshal(googleCredentialsJsonParentSetting)
	if err != nil {
		log.Println("Error marshal credential.json: " + err.Error())
		panic(err)
	}
	return out
}

func printEnvVars() {
	if ServerSetting.RunMode == "debug" {
		fmt.Println("--APP")
		format := "Title: %v\nMinDay: %d\n"
		_, err := fmt.Printf(format, AppSetting.Title, AppSetting.MinDay)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("--SERVER")
		format = "RunMode: %v\nPort: %v\n"
		_, err = fmt.Printf(format, ServerSetting.RunMode, ServerSetting.Port)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("--TRELLO")
		format = "AppKey: %v\nToken: %v\nBoardId: %v\nListNumber: %d\n"
		_, err = fmt.Printf(format, TrelloSetting.AppKey, TrelloSetting.Token, TrelloSetting.BoardID, TrelloSetting.ListNumber)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("--GOOGLE_TOKEN_JSON")
		format = "AccessToken: %v\nTokenType: %v\nRefreshToken: %v\nExpiry: %v\n"
		_, err = fmt.Printf(format, GoogleTokenJson.AccessToken, GoogleTokenJson.TokenType, GoogleTokenJson.RefreshToken, GoogleTokenJson.Expiry)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("--GOOGLE_DRIVE")
		format = "FolderId: %v\n"
		_, err = fmt.Printf(format, GoogleDriveSetting.FolderId)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("--GOOGLE_CREDENTIALS_JSON")
		format = "FirstObject: %v\nProjectId: %v\nClientId: %v\nClientSecret: %v\nAuthURI: %v\nTokenURI: %v\nAuthProviderX509CertUrl: %v\n"
		_, err = fmt.Printf(format, googleCredentialsJsonParentSetting.Installed.ProjectId, googleCredentialsJsonParentSetting.Installed.ClientID,
			googleCredentialsJsonParentSetting.Installed.ClientSecret, googleCredentialsJsonParentSetting.Installed.AuthURI,
			googleCredentialsJsonParentSetting.Installed.TokenURI, googleCredentialsJsonParentSetting.Installed.AuthProviderX509CertUrl)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("RedirectURIs:")
		for _, u := range googleCredentialsJsonParentSetting.Installed.RedirectURIs {
			fmt.Printf("  %s\n", u)
		}
	}
}
