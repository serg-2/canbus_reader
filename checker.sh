#!/bin/bash

target="192.168.199.1"

count=$( ping -c 2 $target | grep icmp* | wc -l )

if [ $count -eq 0 ]
then
    /sbin/reboot
else
    /usr/bin/logger -t "checker" ok
fi
