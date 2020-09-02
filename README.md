# About

An interface for non-user Trello, in the type of a form, to create cards in Trello's board.

The users don't need:
- Trello account
- know using Trello

# Motivation

I needed an application that allows people outside the Marketing Team to open requests for the team and to be placed on the Trello board.
The @baturorkun's project [expenditure-trello](https://github.com/baturorkun/expenditure-trello) inspired me to code this one.
And an opportunity to learn GoLang!

# Before Running

## Trello

Visti [https://trello.com/app-key](https://trello.com/app-key) to generate your App Key and Token.

## Google Drive

Visit [https://developers.google.com/drive/api/v3/quickstart/go](https://developers.google.com/drive/api/v3/quickstart/go) and enable the drive API for your Google Account (Step 1). After doing this, download the credentials.json file.

FAZER O STEP DE PEGAR O TOKEN.JSON, TALVEZ EU TENHA Q FAZER ISSO NO CÓDIGO??

# Build and Running

## Environment Variables

| VAR | TYPE | REQUIRE | DESCRIPTION |
| ------ | --- | ----------- | ----- |
| APP_TITLE  | String | YES | Application Title |
| APP_MINDAY | Int    | YES | [AINDA FAZER ISSO](AINDA FAZER ISSO) |
| SERVER_PORT    | String | YES | Set port number |
| SERVER_RUNMODE | String | YES | `prod` and `debug` (show more logs)  |
| TRELLO_APPKEY     | String | YES | Trello App Key get [here](https://github.com/rodrigoma/form-for-trello/blob/master/README.md#trello)  |
| TRELLO_TOKEN      | String | YES | Trello Token get [here](https://github.com/rodrigoma/form-for-trello/blob/master/README.md#trello)  |
| TRELLO_BOARDID    | String | YES | Every board has a ID on Trello like `KE4wqorD`. You can see it on the URL when then board url is open. `Ex: https://trello.com/b/KE4wqorD/boardname` |
| TRELLO_LISTNUMBER | Int    | YES | Boards has lists like doing, todo, done. ListNumber is order number of the lists and starts with 0 |
| GOOGLE_DRIVE_FOLDERID | String | YES | Every folder has a ID on Drive like `14UgHD-Jhd8dDD`. You can see it on the URL when then folder is open. `Ex: https://drive.google.com/drive/u/1/folders/14UgHD-Jhd8dDD` |
| GOOGLE_CREDENTIALS_JSON_CLIENTID                | String | YES | Field `client_id` on credentials.json file |
| GOOGLE_CREDENTIALS_JSON_CLIENTSECRET            | String | YES | Field `client_secret` on credentials.json file |
| GOOGLE_CREDENTIALS_JSON_PROJECTID               | String | YES | Field `project_id` on credentials.json file |
| GOOGLE_CREDENTIALS_JSON_AUTHURI                 | String | YES | Field `auth_uri` on credentials.json file |
| GOOGLE_CREDENTIALS_JSON_TOKENURI                | String | YES | Field `token_uri` on credentials.json file |
| GOOGLE_CREDENTIALS_JSON_AUTHPROVIDERX509CERTURL | String | YES | Field `auth_provider_x509_cert_url` on credentials.json file |
| GOOGLE_CREDENTIALS_JSON_REDIRECTURIS            | String | YES | Field `redirect_uris` on credentials.json file |
| GOOGLE_TOKEN_JSON_ACCESSTOKEN  | String | YES | Field `access_token` on token.json file |
| GOOGLE_TOKEN_JSON_TOKENTYPE    | String | YES | Field `token_type` on token.json file |
| GOOGLE_TOKEN_JSON_REFRESHTOKEN | String | YES | Field `refresh_token` on token.json file |
| GOOGLE_TOKEN_JSON_EXPIRY       | String | YES | Field `expiry` on token.json file |

You can see examples in the [env.list](https://github.com/rodrigoma/form-for-trello/blob/master/env.list) file

## Development
```
go run main.go
```

## Local
```
go build .
go install
./formfortrello
```

## Docker

Set the env vars in [env.list](https://github.com/rodrigoma/form-for-trello/blob/master/env.list) file
```
docker build . -t formfortrello

docker run -d \
    --name formfortrello \
    --env-file env.list \
    -p 9000:9000 \
    formfortrello
```

## Usage

Open http://localhost:9000 on your browser.