#!/bin/bash
find . -name "*.wsd" | while read src; do
  dir=`dirname $src`
  fileName=`basename $src .wsd`
  dest="$dir/$fileName.txt"
  rm $dest
  cat $src > $dest
done
