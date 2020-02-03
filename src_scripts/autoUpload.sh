#!/bin/bash
echo "initializing deamon to monitor and upload files to ipfs created in filesdump dir."
inotifywait -m ~/Documents/InsightDC/filesdump -e create -e moved_to |
    while read path action file; do
        echo "The file '$file' appeared in directory '$path' via '$action'"
        ipfs add ~/Documents/InsightDC/filesdump/$file | awk '{split($0,a); print a[2]}' >> ~/Documents/InsightDC/Hashes
    done
