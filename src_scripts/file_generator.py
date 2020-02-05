import threading
import time
import random

file_categories = ["Beaver", "Mice", "Beach", "Mountains"]

def create_files():
    # each thread will send a new output to a new file
    files_to_create = 10
    for idx in range(files_to_create):
        with open("/home/marco/Documents/InsightDC/filesdump/hello_world{}.txt".format(int(round(time.time() * 10000))), "w") as f:
            f.write((''.join('{}; {}, {}'.format(random.choice(file_categories), threading.current_thread().ident, idx))))
        # print("file written by thread {}.".format(threading.current_thread().ident))
        time.sleep(5)
    return

def main():
    threads = list()
    for num_threads in range(1):
        x = threading.Thread(target=create_files)
        threads.append(x)
        x.start()

    for index, thread in enumerate(threads):
        thread.join()


main()
