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
./hostscan -d test.com -i 127.0.0.1:3333
```

```
./hostscan -D input/hosts.txt -I input/ips.txt -O out/output.txt -T 5 -t 10 -U
```

## Usage

Please download the version of the corresponding platform in the release

*Notice:*
- Default thread only set to 3, if the network is ok, thread can be set up to rlimit.
- Default UserAgent use `golang-hostscan/xxxx`, if you want to use random UA, please add param '-U'.
- Support the large input file, Now there is no worry about OOM.

```
hostscan --help
  
/ )( \ /  \ / ___)(_  _)/ ___) / __) / _\ (  ( \
) __ ((  O )\___ \  )(  \___ \( (__ /    \/    /
\_)(_/ \__/ (____/ (__) (____/ \___)\_/\_/\_)__)        
Usage of hostscan:
  -D string
        Hosts in file to test
  -F string
        Filter result with List of Response Status Code. 
        Example: 200,201,302
  -I string
        Nginx Ip in file to test
  -O string
        Output File (default "result.txt")
  -T int
        Thread for Http connection. (default 3)
  -U    Open to send random UserAgent to avoid bot detection.
  -V    Output All scan Info. 
        Default is false, only output the result with title.
  -d string
        Host to test
  -i string
        Nginx IP. 
        Example: 1.1.1.1 or 1.2.3.4/24
  -p string
        Port List of Nginx IP. If the flag is set, hostscan will ignore the port in origin IP input. 
        Example: 80,8080,8000-8009
  -t int
        Timeout for Http connection. (default 5)
  -v    Show hostscan version

```

## Demo

*Test the vultarget below*

Host Collsion Success

![demo](./images/demo1.png)

Get status 400

![demo](./images/demo2.png)

## Test Vultarget

### Docker

```
docker pull vultarget/host_collision
docker run -it -p 3333:8080 --rm vultarget/host_collision
```

### Nginx Configuration

#### Reverse proxy server (Core)

```
server {
    listen  8080  default_server;
    server_name _;
    return 400;
}
server {
    listen  8080;
    server_name test.com;


    location / {
        proxy_pass http://127.0.0.1:80;
        proxy_redirect off;
        proxy_set_header Host $host:$server_port;
        proxy_set_header X-Real-IP $remote_addr;
            root    html;
        index   index.html  index.htm;
    }
    access_log logs/test.com.log;
}
```

The first server indicates that, when the host is empty, it will return 400 status

The second server indicates that nginx will forward the service according to the incoming host, and the business accessed by test.com is the service on 127.0.0.1:80

#### Example Web

```
server {
    listen       80;
    server_name  localhost;


    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }


    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

Simple Nginx Web Page.

## ChangeLog

v0.2.3
- Fix the bug of wrong calculation of file line number

v0.2.2
- The -i option supports IP range scanning, such as 1.2.3.4/24
- The -p option supports custom scan ports, such as 80,8000-8009
- The -V option outputs all scan information, disabled by default, only outputting results
- The -F option help you to filter the result with http status code
- Fixed a bug where the progress bar would still be displayed when no parameters were given
- Added some informative output

## References

[Fofapro's Hosts_scan](https://github.com/fofapro/Hosts_scan)
