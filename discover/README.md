# Description
Testing the ability of ChatGPT to generate Go code that can discover SSDP devices.

Prints out a JSON like:-

`{"location":["http://192.168.29.21:8888/ssdp/device-desc.xml","http://192.168.29.21:9999/ssdp/device-desc.xml"],"scheme":["http","http"],"host":["192.168.29.21","192.168.29.21"],"port":["8888","9999"]}`

# Building
You can build a binary out of it using `go build`.

Or you can grab the pre-compiled binaries in **[releases](https://github.com/HunterXProgrammer/tasker-project-assets/releases)**.

### Termux Users
Only for those using Termux in Android device.

Set up the build tools using -

    pkg upgrade -y

&nbsp;

    pkg install -y git golang termux-elf-cleaner

### Build commands
The general build commands -

    curl -s -L "https://github.com/HunterXProgrammer/tasker-project-assets/releases/download/discover/build_discover.sh" | bash

### Run The Binary
After building, run the binary using -

    ~/discover && ./discover ssdp:all

If SSDP devices found then a JSON is printed -

`{"location":["http://192.168.29.21:8888/ssdp/device-desc.xml","http://192.168.29.21:9999/ssdp/device-desc.xml"],"scheme":["http","http"],"host":["192.168.29.21","192.168.29.21"],"port":["8888","9999"]}`

