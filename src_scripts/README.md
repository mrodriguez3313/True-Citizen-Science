This directory contains 5 scripts:
autoUpload.sh
file_generator.py
kill_scripts.sh
script.sh
demo.sh

autoUpload.sh
	This file will upload any file that is in the previous directory called *filesdump*. If it is not there then it will not throw an error, it just won't do anything. This file requires a package be installed called 'inotifywait' that is the core of this automation process that says: if a file is created in said direcory, run this command.
	In our case we are running 'ipfs add'. Which implies that you must have ipfs daemon isntalled and running aswell to run this script.

file_generator.py
	This is a python script used to simulate logs/files being contributed by multiple sources and being automatically stored in 'filesdump'. If the directory does not exist, an error will be thrown. Currently this is a single threaded process, but you can create more threads and by changing the variable 'num_threads' you can also change how many files get created by changing the variable 'files_to_create'

kill_scripts.sh
	This script will kill the file_generator.py and autoUpload.sh. As they are meant to be run in the background. To run this script you need to have the package 'pgrep' installed.

script.sh
	This is a script that called file_generator.py, but runs it in the background.

demo.sh
	This script is meant to automate the environment to autocreate and upload files into ipfs. It first starts a docker container called ipfsnode. This container should have been preset up to have go & ipfs installed in it as well as preconfigured to be connected to a private network. Then the script will run autoUpload and script.sh in the background so to free up your current terminal. after 20 seconds it will kill the processes with './kill_scrips.sh'
