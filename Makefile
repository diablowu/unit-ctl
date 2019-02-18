unit-auth:
	go build -a -v -o dist/unit-auth github.com/diablowu/unit-ctl/cmd/auth
unit-deploy:
	go build -a -v -o dist/unit-deploy github.com/diablowu/unit-ctl/cmd/deploy
unit-bot:
	go build -a -v -o dist/unit-bot github.com/diablowu/unit-ctl/cmd/bot
all: unit-auth unit-deploy unit-bot
	@echo "all"