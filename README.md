# repository_from_sqlboiler
repository_from_sqlboiler can be generated `Go` source code from [sqlboiler](https://github.com/volatiletech/sqlboiler)'s generated `Go` code.

## Setup
It should prepare`local/source` and `local/destination`directory and put on [sqlboiler](https://github.com/volatiletech/sqlboiler)'s generated `Go` model code.
```bash
$ mkdir -p local/source local/destination.
$ export SOURCE_SQLBOILER_MODEL_PATH=$GOPATH/src/github.com/YOUR_APP/model
$ cp $SOURCE_SQLBOILER_MODEL_PATH/*.go local/destination
```

## Run
If you want to try ./repository_from_sqlboiler, You exec `make dry-run`. It is generated file written by`Go` code to `local/destination/repository.go`.
```bash
$ make dry-run
```

