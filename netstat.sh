#!/bin/bash
while true
do
    sleep 1
    clear
    netstat -anp tcp | \
        awk '{ if ($4 ~ /8080/) { ++S[$NF] } } END { for (a in S) print a, S[a] }'
done
