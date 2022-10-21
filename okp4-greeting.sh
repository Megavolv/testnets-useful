#!/bin/bash

WALLET=${1} 	# Wallet name
PASSWORD=${1} 	# Wallet Password 

while true
do
    while read p; do
    seq=`okp4d q account okp41u376ksxhpytfm63rh699w6as95ekfk2mlx5ntr --node tcp://localhost:27657 | grep sequence | awk '{print $2}' | tr -d \"`

    echo "$PASSWORD" | okp4d tx bank send kiwiwallet $p 1uknow --from kiwiwallet --keyring-backend file --node tcp://localhost:27657 --chain-id okp4-nemeton --generate-only --note "Hello from kiwi" > one.json

    echo "$PASSWORD" | okp4d tx sign one.json --offline --keyring-backend file --node tcp://localhost:27657 --from "$WALLET" --chain-id okp4-nemeton --account-number 65 --sequence "$seq" > one_signed.json

    echo "$PASSWORD" | okp4d tx broadcast one_signed.json --from "$WALLET" --keyring-backend file --node tcp://localhost:27657 --chain-id okp4-nemeton -y

    done <adr.txt
done
