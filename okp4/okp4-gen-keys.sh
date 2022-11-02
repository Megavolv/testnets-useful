#!/bin/bash

PREFIX=${1} 	# Prefix for new wallets
PASSWORD=${2} 	# Current wallet Password
WALLETSNUM=${3} # Number of wallets generated
KEYRING_BACKEND=${4}
OKP4BIN=okp4d

FILENAME=${PREFIX}.keys

if [ -z "$KEYRING_BACKEND" ]
then
   KEYRING_BACKEND="file"
fi
echo "Start at `date`"

echo "Будет сгенерировано $WALLETSNUM кошельков. Результат сохранится в файле $FILENAME"

LIST="[]"

for (( i=1; i<=$WALLETSNUM; i++ ))
do
    NEW_KEY=$($OKP4BIN keys add ${PREFIX}${i} --keyring-backend=${KEYRING_BACKEND} --output json)
    LIST=$(echo $LIST | jq  ". += [$NEW_KEY]")
    
	if (( $i%100 == 0))
	then
		date
		echo `date; echo "- added next 100 keys"`
		echo $LIST | jq > $FILENAME # backup save every 1000 keys
	fi
    
done

echo $LIST | jq > $FILENAME