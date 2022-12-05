serve:
		nodemon --exec go run main.go --signal SIGTERM

start:
		./dist/main

build:
		go build -o ./dist/main

format:
		prisma format && prisma generate

studio:
		prisma studio

docs:
		docsify init public
