#!/usr/bin/env bash

set -e

if [[ -d tmp ]]
then
  echo 'TMP ALREADY EXIST! EXITING...'
  exit 0
else
  echo 'THIS IS GOING TO TAKE A WHILE'

  mkdir tmp
  
  for i in {1..6000}
  do
    cat 'tmp.logs.log' >> 'tmp/logs.log'
  done

  echo "DONE! $(ls -lahg tmp | grep logs.log)"
fi
