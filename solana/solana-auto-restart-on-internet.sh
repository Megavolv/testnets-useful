#!/bin/sh

echo `date` "Запуск. Скрипт для однократного перезапуска соланы после появления интернета."

echo date

while true
do
    wget -q --spider http://ya.ru
    if [ $? -ne 0 ]; then
        echo `date` "Offline"

        while true
        do
            wget -q --spider http://ya.ru
            if [ $? -eq 0 ]; then
                echo `date` "ya! online"

                service solanad restart

                echo `date` "exiting"
                exit
        fi
        sleep 10
        done

    fi
    sleep 10
done
