#!/usr/bin/env bash


export GO111MODULE=on
export GOPROXY=https://athens.azurefd.net
#export GOPROXY=http://goproxy.io
#export GOPROXY=http://mirrors.aliyun.com/goproxy
#export GOPROXY=https://gocenter.io

go mod tidy
go mod vendor
