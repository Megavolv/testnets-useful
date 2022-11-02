#!/bin/bash

MAX_TX=2
MAX_NUM_BLOCK=1
CURR=0

LAST_BLOCK=`curl http://localhost:27657/status --no-progress-meter | jq -r .result.sync_info.latest_block_height`

echo "The last known block = $LAST_BLOCK"

for (( i=20000; i<=$LAST_BLOCK; i++ ))
do

    CURR=`okp4d q block $i --node tcp://localhost:27657 | jq .block.data.txs | jq length`
    
    echo "block=$i. last=$CURR. max_tx=$MAX_TX. max_block=$MAX_NUM_BLOCK"

    if [[ $CURR -gt $MAX_TX ]]; then
        MAX_TX=$CURR
        MAX_NUM_BLOCK=$i
        echo "new max $MAX_TX"
    fi
done
