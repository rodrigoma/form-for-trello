##Form For Trello Integration

This application being developed in GoLang and gives you interface (Form) to create cards in Trello. Users don't need Trello account and don't know using Trello. The system provides a fancy HTML form to save cards to Trello.

All you have to set your trello api key and token in app.ini file. This config file is in the conf directory. At first, rename the file to app.ini.

You can generate api key and token from "https://trello.com/app-key".


###### Another important parameters:

BoardID : Every board has a ID on Trello like "KE4wqorD". You can see it on the URL when then board url is open.

> https://trello.com/b/KE4wqorD/batus.

ListNumber : Boards has lists like doing, todo, done. ListNumber is order number of the lists and starts with 0.

Port : You can change server port in server trunk. 

### Run

> go tun main.go

### Install
```
go build .
go install
```

### Running
```
./formfortrello
```

### Running on Docker
```
docker build . -t formfortrello

docker run \
    -p 9000:9000 \
    -e APP_MINDAY=20 \
    -e APP_TITLE='Solicitação de Arte' \
    -e SERVER_PORT=9000 \
    -e SERVER_RUNMODE=debug \
    -e TRELLO_APPKEY='<APP_KEY>' \
    -e TRELLO_BOARDID='<BOARD_ID>' \
    -e TRELLO_LISTNUMBER=0 \
    -e TRELLO_TOKEN='<TOKEN>' \
    formfortrello

```

### Usage

> Open " http://[your ip]:8000 " on your  browser.


