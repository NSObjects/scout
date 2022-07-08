
### 安装：

go 1.18 以上

```
go install github.com/NSObjects/scout@latest
```



## 指纹识别

-u 目标网址多个地址使用,分割

-p 代理

-c 并发,默认5

```
scout finger -u https://www.baidu.com,xxxx.com -p https://proxy.com -c 10
```

go install 

- [x]  CDN检测

- [x]  指纹识别

[ ] 防火墙检测

[ ] 子域名爆破

[ ] 端口扫描
