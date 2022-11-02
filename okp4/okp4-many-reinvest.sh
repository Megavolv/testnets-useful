#!/bin/bash

WALLET=${1} 	# Wallet name
PASSWORD=${2} 	# Wallet Password 

while true
do
    while read address; do
    echo "$WALLET" "$PASSWORD" $address
        ./okp4-flex-reinvest.sh "$WALLET" "$PASSWORD" $address
        
		# Очень-очень грязный хак, предотвращающий забивание памяти.
		# Каждый запуск okp4d оставляет два новых таких процесса.
		# Причина неясна.
	    killall /usr/bin/gnome-keyring-daemon 
	    killall /usr/bin/dbus-daemon

        sleep 10
    done < validators.txt

sleep 1800
done