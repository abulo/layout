#!/bin/bash
path=$(cd `dirname $0`; pwd)
go-bindata -o=initial/view/view.go -pkg=view  ./view/...