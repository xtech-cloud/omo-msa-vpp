# 简介 

虚拟专用代理 （Virtual Private Proxy）

支持http和rpc两种调用方式

# 编译

`go 版本需要大于1.11`

进入omo-msa-vpp目录，执行go install

# 运行
编译完成后运行 ~/go/bin/omo-msa-vpp


# 部署

部署时候，可以按需设置环境变量，以下是默认值

```
export GIN_MODE=release              //默认为 debug
export VPP_HTTP_ADDR=:80
export VPP_HTTPS_ADDR=:443
export VPP_WS_ADDR=:8000
export VPP_CONFIG=/etc/vpp/vpp.yaml  //默认为./conf/vpp.yaml
export VPP_TLS_CRT=/etc/vpp/tls.crt  //默认为./conf/tls.crt
export VPP_TLS_KEY=/etc/vpp/tls.key  //默认为./conf/tls.key
```

设置完环境变量后启动omo-msa-vpp

# 配置文件范例


```
http:
  filter:
    mod: 'blacklist'
    url:
      - '/console/*'
      - '/account/signup'
  forward:
    -
      endpoint: '/ams/*'
      remote: 'http://192.168.0.10:16000'
    -
      endpoint: '/kms/*'
```

过滤模式支持白名单(whitelist)和黑名单(blacklist)两种模式。
白名单模式，在url列表中，没有匹配到地址，会返回405错误。
黑名单模式，在url列表中，匹配到地址，会返回405错误。

没有匹配到正确的转发地址，会返回500错误。




