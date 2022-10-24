#!/bin/bash

PREFIX=${1} 	# Prefix for keys wallet

FILENAME=${PREFIX}.keys

OKP4BIN=okp4d
KEYRING_BACKEND=file #

jq -c '.[]' $FILENAME | while read i; do

    address=$(echo $i | jq -r '.address')

	RES=$($OKP4BIN query bank balances  ${address} --node tcp://localhost:27657  --output=json | jq -r '.balances[].amount')
	
	if [ -z "$RES" ]
	then
 	   continue
	fi

	echo $i | jq -r '.name'
	echo $address
	echo $RES
done

