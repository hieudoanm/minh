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

send-football:
    go run scripts/telegram/messages/football/main.go

send-forex:
    go run scripts/telegram/messages/forex/main.go

send-news:
    go run scripts/telegram/messages/news/main.go

send-weather:
    go run scripts/telegram/messages/weather/main.go
