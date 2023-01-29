import subprocess

# print(subprocess.run(["run.sh", "install"], shell=True, stdout=subprocess.PIPE).stdout.decode('utf-8'))
# print(subprocess.check_output(['run.sh', "install"], shell=True))
data = subprocess.run('./run.sh install', capture_output=True, shell=True)
output = data.stdout
print(output)
# print(subprocess.run([r"C:\Users\mmcho\OneDrive - purdue.edu\ECE 461\Part 1\461_1_23\run.sh", "build"], shell=True))
# print(subprocess.run([r"C:\Users\mmcho\OneDrive - purdue.edu\ECE 461\Part 1\461_1_23\run.sh", "URL_FILE"], shell=True))
# print(subprocess.run([r"C:\Users\mmcho\OneDrive - purdue.edu\ECE 461\Part 1\461_1_23\run.sh", "test"], shell=True))

