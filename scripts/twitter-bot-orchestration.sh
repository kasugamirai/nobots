#!/usr/bin/env zsh

#export TWITTER_USERNAME=your_twitter_username
#export TWITTER_PASSWORD=your_twitter_password

CURDIR=$(cd $(dirname $0); pwd)
BinaryName=nobot
echo "$CURDIR/bin/${BinaryName}"

$CURDIR/bin/${BinaryName} twt init
while true
do
  $CURDIR/bin/${BinaryName} twt fetch-publish -c 10

  echo 'sleep for a while'
  # 60 seconds * 30 = 1800
  sleep 1800
done