#!/bin/bash

rm -r /root/go/bin/christopher 
git pull
sleep 3
supervisorctl stop all 
tail /var/log/long.out.log
pwd
go install 
go build
sleep 3
NOW=$(date +"%m-%d-%Y %T")
echo "{ 
  build_date: $NOW
}" >> version.json
sudo supervisorctl reload
tail /var/log/long.out.log