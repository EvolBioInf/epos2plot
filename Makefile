export VERSION = $(shell bash ./scripts/version.sh) 
export DATE    = $(shell bash ./scripts/date.sh)
all:
	cd cmd && go build -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}" epos2plot.go
	cd cmd && go build -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}" plotSum.go
install:
	cd cmd && go install -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}" epos2plot.go
	cd cmd && go install -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}" plotSum.go
test:
	bash ./scripts/epos2plot.sh
	bash ./scripts/plotSum.sh
clean:
	rm ./cmd/epos2plot
	rm ./cmd/plotSum
.PHONY:	doc
doc:	
	make -C doc

