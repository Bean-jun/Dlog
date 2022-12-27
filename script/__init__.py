import contextlib
import datetime
import os
import re
import shutil
import subprocess


def run_cmd(shell):
    vals = subprocess.Popen(args=shell,
                            stdout=subprocess.PIPE).communicate()[0].decode()
    return vals.strip("\n").strip(" ")


def wirte_file(fb, content, mode="w", encoding="utf-8"):
    with open(fb, mode, encoding=encoding) as f:
        f.write(content)


def read_file(fb, mode="r", encoding="utf-8"):
    with open(fb, mode, encoding=encoding) as f:
        return f.read()


@contextlib.contextmanager
def to_path(path):
    originPath = os.getcwd()
    os.chdir(path)
    try:
        yield originPath
    finally:
        os.chdir(originPath)


def __publish_version(version):
    compile = re.compile("""Version = \"(.+)\"""")
    Version = """Version = \"%s\"""" % version
    Goversion = read_file("pkg/version.go")

    find = compile.findall(Goversion)
    if len(find):
        Goversion = compile.sub(Version, Goversion)
    else:
        Goversion = """package pkg

const (
	Version = "%s"
)\n""" % version

    wirte_file("pkg/version.go", Goversion)


def publish_version():
    """
    发布版本
    """
    version_date = datetime.datetime.now().strftime("%Y.%m.%d.")
    version_hash = run_cmd("git log -1 --format=%h")
    version = version_date + version_hash

    __publish_version(version)

    run_cmd("git add .")
    run_cmd("git commit -m 发布版本-V%s" % version)
    run_cmd("git tag %s -m '发布版本-V%s'" % (version, version))


def build():
    with to_path("src") as currentPath:
        run_cmd("make build-linux")
        targetPath = os.path.join(currentPath, "Release")
        shutil.rmtree(targetPath)
        shutil.copytree("Release", targetPath)
        shutil.rmtree("Release")
    print("build over...")
