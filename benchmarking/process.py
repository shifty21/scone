import os
import re
import glob
import argparse

def wrk_data(wrk_output):
    if (int(wrk_output.get('run_time')) != 20):
        return ""
    return str(wrk_output.get('run_time')) + ',' + str(wrk_output.get('lat_avg')) + ',' + str(wrk_output.get('lat_stdev')) + ',' + str(
        wrk_output.get('lat_max')) + ',' + str(wrk_output.get('lat_stdevpm')) + ',' + str(
        wrk_output.get('req_avg')) + ',' + str(wrk_output.get('req_stdev')) + ',' + str(
        wrk_output.get('req_max')) + ','  + str(wrk_output.get('req_stdevpm')) + ','  + str(
        wrk_output.get('read')) + ',' + str(wrk_output.get('err_connect')) + ',' + str(
        wrk_output.get('err_read')) + ',' + str(wrk_output.get('err_write')) + ',' + str(
        wrk_output.get('err_timeout')) + ',' + str(wrk_output.get('req_sec_tot')) + ',' + str(
        int(wrk_output.get('read_tot'))/1000000) + ',' + str(wrk_output.get('threads')) + ',' + str(
        wrk_output.get('connections'))+ ',' + str(wrk_output.get('total_requests'))  + ',' + str(
        wrk_output.get('total_responses')) + ',' + str(wrk_output.get('min_lat'))+ ',' + str(
        wrk_output.get('50per'))  + ',' + str(wrk_output.get('90per'))+ ',' + str(
        wrk_output.get('99per'))+ ',' + str(wrk_output.get('9999per'))+ ',' + str(wrk_output.get('throughput')) + '\n'

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
    x = re.search("-(\d+)sec-", file)
    if x is not None:
        retval['run_time'] = x.group(1)
    requests = 0
    responses = 0
    thread = 0
    connections = 0
    for line in wrk_output.splitlines():
        # print("Processing line %s" %(line))
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
            retval['total_requests'] = get_number(x.group(1))
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
        # x = re.search(
        #     "^thread (\d+) made (\d+) requests including (\d+) writes and got (\d+) responses.*$", line)
        # if x is not None:
        #     # thread = thread + get_number(x.group(1))
        #     requests = requests +  get_number(x.group(2))
        #     responses =  responses + get_number(x.group(4))
        # x = re.search(
        #     "^thread (\d+) made (\d+) requests including (\d+) lists and got (\d+) responses.*$", line)
        # if x is not None:
        #     # thread = thread + get_number(x.group(1))
        #     requests = requests +  get_number(x.group(2))
        #     responses =  responses + get_number(x.group(4))
        x = re.search("^\s\s(\d+) threads and (\d+) connections.*$", line)
        if x is not None:
            thread = thread + get_number(x.group(1))
            connections = connections + get_number(x.group(2))
        x = re.search("^min_lat: (\d+)[.]\d+\,max_lat: (\d+)[.]\d+\,mean_lat: (\d+)[.]\d+\,stdev_lat: (\d+)[.]\d+,50per: (\d+)[.]\d+,90per: (\d+)[.]\d+,99per: (\d+)[.]\d+,9999per: (\d+)[.]\d+,dur: \d+,req: (\d+),byte: (\d+),econn: \d+,eread: \d+,ewrite: \d+,estatus: \d+,etout: \d+,resp: (\d+),writes: (\d+).*$",line)
        if x is not None:
            retval['min_lat'] = get_ms(x.group(1)+"us")
            retval['50per'] = get_ms(x.group(5)+"us")
            retval['90per'] = get_ms(x.group(6)+"us")
            retval['99per'] = get_ms(x.group(7)+"us")
            retval['9999per'] = get_ms(x.group(8)+"us")
            responses =  responses + get_number(x.group(11))
        x = re.search("^Requests\/sec:[ ]{2,}(\d+).(\d+)$", line)
        if x is not None:
            retval['throughput'] = x.group(1)
    if 'err_connect' not in retval:
        retval['err_connect'] = 0
    if 'err_read' not in retval:
        retval['err_read'] = 0
    if 'err_write' not in retval:
        retval['err_write'] = 0
    if 'err_timeout' not in retval:
        retval['err_timeout'] = 0
    retval['threads'] = thread
    retval['connections'] = connections
    retval['total_responses'] = responses
    # retval['throughput'] = retval['total_requests']/retval['run_time']
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
    print("Processing dir_name %s, path %s"%(dir_name,path))
    list_of_files = sorted(glob.glob(path+"*.log"))
    header = 'run_time,lat_avg,lat_stdev,lat_max,lat_stdevpm,req_avg,req_stdev,req_max,req_stdevpm,'\
    'read,err_connect,err_read,err_write,err_timeout,req_sec_tot'\
    ',read_tot,threads,connections,total_requests,total_responses,min_lat,50per,90per,99per,9999per,throughput\n'
    result = []
    result.append(header)
    for file in list_of_files:
        data = get_per_file_data(file)
        if data.__contains__("None") or data == "":
            continue
        result.append(data)
    with open(dir_name+".csv",'w') as resultcsv:
        resultcsv.writelines(result)

def process_dir(dirs):
    
    list_of_dirs = glob.glob(dirs+"*/")
    for dir in list_of_dirs:
        # print("Processing directory %s" % str(dir))
        if "csv" in dir:
            print("Skipping CSV dir")
            continue
        process_files(str(dir).split("/")[-2],dir)
    if len(list_of_dirs) == 0:
        process_files("sample1.log",str("./"))


if __name__ == '__main__':
    ap = argparse.ArgumentParser()
    ap.add_argument('-p', '--filedir',
                    default='./', help='config file path')
    ap.add_argument('-d', '--dir',
                    default='./', help='config file path')
    args = ap.parse_args()
    
    path = args.dir
    process_dir(path)
