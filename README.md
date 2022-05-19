# httpdumper

- minimum http server
- HTTP Request data dump to stdout and client


# How to install and using for Linux

1. confirm artifacts URL at [Releases](https://github.com/noid11/httpdumper/releases)
    - i.g.: [httpdumper_0.1.2_Linux_x86_64.tar.gz](https://github.com/noid11/httpdumper/releases/download/v0.1.2/httpdumper_0.1.2_Linux_x86_64.tar.gz)
2. download artifact
    - i.g.: `wget https://github.com/noid11/httpdumper/releases/download/v0.1.2/httpdumper_0.1.2_Linux_x86_64.tar.gz`
3. extract binary
    - i.g.: `tar xvf httpdumper_0.1.2_Linux_x86_64.tar.gz`
4. execute httpdumper
    - i.g.: `sudo ./httpdumper`
5. `Ctrl + C` to exit


# Output example

Server Side

```bash
$ sudo ./httpdumper 
2022/05/19 12:24:38 GET / HTTP/1.1
Host: 18.179.80.250
Accept: */*
User-Agent: curl/7.79.1

^C
```

Client Side

```bash
$ curl [server host]
GET / HTTP/1.1
Host: 18.179.80.250
Accept: */*
User-Agent: curl/7.79.1

```

