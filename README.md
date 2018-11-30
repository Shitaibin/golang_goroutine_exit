示例如何优雅的退出使用for-select的goroutine
=====

## 分支解释

- master：使用stop通道，主动告知goroutine退出
- stop_channel：和master相同
- detect_close_channel: 示例可以使用for-range替代for-select，range能检测通道关闭，自动退出
- detect_close_channel_v2：在前一个基础上，增加了监控功能，必须使用for-select，使用ok方法检测通道关闭，退出协程
