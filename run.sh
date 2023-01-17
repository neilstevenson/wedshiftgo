#!/bin/bash
export HOST_IP=$1
if [ "${HOST_IP}" == "" ]
then
 echo Needs HOST_IP as arg
 exit 1
fi
go run cli/client/main.go 
