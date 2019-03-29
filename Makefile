DOCKER_GO_BUILD=go build -mod=readonly -a -installsuffix cgo -ldflags "$$LD_FLAGS"

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${DOCKER_GO_BUILD} -v -o release/linux/amd64/api-ten-minutes

docker:
	docker build -t lotteryjs/api-ten-minutes .

test:
	go test -v .