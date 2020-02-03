Marco Rodriguez 
Project: Creating a platform for dectralized citizen science

extract_hashes:
	This file assumes that you have a file with hashes called 'Hashes'. The hashes are id's for files in ipfs. This script will read the contents of those files and create a file called 'hash_outputs' where each line is the contents of each one of those hashes in 'Hashes'. 'hash_outputs' is used in /src/main.go to take those outputs and add them into mongodb. 

*Currently need to make it so that as hashes get appended to Hashes this file will read the contents and append them to 'hash_outputs', so that it can also add the files into the mongodb through main.go*
