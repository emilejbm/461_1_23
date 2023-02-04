from git import Repo
import os
import sys
import tempfile

# Words to scan for
words = ["Getting Started", "Test", "Help", "Contributing", "License"]

def scan(url: str):
    # Scan README.md for keywords
    with tempfile.TemporaryDirectory() as dir:
        Repo.clone_from(url + ".git", dir)

        # Search for README
        found = 0
        files = os.listdir(dir)
        if "README.md" in files:
            with open(dir + "/README.md", "r") as f:
                txt = f.read().lower()
                for word in words:
                    if word.lower() in txt:
                        found += 1
        
        return found / len(words)

# Start
if len(sys.argv) != 2:
    sys.exit("Incorrect number of arguments")

url = sys.argv[1]
rampUpScore = scan(url)
print(rampUpScore)
