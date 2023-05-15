# Description
Testing the ability of ChatGPT to generate Go code that can discover SSDP devices.

Prints out a JSON like:-

`{"location":["http://192.168.29.21:9999/ssdp/device-desc.xml","http://192.168.29.21:8888/ssdp/device-desc.xml"],"scheme":["http","http"],"host":["192.168.29.21","192.168.29.21"],"port":["9999","8888"]}`

# Building
You can build a binary out of it using `go build`.

### Termux Users
Only for those using Termux in Android device

Set up the build tools using -

    pkg upgrade -y

&nbsp;

    pkg install -y git golang

### Build commands
The general build commands -

    cd
    rm -rf tasker-project-assets >/dev/null 2>&1
    git clone -n --depth=1 --filter=tree:0 https://github.com/HunterXProgrammer/tasker-project-assets
    cd tasker-project-assets
    git sparse-checkout set --no-cone discover
    git checkout
    cd discover
    go build -o discover -buildvcs=false -ldflags="-extldflags -s" discover.go

Or one-liner -

    cd; rm -rf tasker-project-assets >/dev/null 2>&1; git clone -n --depth=1 --filter=tree:0 https://github.com/HunterXProgrammer/tasker-project-assets; cd tasker-project-assets; git sparse-checkout set --no-cone discover; git checkout; cd discover; go build -buildvcs=false -ldflags="-extldflags -s" discover.go

### Run The Binary
After building, run the binary using -

    ./discover ssdp:all

If SSDP devices found then a JSON is printed -

`{"location":["http://192.168.29.21:9999/ssdp/device-desc.xml","http://192.168.29.21:8888/ssdp/device-desc.xml"],"scheme":["http","http"],"host":["192.168.29.21","192.168.29.21"],"port":["9999","8888"]}`

