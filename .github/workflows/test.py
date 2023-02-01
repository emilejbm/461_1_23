import subprocess

print(subprocess.run(["run.sh", "install"], shell=True, stdout=subprocess.PIPE).stdout.decode('utf-8'))