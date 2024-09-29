import multiprocessing
import time
import os
import math

def test(result_list, process_number):
    result_list[process_number] = process_number
    
    for result in result_list:
        print(result)

def main():
    num_processes = 4
    result = multiprocessing.Array('i', num_processes)

    processes = []
    for i in range(num_processes):
        temp_process_number = multiprocessing.Value('i')
        temp_process_number.val = i
        temp = multiprocessing.Process(target=test, args=(result,temp_process_number.val))
        temp.start()
        processes.append(temp)

    for process in processes:
        process.join()

if __name__ == "__main__":
    main()