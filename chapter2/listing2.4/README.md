## User-level threads executing within a single kernel-level thread用户级线程在单个内核级线程中执行

“We are replicating on a small scale what the OS does, in terms of thread scheduling and management, inside the main thread of the process
我们在小规模上复制了操作系统在线程调度和管理方面的功能，即在进程的主线程内部。

The main advantage of user-level threads is performance. Context-switching a suer-level thread is faster than context-switching a kernel-level one.
用户级线程的主要优势是性能。用户级线程的上下文切换比内核级线程的上下文切换更快。

“用户级线程的缺点在于它们执行调用阻塞 I/O 调用的代码时。考虑我们需要从文件中读取的情况。由于操作系统将进程视为只有一个执行线程，如果用户级线程执行这个阻塞读取调用，整个进程将被取消调度。如果同一进程中存在其他用户级线程，它们将无法执行，直到读取操作完成。这并不理想，因为拥有多个线程的一个优势是在其他线程等待 I/O 时执行计算。为了克服这种限制，使用用户级线程的应用程序往往使用非阻塞调用来执行它们的 I/O 操作。然而，使用非阻塞 I/O 并不理想，因为并非每个设备都支持非阻塞调用。”

“用户级线程的另一个缺点是，如果我们有一个多处理器或多核系统，在任何时候我们只能利用一个处理器。操作系统将包含所有用户级线程的单个内核级线程视为单个执行。因此，操作系统在单个处理器上执行内核级线程，所以包含在该内核级线程中的用户级线程不会真正并行执行。”

“Go provides a hybrid system that gives us the great performance of user-level threads without most of the downsides. It achieves this by using a set of kernel-level threads, each managing a queue of goroutines. Since we have more than one kernel-level thread, we can utilize more than one processor if multiple ones are available.”
“Go 提供了一个混合系统，它为我们提供了用户级线程的优秀性能，而没有大多数缺点。它是通过使用一组内核级线程来实现的，每个线程管理一个 goroutine 队列。由于我们有多个内核级线程，因此如果可用，我们可以利用多个处理器。”

## “M:N hybrid threading”

“The system that Go uses for its goroutines is sometimes called the M:N threading model. This is when you have M user-level threads (goroutines) mapped to N kernel-level threads. This contrasts with normal user-level threads, which are referred to as an N:1 threading model, meaning N user-level threads to 1 kernel-level thread. Implementing a runtime for M:N models is substantially more complex than other models since it requires many techniques to move around and balance the user-level threads on the set of kernel-level threads.”

“Go’s runtime determines how many kernel-level threads to use based on the number of logical processors. This is set in the environment variable called GOMAXPROCS. If this variable is not set, Go populates this variable by querying the operating system to determine how many CPUs your system has. You can check how many processors Go sees and the value of GOMAXPROCS by executing the following code.”

“Go 使用的 goroutines 的系统有时被称为 M:N 线程模型。这是指有 M 个用户级线程（goroutines）映射到 N 个内核级线程。这与通常的用户级线程形成对比，通常称为 N:1 线程模型，意味着 N 个用户级线程对应 1 个内核级线程。实现 M:N 模型的运行时比其他模型复杂得多，因为它需要许多技术来在内核级线程集上移动和平衡用户级线程。”

“Go 的运行时根据逻辑处理器的数量确定要使用多少个内核级线程。这个值在名为 GOMAXPROCS 的环境变量中设置。如果这个变量没有设置，Go 将通过查询操作系统来确定您的系统有多少个 CPU。您可以通过执行以下代码来检查 Go 看到的处理器数量和 GOMAXPROCS 的值。”
