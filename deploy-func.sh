#!/bin/bash

FUNCTION_APP_NAME=function-app

if [ -f handler ]; then
    rm main
fi

GOOS=linux GOARCH=amd64 go build -o main

func azure functionapp publish $FUNCTION_APP_NAME