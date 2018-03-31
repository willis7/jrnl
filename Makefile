NAME = jrnl
PWD := $(MKPATH:%/Makefile=%)

clean:
	cd "$(PWD)"
	rm -rf vendor

install:
	go get -v github.com/Masterminds/glide
	cd $GOPATH/src/github.com/Masterminds/glide && git checkout tags/v0.13.1 && make install && cd -

build:
	glide install
	go build -race -o $(NAME)

start:
	./$(NAME)

test:
	go test -race -v $(shell go list ./... | grep -v /vendor/)

coverage:
	go test -race -cover -v $(shell go list ./... | grep -v /vendor/)

vet:
	go vet $(shell go list ./... | grep -v /vendor/)

lint:
	go get -u golang.org/x/lint/golint
	@golint ./... | grep -v vendor | grep -v .pb. | tee /dev/stderr
