#!/bin/bash
FUNCTION_NAME=$1

build () {
  echo "Building $1..."
  GOOS=linux go build -o tmp_build ./app/$1/main.go
  chmod +rwx ./tmp_build
  mkdir -p dist
  zip ./dist/$1.zip ./tmp_build
  rm ./tmp_build
  chmod +rwx ./dist/$1.zip
}

if [ -z "$FUNCTION_NAME" ]
then
      echo "No function specified, building all"
      for i in $(ls -d ./app/*/); do build $(basename ${i}); done
else
      build $FUNCTION_NAME
fi




