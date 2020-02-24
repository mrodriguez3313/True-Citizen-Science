#!/bin/bash
echo "kill python" && kill $(pgrep "python")
echo "kill autoUpload" && kill $(pgrep "autoUpload.sh")
echo "kill autoDownload" && kill $(pgrep "autoDownload.sh")
echo "kill inotifywait" && kill $(pgrep "inotifywait")
echo "kill ipfs pubsub sub Hashes" && kill $(pgrep "ipfs pubsub sub Hashes")
