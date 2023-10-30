## scan_ip (ip 扫描程序)

#### 使用方式

#### 场景1

判断单个 ip

命令示例：scan_ip.exe 127.0.0.1

如果需要保存扫描结果：scan_ip.exe 127.0.0.1 s

结果文件名称：scan_succ_ip.csv
#### 场景2
判断多个端的 ip

命令示例：scan_ip.exe 127.0.*.1

符合 * 的范围是 0-255

如果需要保存扫描结果：scan_ip.exe 127.0.*.1 s

结果文件名称：scan_succ_ip.csv

#### 下载
linux 端下载 ：scan_ip

windows 端下载：scan_ip.exe

#### 源码

windows 打包

```shell
go build -o scan_ip.exe scan_ip/scan_ip.go
```

mac 打包
```shell
set GOOS=darwin
set GOARCH=amd64
go build -o scan_ip_mac scan_ip/scan_ip.go
```

linux 打包
```shell
set GOARCH=amd64
set GOOS=linux
go build -o scan_ip_linux scan_ip/scan_ip.go
```


## scan_ip_port (ip:port 扫描程序)

#### 使用方式

#### 场景1

判断单个 ip:port

命令示例：scan_ip_port.exe 127.0.0.1 80

如果需要保存扫描结果：scan_ip_port.exe 127.0.0.1 80 s

结果文件名称：scan_succ_ip_port.csv
#### 场景2
扫描 ip port 范围

命令示例：scan_ip_port.exe 127.0.0.1 3000 4000

port 的范围是 0-65535

如果需要保存扫描结果：scan_ip_port.exe 127.0.0.1 3000 4000 s

结果文件名称：scan_succ_ip_port.csv

#### 下载
linux 端下载 ：scan_ip_port

windows 端下载：scan_ip_port.exe

#### 源码

windows 打包

```shell
go build -o scan_ip_port.exe scan_port/scan_ip_port.go
```

mac 打包
```shell
set GOOS=darwin
set GOARCH=amd64
go build -o scan_ip_port_mac scan_port/scan_ip_port.go
```

linux 打包
```shell
set GOARCH=amd64
set GOOS=linux
go build -o scan_ip_port_linux scan_port/scan_ip_port.go
```

