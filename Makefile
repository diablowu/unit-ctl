unit-auth:
	go build -a -v -o dist/unit-auth github.com/diablowu/unit-ctl/cmd/auth
unit-deploy:
	go build -a -v -o dist/unit-deploy github.com/diablowu/unit-ctl/cmd/deploy
all: unit-auth unit-deploy
	@echo "all"