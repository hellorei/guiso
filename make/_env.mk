export NO_BROWSERSLIST_UPDATE=1

# WGO
WGO_VERSION=latest

# Variables
FRONTEND_SOURCES := $(shell find . -type f \( -name "*.css" -o -name "*.ts" \))
BINARY_NAME=guiso-app
MAIN_PACKAGE=./app
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
LDFLAGS=-ldflags "-w -s"
MAKEFLAGS += --silent
GOPATH ?= $(HOME)/go
ENV=development
VERSION=0.0.1

# Detect OS and architecture
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

# Default to Linux if UNAME_S is empty
ifeq ($(UNAME_S),)
	UNAME_S := $(shell uname -s)
endif

# Function to kill all child processes
define kill_children
	@echo "Killing all child processes..."
	@pkill -P $$$$
endef