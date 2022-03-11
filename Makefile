export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

build:ydrdsfrpc

# compile assets into binary file
file:
	rm -rf ./assets/frps/static/*
	rm -rf ./assets/frpc/static/*
	cp -rf ./web/frps/dist/* ./assets/frps/static
	cp -rf ./web/frpc/dist/* ./assets/frpc/static

fmt:
	go fmt ./...

frps:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/frps ./cmd/frps

frpc:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/frpc ./cmd/frpc
win:
	env CGO_ENABLED=0 GOOS=windows  go build -trimpath -ldflags "$(LDFLAGS)" -o bin/ydrdc.exe ./cmd/frpc
	env CGO_ENABLED=0 GOOS=windows  go build -trimpath -ldflags "$(LDFLAGS)" -o bin/ydrds.exe ./cmd/frps
test: gotest

gotest:
	go test -v --cover ./assets/...
	go test -v --cover ./cmd/...
	go test -v --cover ./client/...
	go test -v --cover ./server/...
	go test -v --cover ./pkg/...

e2e:
	./hack/run-e2e.sh

e2e-trace:
	DEBUG=true LOG_LEVEL=trace ./hack/run-e2e.sh

alltest: gotest e2e
	
clean:
	rm -f ./bin/frpc
	rm -f ./bin/frps
	rm -f ./bin/ydrds
	rm -f ./bin/ydrdc
