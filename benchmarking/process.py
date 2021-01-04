import os
import re
import glob
import argparse

def wrk_data(wrk_output):
    return str(wrk_output.get('file_name')) + ',' + str(wrk_output.get('lat_avg')) + ',' + str(wrk_output.get('lat_stdev')) + ',' + str(
        wrk_output.get('lat_max')) + ',' + str(wrk_output.get('lat_stdevpm')) + ',' + str(
        wrk_output.get('req_avg')) + ',' + str(wrk_output.get('req_stdev')) + ',' + str(
        wrk_output.get('req_max')) + ','  + str(wrk_output.get('req_stdevpm')) + ',' + str(
        wrk_output.get('tot_requests')) + ',' + str(wrk_output.get('tot_duration')) + ',' + str(
        wrk_output.get('read')) + ',' + str(wrk_output.get('err_connect')) + ',' + str(
        wrk_output.get('err_read')) + ',' + str(wrk_output.get('err_write')) + ',' + str(
        wrk_output.get('err_timeout')) + ',' + str(wrk_output.get('req_sec_tot')) + ',' + str(
        wrk_output.get('read_tot')) + ',' + str(wrk_output.get('threads')) + ',' + str(
        wrk_output.get('total_requests'))  + ',' + str(wrk_output.get('total_responses')) + '\n'


def get_bytes(size_str):
    x = re.search("^(\d+\.*\d*)(\w+)$", size_str)
    if x is not None:
        size = float(x.group(1))
        suffix = (x.group(2)).lower()
    else:
        return size_str

    if suffix == 'b':
        return size
    elif suffix == 'kb' or suffix == 'kib':
        return size * 1024
    elif suffix == 'mb' or suffix == 'mib':
        return size * 1024 ** 2
    elif suffix == 'gb' or suffix == 'gib':
        return size * 1024 ** 3
    elif suffix == 'tb' or suffix == 'tib':
        return size * 1024 ** 3
    elif suffix == 'pb' or suffix == 'pib':
        return size * 1024 ** 4

    return False


def get_number(number_str):
    x = re.search("^(\d+\.*\d*)(\w*)$", number_str)
    if x is not None:
        size = float(x.group(1))
        suffix = (x.group(2)).lower()
    else:
        return number_str

    if suffix == 'k':
        return size * 1000
    elif suffix == 'm':
        return size * 1000 ** 2
    elif suffix == 'g':
        return size * 1000 ** 3
    elif suffix == 't':
        return size * 1000 ** 4
    elif suffix == 'p':
        return size * 1000 ** 5
    else:
        return size

    return False


def get_ms(time_str):
    x = re.search("^(\d+\.*\d*)(\w*)$", time_str)
    if x is not None:
        size = float(x.group(1))
        suffix = (x.group(2)).lower()
    else:
        return time_str

    if suffix == 'us':
        return size / 1000
    elif suffix == 'ms':
        return size
    elif suffix == 's':
        return size * 1000
    elif suffix == 'm':
        return size * 1000 * 60
    elif suffix == 'h':
        return size * 1000 * 60 * 60
    else:
        return size

def parse_wrk_output(file, wrk_output):
    retval = {}
    retval['file_name'] = file
    requests = 0
    responses = 0
    thread = 0
    for line in wrk_output.splitlines():
        x = re.search("^\s+Latency\s+(\d+\.\d+\w*)\s+(\d+\.\d+\w*)\s+(\d+\.\d+\w*)\s+(\d+\.\d+\w*).*$", line)
        if x is not None:
            retval['lat_avg'] = get_ms(x.group(1))
            retval['lat_stdev'] = get_ms(x.group(2))
            retval['lat_max'] = get_ms(x.group(3))
            retval['lat_stdevpm'] = get_number(x.group(4))
        x = re.search("^\s+Req/Sec\s+(\d+\.\d+\w*)\s+(\d+\.\d+\w*)\s+(\d+\.\d+\w*)\s+(\d+\.\d+\w*).*$", line)
        if x is not None:
            retval['req_avg'] = get_number(x.group(1))
            retval['req_stdev'] = get_number(x.group(2))
            retval['req_max'] = get_number(x.group(3))
            retval['req_stdevpm'] = get_number(x.group(4))
        x = re.search("^\s+(\d+)\ requests in (\d+\.\d+\w*)\,\ (\d+\.\d+\w*)\ read.*$", line)
        if x is not None:
            retval['tot_requests'] = get_number(x.group(1))
            retval['tot_duration'] = get_ms(x.group(2))
            retval['read'] = get_bytes(x.group(3))
        x = re.search("^Requests\/sec\:\s+(\d+\.*\d*).*$", line)
        if x is not None:
            retval['req_sec_tot'] = get_number(x.group(1))
        x = re.search("^Transfer\/sec\:\s+(\d+\.*\d*\w+).*$", line)
        if x is not None:
            retval['read_tot'] = get_bytes(x.group(1))
        x = re.search(
            "^\s+Socket errors:\ connect (\d+\w*)\,\ read (\d+\w*)\,\ write\ (\d+\w*)\,\ timeout\ (\d+\w*).*$", line)
        if x is not None:
            retval['err_connect'] = get_number(x.group(1))
            retval['err_read'] = get_number(x.group(2))
            retval['err_write'] = get_number(x.group(3))
            retval['err_timeout'] = get_number(x.group(4))
        x = re.search(
            "^thread (\d+) made (\d+) requests including (\d+) writes and got (\d+) responses.*$", line)
        if x is not None:
            thread = get_number(x.group(1))
            requests = requests +  get_number(x.group(2))
            responses =  responses + get_number(x.group(4))
    if 'err_connect' not in retval:
        retval['err_connect'] = 0
    if 'err_read' not in retval:
        retval['err_read'] = 0
    if 'err_write' not in retval:
        retval['err_write'] = 0
    if 'err_timeout' not in retval:
        retval['err_timeout'] = 0
    retval['threads'] = thread
    retval['total_requests'] = requests
    retval['total_responses'] = responses
    return retval


def get_per_file_data(filename):
    file = open(filename)
    wrk_output =  file.read() 
    wrk_output_dict = parse_wrk_output(file.name.split("/")[-1], wrk_output)
    # print(str(wrk_output_dict) + "\n\n")
    # print("****wrk output csv line: \n\n")
    wrk_output_csv = wrk_data(wrk_output_dict)
    return wrk_output_csv


def process_files(dir_name,path):
    list_of_files = sorted(glob.glob(path+"*.log"))
    header = 'lat_avg,lat_stdev,lat_max,lat_stdevpm,req_avg,req_stdev,req_max,req_stdevpm,'\
    'tot_requests,tot_duration,read,err_connect,err_read,err_write,err_timeout,req_sec_tot'\
    ',read_tot,threads,total_requestss,total_responses\n'
    result = []
    result.append(header)
    for file in list_of_files:
        data = get_per_file_data(file)
        result.append(data)
    with open(dir_name+".csv",'w') as resultcsv:
        resultcsv.writelines(result)

def process_dir(dirs):
    
    list_of_dirs = glob.glob(dirs+"*/")
    
    print("Processing directories in path %s with directories %s"%(dirs, list_of_dirs))
    for dir in list_of_dirs:
        print("Processing directory %s" % str(dir))
        process_files(str(dir).split("/")[-2],dir)


if __name__ == '__main__':
    ap = argparse.ArgumentParser()
    ap.add_argument('-p', '--filedir',
                    default='./', help='config file path')
    ap.add_argument('-d', '--dir',
                    default='./', help='config file path')
    args = ap.parse_args()
    
    path = args.dir
    process_dir(path)
