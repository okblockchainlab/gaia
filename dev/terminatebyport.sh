#!/bin/bash

NAME=$1
MYNAME="terminatebyport.sh"

if [ $# -eq 0 ]; then
    echo "$MYNAME <process name>"
    exit
fi

lsof -i -n -P|grep "$NAME"|grep -v grep |grep -v $MYNAME |awk '{print "kill -9 "$2", "$1}'
lsof -i -n -P|grep "$NAME"|grep -v grep |grep -v $MYNAME |awk '{print "kill -9 "$2}' | sh
#ps -ef|grep "$NAME"|grep -v grep |grep -v $MYNAME |awk '{print "kill -9 "$2}' | sh
echo "All <$NAME> killed!"


