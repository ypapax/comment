#!/usr/bin/env bash
set -ex
build() {
  pushd apps/api
  go install
  popd
}

run() {
  build
  api -conf ./conf.local.yaml
}

runc() {
  build
  docker-compose build
  docker-compose up
}

runca() {
  build
  docker-compose build api
  docker-compose up api
}

runcdb() {
  build
  docker-compose build postgres
  docker-compose up postgres
}

testl() {
  pushd test
  go test -v "$@"
  popd
}

"$@"
