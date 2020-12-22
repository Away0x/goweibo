APP_NAME = "goweibo"
PROXY = "https://goproxy.cn"

default:
	go build -o ${APP_NAME}
	# env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	env GOPROXY=${PROXY} go mod download

version:
	go run main.go version

dev:
  # go install -u github.com/cosmtrek/air
	env GOWEIBO_APP_RUNMODE=development air -c air.toml

deploy:
	make && ./bin/deploy restart

api-doc:
  # go install -u github.com/swaggo/swag/cmd/swag
	swag init -g routes/api.go

mock:
	env GOWEIBO_APP_RUNMODE=development go run main.go mock

api-test:
	env GOWEIBO_APP_RUNMODE=test go test ./test/controllers/api/...

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make version - print version"
	@echo "make dev - run go fresh"
	@echo "make deploy - deploy app"
	@echo "make mock - mock data"
	@echo "make api-test - test api"
	@echo "make doc - generate swagger api docs"
	@echo "make clean - remove binary file"
