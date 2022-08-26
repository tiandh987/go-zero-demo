#!/bin/bash

# 压力测试 /heavy 接口

# hey
#   压力测试工具
#
#   Usage:
#      hey [options...] <url>
#
#   Options:
#     -z  Duration of application to send requests. When duration is reached,
#         application stops and exits. If duration is specified, n is ignored.
#         Examples: -z 10s -z 3m.

hey -z 60s http://localhost:8080/heavy