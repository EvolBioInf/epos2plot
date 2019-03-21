export VERSION = $(shell bash ./scripts/version.sh) 
export DATE    = $(shell bash ./scripts/date.sh)
all:
	cd cmd/epos2plot && go build -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}"
	cd cmd/plotSum   && go build -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}"
install:
	cd cmd/epos2plot && go install -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}"
	cd cmd/plotSum   && go install -ldflags "-X main.VERSION=${VERSION} -X main.DATE=${DATE}"
test:
	bash ./scripts/epos2plot.sh
	bash ./scripts/plotSum.sh
clean:
	rm ./cmd/epos2plot/epos2plot
	rm ./cmd/plotSum/plotSum
.PHONY:	doc
doc:	
	make -C doc

