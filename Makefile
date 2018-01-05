
install:
	@go install ./cmd/gabel

build:
	@GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/gabel ./cmd/gabel
	@GOOS=linux GOARCH=386 go build -o build/linux/386/gabel ./cmd/gabel

	@GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/gabel.exe ./cmd/gabel
	@GOOS=windows GOARCH=386 go build -o build/windows/386/gabel.exe ./cmd/gabel

	@GOOS=darwin GOARCH=amd64 go build -o build/darwin/amd64/gabel ./cmd/gabel
	@GOOS=darwin GOARCH=386 go build -o build/darwin/386/gabel ./cmd/gabel

clean:
	rm -rf ./build
