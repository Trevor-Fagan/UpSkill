import threading
import time
import os

def test(thread_number):
    time.sleep(1)
    print(thread_number)

def main():
    start_time = time.time()
    t1 = threading.Thread(target=test, args=(1,))
    t2 = threading.Thread(target=test, args=(2,))

    t1.start()
    t2.start()

    t1.join()
    t2.join()

    end_time = time.time()
    print("Done in " + str(end_time - start_time))

if __name__ == "__main__":
    main()