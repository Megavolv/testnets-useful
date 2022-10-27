# Полезные скрипты в нодоводстве

## OKP4

### okp4-fast-gen-keys

Program for quick key generation for okp4. Written in golang. Supports multithreading.

How to use:

	git clone github.com/Megavolv/testnets-useful/
	cd okp4-fast-gen-keys; go build
	bash: date; nice -n 19 ./okp4-fast-gen-keys --name="mykey" --num 1000000 > keys.json; date

This command will generate 1 million keys with name from mykey0 to mykey999999 and save them to keys.json.

For example:

*{"name":"prefix76","type":"local","address":"okp41w496k6pj0f5f9x43vva6enfacngek25s00gyvz","pubkey":"PubKeySecp256k1{026BD1E6385203662198F2EDF7FE4B2316357690D239B1C0021864C99ADAA4BC7C}","mnemonic":"scrap staff wisdom waste romance coast usage pride diet vivid ramp club length inside october dose remove swear soda angry volume expect aisle debris"}*

#### Benchmark

On my computer (6 cores * 2375 Mhz), the generation of 1 million keys takes ~ 5 minutes, so ~ 534 keys/1 core/1 sec.

### okp4-greetings.sh

Отправка минимальной транзакции с приветствием по списку адресов из файла adr.txt с периодичностью в 1 ч.

### okp4-flex-reinvest.sh

Модификация скрипта для реинвестирования. На вход подается название кошелька, пароль и адрес валидатора.

### okp4-many-reinvest.sh

Используется для реинвестирования списка валидаторов.
На вход: название кошелька, пароль.

Список валидаторов берется из *validators.txt*

### okpr-max-tx.sh

Скрипт поиска блока с максимальным количеством транзакций.

### okp4-gen-keys.sh

Генерация заданного количества кошельков.

На вход: перфикс группы паролей (любое слово или буквы), пароль, количество генерируемых кошельков.

Результат сохраняется в кошельке okp4 и в json файле ${prefix}.keys

### okp4-gen-dryrun-keys.sh

Генерация заданного количества кошельков.

На вход: перфикс группы паролей (любое слово или буквы), количество генерируемых кошельков.

Результат сохраняется только в json в файле ${prefix}.keys

### okp4-del-keys.sh

Удаление кошельков, используя файл с массивом json-кошельков (из предыдущего этапа).

### okp4-check-balance.sh

Массовая проверка балансов адресов из файла с массивом json-кошельков. Отображаются не пустые адреса.

### okp4-join-txsmsg.sh

Объединяет все файлы транзакций в один файл.

## Massa

### massa_rols_checker.sh

Простой скрипт для контроля состояния ноды. Проверяет ролы, покупает, перезапускает сервис при необходимости.

Отправляются уведомления в телеграмм с помощью pigeon.

	while true; do sudo ./massa-rols-checker.sh [node name] [wallet password]; sleep 10; done

## Солана

### solana-auto-restart-on-internet.sh

Однократно перезапускает солану при появлении интернета. Запускается заранее, при ожидании технических работ на сервере.
