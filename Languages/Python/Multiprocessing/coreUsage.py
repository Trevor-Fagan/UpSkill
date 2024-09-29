import multiprocessing
import time
import os
import math

'''
Used to visualize multi processing. The intent is to view the time it takes to perform
these actions based on the number of processes that we schedule to run. On my hardware,
sum(range(10**6)) 50 times takes ~1 second. This is running on an 8 core machine. We can
see that up to those 8 cores, it takes 1 second as the processes are running independently
on each core. Double it, it takes 2 seconds, double again and it takes 4 seconds, etc. Can
also visualize CPU core usage with htop and see all cores max out when running.
'''

def test(process_number):
    for i in range(50):
        sum(range(10**6))  # Adjust the range for more complexity if needed

    print(process_number)

def main():
    start_time = time.time()

    processes = []
    for i in range(64):
        temp = multiprocessing.Process(target=test, args=(i+1,))
        temp.start()
        processes.append(temp)

    for process in processes:
        process.join()

    end_time = time.time()
    print("Done in " + str(end_time - start_time))

if __name__ == "__main__":
    main()