run:
	go run ./cmd/bot/main.go

build:
	go build ./cmd/bot/main.go

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run
