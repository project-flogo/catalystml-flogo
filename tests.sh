#!/bin/sh
cd action
go test 
cd ..
cd operations
for op  in $(find . -mindepth 2 -maxdepth 2 -type d)
do 
    cd $op
    go test 
    cd ../..
done