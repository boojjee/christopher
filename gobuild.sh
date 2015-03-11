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
sudo supervisorctl reload
tail /var/log/long.out.log