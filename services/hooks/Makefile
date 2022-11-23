start:
    ./dist/main

build:
    go build -o ./dist/main

serve:
    nodemon --exec go run main.go --signal SIGTERM

docs:
    docsify init public

ngrok:
    ngrok http 8080

telegram-football:
    go run scripts/telegram/messages/football/main.go

telegram-forex:
    go run scripts/telegram/messages/forex/main.go

telegram-news:
    go run scripts/telegram/messages/news/main.go

telegram-weather:
    go run scripts/telegram/messages/weather/main.go
