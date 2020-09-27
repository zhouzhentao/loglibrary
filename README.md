# loglibrary
### 设计思路：
- 默认情况下执行直接输出到终端，如果读取到配置文件里面有logpath的情况下 输出到logpath，如果有kafka配置就输出到kakfa
- 输出log输出位置的文件名行号
- 文件输出不同 errorlog 输出errorlog文件 常规输出outlog文件
- 易用性封装
- 异步输出log