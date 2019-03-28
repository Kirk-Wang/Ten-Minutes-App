DOCKER_GO_BUILD=go build -mod=readonly -a -installsuffix cgo -ldflags "$$LD_FLAGS"

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${DOCKER_GO_BUILD} -v -o release/linux/amd64/ten-minutes-app-api

docker:
	docker build -t lotteryjs/ten-minutes-app-api .

test:
	go test -v .