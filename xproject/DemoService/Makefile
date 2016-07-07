SERVICE := DemoService
VERSION := 0.0.1

GO_CFLAGS=-ldflags "-s -X main.g_Service=$(SERVICE) \
					   -X main.g_Version=$(VERSION) \
					   -X main.g_BuildDate=`date -u +%Y-%m-%d,%H:%M:%S` \
					   -X main.g_CommitID=`git log -p -1 | sed -n '1p' | cut -d " " -f 2`"

build:
	@echo "building..." && \
	go install $(GO_CFLAGS) && \
	echo "build done"


run: build
	mkdir -p log && \
	$(SERVICE) -c demo.conf

spec:
	@echo "generating swagger spec..." && \
	swagger generate spec -o ./swagger.json && \
	jsonpretty -f ./swagger.json && \
	echo "generate swagger sped done"