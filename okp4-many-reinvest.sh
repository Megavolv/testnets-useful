#!/bin/bash

WALLET=${1} 	# Wallet name
PASSWORD=${2} 	# Wallet Password 

while true
do
    while read address; do
        ./okp4-flex-reinvest.sh0 "$WALLET" "$PASSWORD" $address
        sleep 10
    done < validators.txt

sleep 1800
done