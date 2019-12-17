# hconf

## provider

## consul

consul 添加和查看配置

``` sh
curl -X PUT 127.0.0.1:8500/v1/kv/test_namespace/test_service --data-binary @test.json
curl -X GET 127.0.0.1:8500/v1/kv/test_namespace/test_service\?raw=true
```

## 链接

- toml 语法: <https://github.com/toml-lang/toml>
- properties 文件格式: <https://docs.oracle.com/cd/E23095_01/Platform.93/ATGProgGuide/html/s0204propertiesfileformat01.html>
- ini 维基百科: <https://en.wikipedia.org/wiki/INI_file>
