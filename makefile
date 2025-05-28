# for updating the version of the project and pushing the tag to the repository 
# NEVER GO BEYOND 1.0.0 (breaking change use 0.x.0 ) 
# keep the 0.x.x -> 1.x.x versioning in the v1 directory from 1.0.0 -> 2.0.0 make v2 folder and new go.mod file in it

VERSION = 0.0.3

MODULE_NAME := github.com/fengdotdev/golibs-funcs


#my GOLIBS For all projects Some are in development, some are stable, some are experimental, some are hidden NOT ALL ARE USED IN THIS PROJECT


# Core Libraries -------------------------------------------------------------------------------------

# Traits: Library for handling reusable traits (interfaces) in Go.
TRAITS = github.com/fengdotdev/golibs-traits

# Testing: Library for testing utilities in Go.
TESTING = github.com/fengdotdev/golibs-testing

# kwowledge: Library for knowledge management in Go. aka info like countries, phones, etc.
KWOW = github.com/fengdotdev/golibs-kwowledge

# CommonTypes: Library for common types used across various Go projects.
COMMON = github.com/fengdotdev/golibs-commontypes

# Dummy: Library for dummy data generation in Go. using common types and traits
DUMMY = github.com/fengdotdev/golibs-dummy

# Funcs set of utility functions simplifying standard library operations.
FUNCS = github.com/fengdotdev/golibs-funcs


# Filesystem Libraries -------------------------------------------------------------------------------------

# Nativedrive: Library for native drive operations in Go.
NDRIVE = github.com/fengdotdev/golibs-nativedrive

# VDrive: Library for virtual drive operations in Go.
VDRIVE = github.com/fengdotdev/golibs-vdrive


# 1DriveClient: Library for interacting with OneDrive in Go.
ONEDRIVE = github.com/fengdotdev/golibs-1driveclient


# Cross-platform Libraries --------------------------------------------------------------------------------

# Bridge: Library for bridging across different languages and platforms in Go. eg: android, ios,
BRIDGE = github.com/fengdotdev/golibs-bridge


# Math And Machine Learning Libraries --------------------------------------------------------------------

# ML: Library for machine learning utilities in Go.
ML = github.com/fengdotdev/golibs-ml

# Statistics: Library for statistical analysis in Go.
STAT = github.com/fengdotdev/golibs-statistics

# LinealAlgebra: Library for linear algebra operations in Go.
LA = github.com/fengdotdev/golibs-linealalgebra


# Data Processing Libraries ------------------------------------------------------------------------------

# DataContainer: Library for handling data containers in Go.
DC = github.com/fengdotdev/golibs-datacontainer


# Stream: Library for stream processing in Go.
STREAM = github.com/fengdotdev/golibs-stream

# Options: Library for handling options in Go.
OPTIONS = github.com/fengdotdev/golibs-options

# Future: Library for handling futures in Go.
FUTURE = github.com/fengdotdev/golibs-future


# Async: Library for handling asynchronous operations in go for wasm and embedded systems.
ASYNC = github.com/fengdotdev/golibs-async


# Database

# CoipoDB: Library for Key-Value database in Go.
COIPODB = github.com/fengdotdev/golibs-coipodb

# JanusDB: Library Relational database in Go.
JANUSDB = github.com/fengdotdev/golibs-janusdb

# VectorDB : Library for a vector database in Go.
VECTORDB = github.com/fengdotdev/golibs-vectordb

# JSONDB: Library for a JSON database in Go.
JSONDB = github.com/fengdotdev/golibs-jsondb


# UI

# UtilityCSS: Library for utility-first CSS in Go.
CSS = github.com/fengdotdev/golibs-utilitycss

# StaticPages: Library for handling static web pages in Go.
STATIC = github.com/fengdotdev/golibs-staticpages

# CoipoWASM: Library for WebAssembly in Go. 
COIPOWASM = github.com/fengdotdev/golibs-coipowasm

# CoipoComponents: Library for reusable GUI components in Go.
COIPOCOMPONENTS = github.com/fengdotdev/golibs-coipocomponents

.PHONY: sand init helloworld fix get tag test

sand: 
	go run cmd/sandbox/main.go


# create a new go module
init:
	go mod init $(MODULE_NAME)
	make get
	make helloworld


# create some directories for the project
helloworld:
# create  cmd directory
	mkdir -p cmd
	mkdir -p cmd/sandbox
# keep the 0.x.x -> 1.x.x versioning in the v1 directory from 1.0.0 -> 2.0.0 make v2 folder and new go.mod file in it
	mkdir -p v1
# create the main.go file
	echo "package main" > cmd/sandbox/main.go
	echo "import \"fmt\"" >> cmd/sandbox/main.go
	echo "func main() {" >> cmd/sandbox/main.go
	echo "fmt.Println(\"Hello, world!\")" >> cmd/sandbox/main.go
	echo "}" >> cmd/sandbox/main.go

fix:	
# fix the go.mod file
	go mod tidy
	go mod vendor
	go mod verify
	

# update the go.mod file with the latest dependencies
get:
	go get $(TRAITS)@latest
	go get $(TESTING)@latest

# update the version of the project
tag:
		git tag v${VERSION} && git push origin v${VERSION}

# run all tests
test:
	go clean -testcache
	go test -count=1 -v ./...



