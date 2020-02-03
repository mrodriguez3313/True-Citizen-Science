import threading
import time

def create_files():
    # each thread will send a new output to a new file
    file_num = 0
    while True:
        with open("/home/marco/Documents/InsightDC/filesdump/hello_world{}.txt".format(int(round(time.time() * 10000))), "w") as f:
            f.write((''.join('{}, {}'.format(threading.current_thread().ident, file_num))))
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
