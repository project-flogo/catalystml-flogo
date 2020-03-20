#!/bin/sh
rtn=0
cd action
go test 
cd ..
cd operations
for op  in $(find . -mindepth 2 -maxdepth 2 -type d)
do 
    cd $op
    if ! go test; then
        let "rtn=rtn+1"
    fi
    go test 
    cd ../..
done
return $rtn