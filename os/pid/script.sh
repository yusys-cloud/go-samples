#!/bin/sh
sleep 1
echo "I'm the script with pid $$"
for i in 1 2 3; do
        sleep 1
        echo "Still running $$"
done