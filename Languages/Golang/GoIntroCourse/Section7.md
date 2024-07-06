# Section 7 - Channels and Go Routines
Go routines allow us to perform tasks in parallel.  A Go routine happens within our
running program (a process). A single Go routine executes everything we are doing
line by line, sequentially. Where we can use Go routines is on function calls that 
require time to complete- such as an HTTP request. These function calls block our
main Go routine. To fix this, we can introduce additional Go routines. To do this,
we can use the following syntax:
`go checkLink(link)`
This will generate a new, independent go routine (thread) to run the code within the checkLink function.

## Go Scheduler
A Go scheduler runs on a single core on our machine, even if you have a multi-core machine.
The scheduler will then run *one* Go routine until it completes or reaches a
blocking call, then move onto the next one. You are able to overwrite the Go scheduler
so that it takes advantage of more cores.

*Concurrency is not parallelism*

Concurrency => Multiple threads executing code. If one blocks, another is picked up
Parallelism => Multiple threads executing at the same time (by running them on different cores)

When we launch a Go program, it created a single Go routine that is referred to as
the main routine. Every one created from within the program are child routines.
The child routines behave differently than the main routine.

Since the Go routine may end before the child routines end, we need to utilize Go
channels in order to stop this behavior. A channel is an intermediary between all
of the running routines. When you create a channel, you define the type of data
that you are sending between the routines to communicate.
`c := make(chan string)`

We can wait for a message to come into a channel by doing
`fmt.Println(<- c)`
This would be a blocking call. It would wait until a message is received before continuing.

**Never reference the same variable within two Go routines**
