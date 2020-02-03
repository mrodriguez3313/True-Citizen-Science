#!/bin/bash
docker stop ipfsnode
docker start ipfsnode
./autoUpload.sh &
./script.sh &

sleep(20)
./kill_scripts.sh
