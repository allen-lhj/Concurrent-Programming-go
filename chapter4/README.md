## Synchronization with mutexes

Protecting critical sections with mutex locks使用互斥锁保护关键部分

Improving performance with readers–writer locks使用读者-写者锁提高性能

Implementing a read-preferred readers–writer lock实现读优先的读者-写者锁
  

“deFINITION Mutex, short for mutual exclusion, is a form of concurrency control with the purpose of preventing race conditions. A mutex allows only one execution (such as a goroutine or a kernel-level thread) to enter a critical section. If two executions request access to the mutex at the same time, the semantics of the mutex guarantee that only one goroutine will acquire access to the mutex. The other execution will have to wait until the mutex becomes available again.”

“互斥锁”，mutual exclusion的缩写，是一种并发控制形式，其目的是防止竞态条件。互斥锁允许只有一个执行（如 goroutine 或内核级线程）进入临界区。如果有两个执行同时请求访问互斥锁，互斥锁的语义保证只有一个 goroutine 能够获取到互斥锁。其他执行将不得不等待互斥锁再次可用。”


“In Go, mutex functionality is provided in the sync package, under the type Mutex. This type gives us two main operations, Lock() and Unlock(), which we can use to mark the beginning and end of our critical code sections, respectively. As a simple example, we can modify our stingy() and spendy() functions from the previous chapter to protect our critical sections. In the following listing, we’ll use the mutex to protect the shared money variable, preventing both goroutines from modifying it at the same time.
在 Go 语言中，互斥锁功能由 sync 包提供，在 Mutex 类型下。此类型为我们提供了两个主要操作， Lock() 和 Unlock() ，
我们可以使用它们分别标记临界代码段的开始和结束。
作为一个简单的例子，我们可以修改上一章中的 stingy() 和 spendy() 函数，以保护我们的临界区。
在下面的代码示例中，我们将使用互斥锁来保护共享的 money 变量，防止两个 goroutine 同时修改它。”


### How are mutexes implemented

Mutexes are typically implemented with help from the operating system and hardware. If we had a system with just one processor, we could implement a mutex just by disabling interrupts while a thread is holding the lock. This way, another execution will not interrupt the current thread, and there is no interference. However, this is not ideal, because badly written code can end up blocking the entire system for all the other processes and threads. A malicious or poorly written program can have an infinite loop after acquiring a mutex lock and crash the system. Also, this approach will not work on a system with multiple processors, since other threads could be executing in parallel on another CPU
互斥锁通常借助操作系统和硬件来实现。如果我们只有一颗处理器的系统，在一个线程持有锁时，可以通过禁用中断来实现互斥锁。这样，其他执行不会中断当前线程，也就不会有干扰。然而，这并不理想，因为写得不好的代码可能会导致整个系统中其他进程和线程全部被阻塞。一个恶意或写得不好的程序在获取互斥锁后可能会进入无限循环，从而导致系统崩溃。同时，这种方法在多处理器系统上不起作用，因为其他线程可能会在另一颗 CPU 上并行执行。

### Mutexes and sequential processing 互斥锁与顺序处理

见listing4.3_4

### Improving performance with readers-writer mutexes

At times, mutexes might be too restrictive.We can think of mutexes as blunt tools that solve concurrency problems by blocking concurrency.
But this might needlessing restrict performance and scalability for some applications.
Readers-writer mutextes give us a variation on standard mutexes that only block concurrency when we need to update a shared resource. Using readers-writer mutexes, we can improve the performance of read-heavy applications where we are doing a large number of read operations on shared data in comparsion with updates

有时，互斥锁可能过于限制性。我们可以把互斥锁看作是通过阻塞并发来解决并发问题的钝工具。一次只有一个 goroutine 能执行我们用互斥锁保护的临界区。这对于保证我们不会遭遇竞态条件非常有效，但对于某些应用来说，这可能不必要地限制了性能和可扩展性。读写互斥锁为我们提供了标准互斥锁的一个变体，仅在需要更新共享资源时阻塞并发。使用读写互斥锁，我们可以提升读操作占多数的应用的性能，这类应用中与更新相比，我们对共享数据执行了大量的读操作。

### Building our own read-preferred readers-writer mutex



