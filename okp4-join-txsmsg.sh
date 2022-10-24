#!/bin/bash

DIRECTORY=${1} 	# Директория с json файлами. В каждом файле - по одной транзакции.
RESULTFILE=${2} # Результирующий файл

NUM=$(find "tx" -name '*.json' | wc -l)

if (( $NUM < 2 )); then
	echo "Найдено недостаточно файлов"
	exit
fi

ALLLIST=$(find "tx" -name '*.json')

RESULTMSG=$(cat $(echo "${ALLLIST}" | head -1)) # Первый файл используется для собирания всех tx

WORKLIST=$(find "tx" -name '*.json' | tail +2)

for name in ${WORKLIST}; do
    M2M=$(cat $name | jq '.body.messages')
    RESULTMSG=$(echo $RESULTMSG | jq ".body.messages += [$M2M]")
done

let "GAS = 200000*$NUM"

RESULTMSG=$(echo $RESULTMSG | jq ".auth_info.fee.gas_limit = $GAS")
echo $RESULTMSG > $RESULTFILE
