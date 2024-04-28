import os


def getChangedFiles():
    os.system("git remote add build git@github.com:lxhandsome/leetcode.git")
    os.system("git fetch build master")
    files = os.popen("git diff HEAD  build/master  --name-only").readlines()
    os.system("git remote remove build")
    return files


def getGoChangedDirs():
    set = {}
    dirs = []
    files = getChangedFiles()
    for file in files:
        fn = file.strip()
        if not fn.endswith("go"):
            continue
        dir = os.path.dirname(fn)
        set[dir] = ""
    for dir in set:
        dirs.append("./"+dir)
    return dirs


def genTestHtml(dirs):
    pkgCmd = " ".join(dirs)
    print(pkgCmd)
    tmp = "tmp"
    os.makedirs(tmp, exist_ok=True)
    outFile = os.path.join(tmp, "cover.out")
    htmlFile = os.path.join(tmp, "cover.html")
    testCmd = f"go test {pkgCmd} -coverprofile {outFile} -cover count"
    coverCmd = f"go tool cover -html {outFile} -o {htmlFile}"
    print(testCmd)
    print(coverCmd)
    os.system(testCmd)
    os.system(coverCmd)


if __name__ == "__main__":
    dirs = getGoChangedDirs()
    genTestHtml(dirs)
