#!/bin/bash

PREFIX=${1} 	# Prefix for new wallets
PASSWORD=${2} 	# Current wallet Password
WALLETSNUM=${3} # Number of wallets generated
KEYRING_BACKEND=${4}

	if [ -z "$KEYRING_BACKEND" ]
	then
 	   KEYRING_BACKEND="file"
	fi

FILENAME=${PREFIX}.keys

echo "Будет сгенерировано $WALLETSNUM кошельков. Результат сохранится в файле $FILENAME"


OKP4BIN=okp4d
KEYRING_BACKEND=file #

LIST="[]"

for (( i=1; i<=$WALLETSNUM; i++ ))
do

    NEW_KEY=$(echo ${PASSWORD} | $OKP4BIN keys add ${PREFIX}${i} --keyring-backend=${KEYRING_BACKEND} --output json)
    
    name=${PREFIX}${i} #$(echo $NEW_KEY | jq -r '.name')
    address=$(echo $NEW_KEY | jq -r '.address')
    mnemonic=$(echo $NEW_KEY | jq -r '.mnemonic')
    pubkey=$(echo $NEW_KEY | jq '.pubkey')

	echo "Кошелек: $name. Публичный ключ: $address"     
    
    #LIST=$(echo $LIST | jq  ". += [{\"name\":\"${name}\", \"address\":\"${address}\", \"mnemonic\":\"${mnemonic}\", \"pubkey\":${pubkey}}]")
    LIST=$(echo $LIST | jq  ". += [{\"name\":\"${name}\", \"address\":\"${address}\", \"mnemonic\":\"${mnemonic}\"}]")
done

echo $LIST | jq > $FILENAME


