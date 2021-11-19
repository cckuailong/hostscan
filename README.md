# hostscan

[中文Readme](./README_zh.md)

## What is hostscan

**Auto** Host Collsion Tool, In order to help RedTeam quickly expand the network boundary and access more target points

## Why hostscan ?? 

In many cases, when accessing the target website, it cannot be accessed using its real IP, and only the domain name can access the back-end business services. This is because the Reverse proxy server (such as nginx) is configured to prohibit direct IP access.

The business behind nginx is sometimes hidden:
1. Services that are not open to external networks, such as test services
2. The original DNS was resolved to the external network, but the A record was later deleted (the business after nginx was not deleted and transferred to the internal network access)

How to access these hidden businesses? This requires the appearance of today's protagonist-Host collision technology

## Example

```
./hostscan -d 127.0.0.1 -i 1.1.1.1
```

```
./hostscan -D input/hosts.txt -I input/ips.txt -O out/output.txt -T 5 -t 10
```

## Usage

Please download the version of the corresponding platform in the release

```
./hostscan --help
  
/ )( \ /  \ / ___)(_  _)/ ___) / __) / _\ (  ( \
) __ ((  O )\___ \  )(  \___ \( (__ /    \/    /
\_)(_/ \__/ (____/ (__) (____/ \___)\_/\_/\_)__)        
Usage of ./main:
  -D string
        Hosts in file to test
  -I string
        Nginx Ip in file to test
  -O string
        Output File (default "result.txt")
  -T int
        Thread for Http connection. (default 3)
  -d string
        Host to test
  -i string
        Nginx IP
  -t int
        Timeout for Http connection. (default 5)
  -v    Show hostscan version

```

## Demo

![demo](./images/demo.png)

## References

[Fofapro 的 Hosts_scan](https://github.com/fofapro/Hosts_scan)
