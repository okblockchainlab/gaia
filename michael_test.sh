#!/usr/bin/env bash

cd /Users/oker/code/src/github.com/cosmos/gaia/dev

./start.sh;tail -f nohup.out
tail -f /Users/oker/code/src/github.com/cosmos/gaia/dev/nohup.out


gaiacli query account cosmos16wdljyzt5nvdjwrd06jup9hvc8wwhn5ulsttz8
gaiacli query account cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk

gaiacli keys add jack
(cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk)

gaiacli tx send Turing cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk 1000stake -yes -b block

gai