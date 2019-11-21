#!/usr/bin/env bash

for (( VAR = 0; VAR < 15; ++VAR )); do
    w1="prikey"${VAR}".txt"
    w2="addr"${VAR}".txt"
    flooder account save_to_file --number 3000 --mnemonic_file_path ${w1} --address_file_path ${w2}
done

