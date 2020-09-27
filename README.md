# loglibrary
改善能直接调用并且可以自由选择终端或者文件或者网络方式format输出日志
设计思路：
	易用性封装
	默认情况下执行直接输出到终端，如果读取到配置文件里面有logpath的情况下 输出到logpath，如果有kafka配置就输出到kakfa
	输出文件名行号
	文件输出不同 errorlog 输出errorlog文件 常规输出常规