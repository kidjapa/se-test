#!/bin/bash

# delete old node process
declare -a portsServices=("3000")

# get length of an array
portArrayLength=${#portsServices[@]}

# Kill all ports if exists
for (( i=1; i<${portArrayLength}+1; i++ ));
do
  echo "Killing port: - "  $i " / " ${portArrayLength} " : " ${portsServices[$i-1]}
  echo "Killing PID: " $(lsof -t -i:${portsServices[$i-1]})
  lsof -n -i:${portsServices[$i-1]} | grep LISTEN | tr -s ' ' | cut -f 2 -d ' ' | xargs kill -9
done