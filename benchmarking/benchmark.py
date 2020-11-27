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
    subprocess.run("echo jackrecacher > ../sample.log", shell=True)
    for th in thread_range:
        for conn in connection_range:
            for tr in time_range:
                for exp in range(number_of_experiments):
                    # command = 'wrk -t{} -c{} -d{}s -H "X-Vault-Token: {}" -s write-random-secrets.lua {} -- 10000 > /home/logfiles/prod-test-write-1000-random-secrets-t{}-c{}-{}sec-exp{}.log'.format(th,conn,tr,vault_token,vault_address ,th,conn,tr,exp)
                    
                    # time.sleep(10)
                    pass