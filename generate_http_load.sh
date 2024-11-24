#!/bin/bash

function check_vegeta_binary() {
  command -v vegeta > /dev/null
  if [[ $? != 0 ]]; then
    echo 'Error: vegeta is not installed.'
    exit 1
  fi
}

function main() {
    check_vegeta_binary
    echo "GET http://localhost:8080" | vegeta attack -rate 100 -duration=15s | vegeta report
}

main