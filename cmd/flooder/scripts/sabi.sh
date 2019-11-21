#!/usr/bin/env bash

for (( VAR = 0; VAR < 100; ++VAR )); do
    w2="data/addr"${VAR}".txt"
    addrs=""
    for i in `cat ${w2}`
    do
    addrs=${addrs}${i}","
    done
    s=${addrs%%","}
    gaiacli tx send Turning ${s} 1000stake --fees 350stake -yes -b block --gas auto<<EOF
12345678
EOF
    gaiacli tx staking delegate cosmosvaloper16wdljyzt5nvdjwrd06jup9hvc8wwhn5u6yl7w5 100stake ${s} --fees 1000stake -yes -b block --gas auto --from Turning<<EOF
12345678
EOF
    sleep 2
    echo "current:"${VAR}
    #gaiacli query staking delegation cosmos14jl54akhlykp54nhe6gqppmevevxef7frgz0gx cosmosvaloper16wdljyzt5nvdjwrd06jup9hvc8wwhn5u6yl7w5
done