
build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/linux/amd64/ten-minutes-app-api

build_linux_i386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -a -o release/linux/i386/ten-minutes-app-api

docker:
	docker build -t kirkwwang/ten-minutes-app-api .

test:
	go test -v .