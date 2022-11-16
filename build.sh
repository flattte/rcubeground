#!/bin/bash

f="./Rubik-Cube"
if [ -d ./Rubik-Cube ]; then
    echo "front end present"
else
    git clone git@github.com:Patrick2402/Rubik-Cube.git
fi

if [ $1 == "rf" ]; then
    cd $f && npm run build && cd ..
fi  

exec_name=$(cat ./backend/go.mod | grep module | awk '{print $2}')
cd ./backend && go build
./$exec_name
