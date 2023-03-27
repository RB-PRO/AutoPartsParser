all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/AutoPartsParser.git

pull:
	git pull git@github.com:RB-PRO/AutoPartsParser.git

pushW:
	git push https://github.com/RB-PRO/AutoPartsParser.git

pullW:
	git pull https://github.com/RB-PRO/AutoPartsParser.git

build-config:
	go env GOOS GOARCH

build-windows-to-linux:
	set GOARCH=amd64 set GOOS=linux go build .\cmd\main\main.go  

build-linux-to-windows:
	export GOARCH=amd64 export GOOS=windows go build cmd/main/main.go

build-car:
	set GOARCH=amd64
	set GOOS=linux
	export GOARCH=amd64
	export GOOS=linux
	export CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o main ./cmd/main/main.go
	scp main lap.json root@194.87.107.129:go/AutoPartsParser/