#!/bin/bash

KEYSFILE=${1} # Файл со списком адресов в json
DESTFILE=${2} # Выходной файл

OKP4BIN=okp4d

COLLECTOR='{"body":{"messages":[{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"okp41u376ksxhpytfm63rh699w6as95ekfk2mlx5ntr","to_address":"okp41r8rflpv2jnggymk09kvut4dngkv3j34z5xn62u","amount":[{"denom":"uknow","amount":"1"}]}],"memo":"Very-very big tx from kiwi.","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""},"tip":null},"signatures":[]}'

ADDRESSES=$(cat $KEYSFILE | jq '.[].address')

NUM=$(cat $KEYSFILE | jq '.[].address' | wc -l)

I=0


for address in ${ADDRESSES}; do
    let I+=1
    echo "Адрес $I из $NUM"

	MESSAGE='[{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"okp41u376ksxhpytfm63rh699w6as95ekfk2mlx5ntr","to_address":'$address',"amount":[{"denom":"uknow","amount":"1"}]}]'
	COLLECTOR=$(echo $COLLECTOR | jq ".body.messages += $MESSAGE")
	
	if (( "$I" == '100'))
    then
        break
    fi

	
done

let "GAS = 200000*$NUM"

COLLECTOR=$(echo $COLLECTOR | jq ".auth_info.fee.gas_limit = $GAS")

echo $COLLECTOR > $DESTFILE
