#!/usr/bin/env zsh

CURDIR=$(cd $(dirname $0); pwd)
BinaryName=nobot
echo "$CURDIR/bin/${BinaryName}"

while true
do
  $CURDIR/bin/${BinaryName} twt car
done
