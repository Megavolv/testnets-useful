#!/bin/bash

restart() {
	./pigeon --text "$1: It seems massad doesnt work. Try to restart" --bot "fl"
    echo "Massad doesnt work. Try to restart"
    sudo service massad restart
}

NODENAME=${1} 	# название сервера
PASSWORD=${2} 	# пароль от бумажника
MASSABIN=./massa-client # путь к исполняемому файлу с массой

while true
do
# Проверяем текущий статус массы
	STATUS=$($MASSABIN -p $PASSWORD get_status)
	if [[ $STATUS == *"error"* ]]; then
			restart $NODENAME
			sleep 5s
		else
			break
	fi
done

ADDRESS=$($MASSABIN wallet_info -j -p $PASSWORD | jq -r 'keys[0]')

# Если вместо адреса переменная содержит "keys", то это тоже означает, что масса не работает.
if [ "$ADDRESS" == "keys" ]; then
	restart $NODENAME
	exit
fi

#ACTIVE_ROLLS_COUNT=$($MASSABIN -p ${PASSWORD} wallet_info -j | jq ".${ADDRESS}.address_info.active_rolls")
CANDIDATE_ROLLS_COUNT=$($MASSABIN -p ${PASSWORD} wallet_info -j | jq ".${ADDRESS}.address_info.candidate_rolls")

# if [[ $ACTIVE_ROLLS_COUNT -eq 0 && $CANDIDATE_ROLLS_COUNT -eq 0 ]]; then
if [[ $CANDIDATE_ROLLS_COUNT -eq 0 ]]; then
		./pigeon --text "${NODENAME}: need to buy rolls" --bot "fl"
		echo "Need to buy rolls"

		($MASSABIN -p ${PASSWORD} buy_rolls $ADDRESS 1 0)
	else
		echo "No action is needed"
fi
