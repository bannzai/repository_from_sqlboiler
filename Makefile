

dry-run:
	go run main.go

build:
	go build -o repository_from_sqlboiler

source: 
	cp $(SOURCE_SQLBOILER_PATH)/*.go local/source
