export GO111MODULE=on

all: clean gittools

REVISION := $(shell git describe --tags --match 'v*' --always --dirty 2>/dev/null)
REVISIONDATE := $(shell git log -1 --pretty=format:'%ad' --date short 2>/dev/null)
PKG := github.com/chiyutianyi/gittools/pkg/version
LDFLAGS = -s -w
ifneq ($(strip $(REVISION)),) # Use git clone
	LDFLAGS += -X $(PKG).revision=$(REVISION) \
		   -X $(PKG).revisionDate=$(REVISIONDATE)
endif

gittools: Makefile cmd/*.go pkg/*/*.go
	go build -ldflags="$(LDFLAGS)" -o gittools ./cmd

gittools-macos: Makefile cmd/*.go pkg/*/*.go
	go build -ldflags="$(LDFLAGS)" -o gittools-$(REVISION) ./cmd

gittools-linux: Makefile cmd/*.go pkg/*/*.go
	 GOOS=linux go build -ldflags="$(LDFLAGS)" -o gittools-$(REVISION).x86_64 ./cmd

gittools-linux-latest: Makefile cmd/*.go pkg/*/*.go
	 GOOS=linux go build -ldflags="$(LDFLAGS)" -o gittools-latest.x86_64 ./cmd

clean:
	rm -f gittools gittools-* gittools-*.x86_64