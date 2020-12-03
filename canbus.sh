#!/bin/bash
/home/pi/canbus/canbus can0 receive | /home/pi/canbus/analyzer -json 2>&1 | /home/pi/canbus/iptee -u 192.168.199.1 4444
