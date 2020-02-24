#!/bin/bash
# rm -r ../filesdump; rm ../Hashes; rm ../hash_outputs
./script.sh &
sleep 2
./autoUpload.sh &


sleep 10
./kill_scripts.sh
cd ..
pwd
sleep 4
echo ""
echo "Number of files created: "
ls -l filesdump | wc -l
echo ""
sleep 2
echo "extracting hashes: "
sleep 1
./extract_hashes_batch Hashes
echo ""
sleep 6
rm -r filesdump; rm Hashes; rm hash_outputs
mongo
