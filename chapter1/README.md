curl -s https://www.rfc-editor.org/rfc/rfc1122.txt | wc

when we run this command on a unix system, the command line is forking two concurrent processes. we can check this by opening another terminal and running ps -a

## Concurrency with threads

Processes are the heavyweight answer to concurrency. They provide us with good isolation, but they consume lots of resources and take a while to create.