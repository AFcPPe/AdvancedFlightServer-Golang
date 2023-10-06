
# AdvancedFlightServer-Go

[English Version](Readme.md)

## 简介

AdvancedFlightServer-Go 是一个用 Go 语言编写的模拟飞行联机服务器，支持 Swift、Echo、Euroscope 或其他自定义的客户端。

## 特点

- 更快的速度
- 使用 FSD Version 3.000 Draft 9 协议
- 使用 JSON 进行数据输出
- 使用 PostgreSQL 数据库进行账户存储
- 支持高并发
- 跨平台支持（开发测试环境为 Windows 11 和 CentOS 8，理论上支持所有 Golang 支持的平台）
- 解决了 FSD V3.000 Draft 9 在 Linux 系统运行时随机崩溃的问题
- 自带的METAR数据源

## 使用指南

1. 使用 `go build AdvancedFlightServer/main` 进行编译。
2. 运行编译后的可执行文件。
3. 第一次运行程序会生成一个 `config.ini` 文件，请先关闭程序。
4. 填写 `config.ini` 文件中的相关条目。
5. 重新打开程序即可。

## 开源许可

本软件使用 AGPL v3.0 开源许可。

## 贡献

欢迎通过提交 PR 的方式为本代码库做出贡献。

## 商业版本

除了开源版本外，我们还提供商业版本，该版本具有以下特性：

- 使用 Vatsim v3.4 协议
- 支持管制员修改飞行计划
- 支持 Euroscope 内置 ATIS
- 支持管制员 INFO line 和 Logoff time
- 支持 VisualPosition 刷新（飞机在地面上时，每 0.2 秒钟刷新一次位置信息）
- 支持锁定飞行计划（管制员修改过的飞行计划自动锁定，直到下线）
- 支持标牌操作记录等

请注意，商业版本的详细信息和许可要求可能会有所不同。如有兴趣了解商业版本，请联系我们获取更多信息。