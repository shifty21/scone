import yaml
import time
import subprocess
def load_config(configfile):
    print("loading config from %s" % configfile)
    with open(configfile, "r") as ymlfile:
        return yaml.load(ymlfile, Loader=yaml.FullLoader)

if __name__=="__main__":
    config = load_config("resource.yaml")
    number_of_experiments = config["number_of_experiments"]
    thread_range = config["thread_range"]
    connection_range = config["connection_range"]
    time_range = config["time_range"]
    vault_address = config["vault_address"]
    vault_token = config["vault_token"]
    # log_files = config["log_files"]
    cmd = config["cmd_list"]
    print(cmd)
    for x in cmd:
        print("Running for %s with result in %s" % (x["cmd"], x["dir"]))
        for th in thread_range:
            for conn in connection_range:
                for tr in time_range:
                    for exp in range(number_of_experiments):
                        command = x["cmd"].format(th,conn,tr,vault_token,vault_address,x["dir"] ,th,conn,tr,exp)
                        subprocess.run(command, shell=True)
                        time.sleep(10)
        print("\n==============================\n")