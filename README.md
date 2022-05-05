# Concurrency in Go
## A design philosophy

Presented at the IMX Bootcamp
04 May 2022

Link below will take you to the presentation

* [Golang Basics](https://talks.godoc.org/github.com/alexi21/concurrency-in-go/presentation.slide)

### Speaker notes:

We are going to discuss the background and context to concurrency in Go, and some of the features of the language that enable concurrency.
Most importantly thank you all for being here!

I want to start with some simple premises that are hopefully not too controversial
- Concurrency in Go is a design philosophy
- Readability is the most important feature of a programming language
- And, composition not inheritance
While not exhaustive, each of these go to the core of programming in Go, and I want to try and keep them in mind during this presentation

* Why concurrency in go?

Based on the primitives and idioms of Go, researching the motivation for its creation, and listening to or reading the thoughts of people such as Rob Pike,
one of the creators of Go, my thesis is that concurrency in Go is a design philosophy. Which is to say, Go was created with concurrency as one of its core principles.
Go allows us to write composable sequential code, that remains reasonable and correct, while also allowing us to create sophisticated concurrent structures.
In his talk "Concurrency is not parallelism", Rob Pike argues that the world is not object orientated, it is parallel!
However, by and large, programming languages have not been designed with that world view in mind.
Given the rise of multicore CPUs, it would seem self-evident that a language should provide first-class support for some sort of concurrency or parallelism

* CSP

The concurrency primitives and patterns in Go owe a great deal to Tony Hoare's paper _Communicating Sequential Processes_ written in 1978
In which he describes three architectural patterns for effectively writing concurrent code
Guarded commands as described by Dijkstra in his paper from 1975, "Guarded Commands, Nondeterminancy and Formal Derivations of Programs"
A parallel command, specifying concurrent execution of its constituent sequential processes
And the necessity of communication between concurrent processes
These form the building blocks for a code structure that enables the concurrent composition of communication channels and sequential procedures.

While Go was written to facilitate such concerns as simplifying dependency management, fast compilation, and providing a minimal yet familiar programming interface.
It was also designed with the principle of building concurrency into the language, and designing concurrency primitives to be first class citizens,
allowing them to be easily parameterized and shared between procedures
Go provides tools which allow us to structure our code using concurrency, in a manner that makes concurrency useful.
Goroutines provides us with a way of doing things concurrently,
channels provide the ability to communicate between processes that are executing concurrently,
and the select statement provides a guarded command which allows us to control and multiplex concurrent communication and synchronization.

* Structure vs Execution

In the context of Go:
Concurrency is a way of structuring your program to make it easier to understand and scalable, and parallelism is just the execution of multiple goroutines in parallel
If concurrency is about structure, then it would seem vital that we can both communicate and synchronize the concurrent processes as required.
Go provides the tools to do just that, and it does so in a manner that follows the principles of CSP

* CSP

There are three principles to CSP, build each individual process for sequential execution, coordinate these processes using channels
- do not share state. Finally, in order to scale, simply replicate the pattern.

A typical sequential process is an IO operation, read a file - process the data - and the write to a file.
We can imagine introducing concurrency into the data processing, however we cannot process the data before reading the file, and so on.

We can think of a channel in terms of the pipe command, we use the pipe command in UNIX based systems to write the output of one command
into the input of the following command. A channel can be shared, and unless buffered stores no data such that it must be read from in order to
write to it. If buffered it is like a FIFO queue, data is read from the channel in the same order as it was written.

* Coordination and synchronization

Channels can be passed to and from goroutines, and shared across multiple readers and writers.
In order to orchestrate the sequential processes we can use the select statement, which we will look at shortly, to control execution using channels

* Composition

The possibility of parallelism is enabled through the correct design of concurrent processes composed together.
The approach is thus the composition of independently executing functions of otherwise regular procedural code.
At its core, through the semantics of interfaces, Go encourages composition over inheritance.
Interface composition (which is its own subject) demands its own style of writing code, clearly the goal of composition is at the heart of Go.
The point of this is to understand that the design process should be based on the composition of concurrent processes.
This means breaking a problem down into component parts that can be run concurrently and composed together to solve the problem.

* Goroutines

Lets say we instantiate a function in Go. The program will run the function and then exit, before moving to the next sequential part of the program.
If instead we define the function call by placing the word "go" in front of it, the function starts running and "concurrently" the program continues to run sequentially.
This may or may not occur at the same time, but at an abstract level we can think of the program as forking at that point and now moving conceptually along two paths.
They are "like" threads, they run in the same address space, and they are very cheap, they are multiplexed dynamically onto OS threads as required.
The system takes care of all scheduling and blocking, and that scheduling and blocking is done independently between all running goroutines.

* Channels

_"Don't communicate by sharing memory; instead, share memory by communicating"_ Rob Pike

* Deadlock

A deadlock in general terms is when you have two processes both waiting for the other to unblock so they can continue.
The deadlock detector in Go only works if the entire program deadlocks
In the previous example, the block created by the attempt to read from the empty channel locks the main function indefinitely creating a deadlock

* Select

The select statement is used to bring channels together and enables us to compose channels and create more sophisticated abstractions

- `select` is used as "switching" statement for channel operations
- `select` uses case statements, however the case statements are not tested sequentially
- The execution will block if none of the cases are met
- The `default` operator can be used as a fall back if otherwise execution is blocked

