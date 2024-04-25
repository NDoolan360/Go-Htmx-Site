# Usage:
# make          # install and build
# make build    # builds our `tailwind.css` output file
# make install  # installs the latest version of the `tailwindcss` CLI
# make clean    # removes all binaries and artifacts

.PHONY: build install dev clean

# Get OS & ARCH info
SYSTEM := $(shell uname -s)
ifeq ($(SYSTEM),Linux)
    OS=linux
else ifeq ($(SYSTEM),Darwin)
    OS=macos
endif

PLATFORM := $(shell uname -m)
ifeq ($(PLATFORM),x86_64)
    ARCH=x64
else ifeq ($(PLATFORM),arm64)
    ARCH=arm64
endif

BINARY=tailwindcss-$(OS)-$(ARCH)
CSS_FILE=$(PWD)/public/styles.css
BUILD_FILE=$(PWD)/public/tailwind.css

all: install build

build: $(BUILD_FILE)

$(BUILD_FILE): $(BINARY) $(CSS_FILE)
	./$(BINARY) build -i $(CSS_FILE) -o $(BUILD_FILE) --minify

install: $(BINARY)

$(BINARY):
	curl -sLo $(BINARY) https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(OS)-$(ARCH)
	chmod +x $(BINARY)

dev: $(BUILD_FILE)
	go run main.go

clean:
	rm -vf $(BINARY) $(BUILD_FILE)
