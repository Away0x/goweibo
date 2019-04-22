APP_NAME = "gin_weibo"

default:
	go build -o ${APP_NAME}
	# env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	go mod download

dev:
	fresh -c ./fresh.conf

mock:
	go run ./main.go -m

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make dev - run go fresh"
	@echo "make mock - mock data"
	@echo "make clean - remove binary file"
