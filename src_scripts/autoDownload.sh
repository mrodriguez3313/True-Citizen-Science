#!/bin/bash
ipfs pubsub sub Hashes > pub_Hashes &
# I want to add only what was newly added into the file
inotifywait -m ~/Documents/InsightDC/src_scripts/pub_Hashes -e modify -e create |
    while read path action file; do
        if tail -n1 $path | grep Qm; then
            echo "A hash was recieved and saved to pub_Hashes '$file' in directory '$path' "
            echo "" >> pub_Hashes
            echo "`tail -n2 pub_Hashes | head -n1`" > Hashes
            # cat pub_Hashes >> Hashes
            cat /dev/null > pub_Hashes
            ~/Documents/InsightDC/extract_hashes_serial Hashes
            cat Hashes
            # rm Hashes
            # touch pub_Hashes
        fi
    done
