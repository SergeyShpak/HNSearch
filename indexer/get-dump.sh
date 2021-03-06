#!/bin/bash

URL=https://www.dropbox.com/s/duv704waqjp3tu1/hn_logs.tsv.gz?dl=1
DIR=/usr/HNIndexer/data
FILE=$DIR/hn_logs.tsv

if [ -f $FILE ]; then
  exit 0
fi

mkdir -p $DIR
wget -O $FILE.gz $URL && gzip -f -d $FILE.gz
rm -f $FILE.gz