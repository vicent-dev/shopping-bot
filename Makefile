include .env

run:
	go run ./cmd/bot/main.go

build:
	CGO_ENABLED=0 go build  -o shopping-bot ./cmd/bot/main.go

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run

deploy:
	make build
	cp shopping-bot.service shopping-bot-backup.service
	sed -i "s/DEPLOY_USER/$(DEPLOY_USER)/g" shopping-bot.service
	sed -i "s/DEPLOY_PATH/$(DEPLOY_PATH)/g" shopping-bot.service
	scp shopping-bot.service $(DEPLOY_USER)@$(DEPLOY_SEVER):/etc/systemd/system/
	scp shopping-bot $(DEPLOY_USER)@$(DEPLOY_SEVER):$(DEPLOY_PATH)
	ssh $(DEPLOY_USER)@$(DEPLOY_SEVER) 'systemctl daemon-reload && systemctl restart shopping-bot'
	mv shopping-bot-backup.service shopping-bot.service