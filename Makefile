

dry-run:
	make run main.go

build:
	make build -o repository_from_sqlboiler

source: 
	cp $(SOURCE_SQLBOILER_PATH)/*.go local/source
