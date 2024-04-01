CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.54.2
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=github.com/romanzimoglyad/inquiry-backend/cmd/app
GRPC_PACKAGE=route256/loms/cmd/grpc_app
all: format build test lint

build: bindir
	env GOOS=linux go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint:
	golangci-lint run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}



install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})
generate:
	protoc -I api/v1 -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7/ \
	--go_out ./pb/api_v1 --go_opt=paths=source_relative \
	--go-grpc_out ./pb/api_v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ./pb/api_v1 --grpc-gateway_opt paths=source_relative \
	api/v1/api.proto


run-all: build
	sudo docker compose up --force-recreate --build
