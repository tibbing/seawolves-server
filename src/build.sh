#!/bin/bash
FUNCTION_NAME=$1

build () {
  echo "Building $1..."
  GOOS=linux go build -o main ./app/$1/main.go
  chmod +rwx ./main
  mkdir -p dist
  rm -f ./dist/$1.zip
  zip ./dist/$1.zip ./main
  rm ./main
  chmod +rwx ./dist/$1.zip
}

if [ -z "$FUNCTION_NAME" ]
then
      echo "No function specified, building all"
      for i in $(ls -d ./app/*/); do build $(basename ${i}); done
else
      build $FUNCTION_NAME
fi




