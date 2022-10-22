#!/bin/bash

MAX=0
NUM=0

LAST_BLOCK=`curl http://localhost:27657/status --no-progress-meter | jq -r .result.sync_info.latest_block_height`

echo "The last known block = $LAST_BLOCK"

for (( i=1; i<=$LAST_BLOCK; i++ ))
do

    NUM=`okp4d q block $i --node tcp://localhost:27657 | jq .block.data.txs | jq length`
    echo "block=$i. max=$MAX"

    if [[ $NUM -gt $MAX ]]; then
    MAX=$NUM
    echo "new max $MAX"
    fi
done
