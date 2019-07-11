
clean:
	rm -f ./repository_from_sqlboiler

dry-run: clean build
	./script/development/generate.sh

build:
	go build -o repository_from_sqlboiler

source: 
	cp $(SOURCE_SQLBOILER_PATH)/*.go local/source

test: 
	go test ./pkg/...
