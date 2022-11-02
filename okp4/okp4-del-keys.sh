#!/bin/bash

PREFIX=${1} 	# Prefix for new wallets
PASSWORD=${2} 	# Current wallet Password

FILENAME=${PREFIX}.keys

OKP4BIN=okp4d
KEYRING_BACKEND=file #

echo "Будут удалены все кошельки из файла $FILENAME"

FILE=$(<$FILENAME)

LIST=$(echo $FILE | jq -r '.[] .name')

for name in $LIST
do
	echo "Удаляем ключ $name"
	echo ${PASSWORD} | $OKP4BIN keys delete ${name} --keyring-backend=${KEYRING_BACKEND} -y
done

#rm $FILENAME
#echo "Файл $FILENAME удален"