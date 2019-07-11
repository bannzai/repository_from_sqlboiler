#! /bin/bash
set -eu
set -o pipefail

PWD=`dirname $0`
APP_DIR="$PWD/../../"
cd $APP_DIR

./repository_from_sqlboiler generate --source=local/source/source.go --destination=local/destination/repository.go --template=configs/template/repository.go.tpl
