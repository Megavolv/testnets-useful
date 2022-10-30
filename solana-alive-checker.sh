#!/bin/bash

name=${1} // Название сервера

while :
do
    status=$(service solanad status  | grep dead | wc -l)

    if ((status >= 1)); then
        ./pigeon --text "$name: solana service is dead;" --bot fl
	sleep 5m
    fi
    sleep 1s
done