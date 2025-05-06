## “Thread communication using memory sharing”

“In concurrent programming using memory sharing, we allocate a part of the process’s memory—for example, a shared data structure or a variable—and we have different goroutines work concurrently on this memory.

在使用内存共享的并发编程中，我们会分配一部分进程的内存——例如，
共享的数据结构或变量——并让不同的 goroutine 并发地在这块内存上工作。

### Race conditions 竞态条件

“Race conditions are what happens when your program is trying to do many things at the same time, and its behavior is dependent on the exact timing of independent unpredictable events. As we saw in the previous section, our letter frequency program ends up giving unexpected results, but sometimes the outcome is even more dramatic. Our concurrent code might be happily running for a long period, and then one day it may crash, resulting in more serious data corruption. This can happen because the concurrent executions are lacking proper synchronization and are stepping over each other”

竞态条件是指当你的程序试图同时执行多项任务时发生的情况，其行为依赖于独立且不可预测事件的精确时序。正如我们在前一节中看到的，我们的字母频率程序最终给出了意想不到的结果，但有时结果更加严重。我们的并发代码可能会长时间正常运行，然后某一天突然崩溃，导致更严重的数据损坏。这可能是因为并发执行缺乏适当的同步，导致相互之间发生冲突。
