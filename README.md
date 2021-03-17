# 简介

v2ray 的 gomobile 绑定, 目前只支持 Android

# 函数

- `Start(jsonConfig string) error` 启动 V2ray
- `Stop()` 停止 V2ray
- `Status() int` 返回 V2ray 的状态, `-1` 则是未启动

# 编译

要求

- gomobile 和 gobind 已在环境变量中

开始编译:

```sh
make android
```
