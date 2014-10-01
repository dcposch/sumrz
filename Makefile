
build: bin/sumrz

build-all: \
  bin/sumrz-darwin-386 \
  bin/sumrz-darwin-amd64 \
  bin/sumrz-windows-386 \
  bin/sumrz-windows-amd64 \
  bin/sumrz-linux-386 \
  bin/sumrz-linux-amd64 \
  bin/sumrz-linux-arm

bin/sumrz: src/*.go
	go build -o bin/sumrz src/*.go

bin/sumrz-%: src/*.go
	$(eval EXT := $(subst -, ,$(*)))
	GOOS=$(word 1, $(EXT)) GOARCH=$(word 2, $(EXT)) go build -o bin/sumrz-$* src/*.go

