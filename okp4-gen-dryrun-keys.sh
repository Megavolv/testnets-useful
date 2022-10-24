#!/bin/bash

PREFIX=${1} 	# Prefix for new wallets
WALLETSNUM=${2} # Number of wallets generated
OKP4BIN=okp4d

FILENAME=${PREFIX}.keys

LIST="[]"

for (( i=1; i<=$WALLETSNUM; i++ ))
do
	echo "Wallet $i of $WALLETSNUM"
    NEW_KEY=$($OKP4BIN keys add ${PREFIX}${i} --keyring-backend=test --dry-run --output json)
    LIST=$(echo $LIST | jq  ". += [$NEW_KEY]")

	if (( $i%2 == 0))
	then
		echo $LIST | jq > $FILENAME # backup save every 1000 keys
	fi
    
done

echo $LIST | jq > $FILENAME

echo "Created file $FILENAME"
