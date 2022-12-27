import sys
from ..script import *


if __name__ == "__main__":
    shell = sys.argv[1:]
    if shell[0] == "publish":
        publish_version()
