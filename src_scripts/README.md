This directory contains 5 scripts:
autoDownload.sh
autoUpload.sh
file_generator.py
kill_scripts.sh
script.sh
demo.sh

autoDownload.sh (run `./autoDownload.sh &`)
	This program is meant to run in the backgroud of the researcher's computer to listen for data that is being published over the IPFS pubsub network under a specific topic and call extract_hashes on the new found content. This script listens for data on the 'Hashes' topic by default. This is the file for the inotifywait package to monitor when something gets published under that topic. 

autoUpload.sh (run `./autoUpload.sh &`)
	This file will upload any file that is in the directory called '*filesdump*'. If it is not there then it will not monitor the directory. This file requires a package be installed called 'inotifywait' that is the core of this automation process that says: if a file is created in said direcory, run this command.
	In our case we are running 'ipfs add'. Which implies that you must have ipfs daemon installed and running as well to run this script. 

file_generator.py
	This is a python script used to simulate logs/files being contributed by multiple sources and being automatically stored in 'filesdump'. If the directory does not exist, an error will be thrown. Currently this is a single threaded process, but you can create more threads and by changing the variable 'num_threads' you can also change how many files get created by changing the variable 'files_to_create' 

kill_scripts.sh (run `./kill_scripts.sh`)
	This script will kill the file_generator.py, autoUpload.sh, autoDownload.sh, python scripts, you might need to kill inotifywait and 'ipfs pubsub sub Hashes' manually. All these mentioned programs are meant to be run in the background. To run this script you need to have the package 'pgrep' installed.

script.sh (`run ./script &`)
	This is a script that called file_generator.py, but runs it in the background.

demo.sh (run `./demo.sh`)
	This script is meant to automate the environment to autocreate and upload files into ipfs. The container(s) should have been preset up to have go & ipfs installed in it as well as preconfigured to be connected to a private network (easy to do with init_network.sh). Then the script will run autoUpload and script.sh in the background so to free up your current terminal; print number of files created; after 20 seconds it will call main.go and then open the mongodb shell'. From there you can make queries to see the things added in the db called PrivateIPFSDB, in the collection AllProjects, and/or quit by typing exit or Ctrl+C.
