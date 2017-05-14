SOURCE_FILES?=$$(glide novendor)

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/stretchr/testify
	dep ensure

test:
	gotestcover -coverprofile=coverage.out $(SOURCE_FILES) -run .

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

build:
	go build