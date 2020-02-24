import threading
import time
import random
import os
from signal import *
import sys

file_categories = ["Beaver", "Mice", "Beach", "Mountains"]

def create_files():
    # each thread will send a new output to a new file
    files_to_create = 20
    for idx in range(files_to_create):
        with open("/home/marco/Documents/InsightDC/filesdump/hello_world{}.txt".format(int(round(time.time() * 10000))), "w") as f:
            f.write((''.join('{}; {}, {}'.format(random.choice(file_categories), threading.current_thread().ident, idx))))
        time.sleep(10)
    return

def join_threads(threads):
    for index, thread in enumerate(threads):
        thread.join()

def main():
    threads = list()
    path="/home/marco/Documents/InsightDC/filesdump/"
    isdir = os.path.isdir(path)
    if not isdir:
        try:
            os.mkdir(path)
        except OSError:
            print("Directory at %s failed to be created." % path)
            return
        else:
            print("Directory at %s created." % path)

    for num_threads in range(1):
        x = threading.Thread(target=create_files)
        threads.append(x)
        time.sleep(1)
        x.start()


    join_threads(threads)
main()
