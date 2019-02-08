export VERSION = $(shell bash ./version.sh)
export DATE    = $(shell bash ./date.sh)
all:
	go build -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}"
install:
	go install -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}"
