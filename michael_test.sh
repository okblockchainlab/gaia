#!/usr/bin/env bash

cd /Users/oker/code/src/github.com/cosmos/gaia/dev

./start.sh;tail -f nohup.out
tail -f /Users/oker/code/src/github.com/cosmos/gaia/dev/nohup.out

nohup flooder account save_to_file --number 300000 >> ./output.log 2>&1 &


gaiacli query account cosmos16wdljyzt5nvdjwrd06jup9hvc8wwhn5ulsttz8
gaiacli query account cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk

gaiacli keys add jack
(cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk)

gaiacli tx send Turing cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk 1000stake -yes -b block

gaiacli query staking validators

gaiacli tx staking delegate cosmosvaloper16wdljyzt5nvdjwrd06jup9hvc8wwhn5u6yl7w5 10stake --from jack -yes -b block


gaiacli query account cosmos1wfsqpu02dz5exmzn25yc899435js6wnd9nkmn4


gaiacli tx send Turing


nohup gaiad start --log_level consensus:info,state:info --home cache/node0/gaiad >> ./output1.log 2>&1 &

gaiad start --log_level consensus:info,state:info --home cache/node0/gaiad


gaiacli tx send Turing cosmos18dsnjdq53jmk8c7d7w96fczpafjr6xjp5zyrjk 1000stake -yes -b block
