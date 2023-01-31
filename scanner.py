# Need to set up virtual environment and install Gitpython
# GitPython requires Python v3.7, we are using v3.6.8

from git import Repo
import os
import shutil

words = ["Ben Schwartz", "Commands", "PQWERETS", "License"]

def func(url: str):
    dir = "py_workspace"
    os.mkdir(dir)
    Repo.clone_from(url + ".git", dir)

    # Search for README
    found = 0

    files = os.listdir(dir)
    if "README.md" in files:
        with open(dir + "/README.md", "r") as f:
            txt = f.read()
            for word in words:
                if word in txt:
                    print(word)
                    found += 1

    shutil.rmtree(dir)
    return found / len(words)
    






a = func("https://github.com/benschwartz9/461_1_23")
print(a)

b = func("https://github.com/axios/axios")
print(b)