## I/O model notes

### Golang IO Cookbook

[source](https://github.com/jesseduffield/notes/wiki/Golang-IO-Cookbook)

![](https://i.imgur.com/kpIFNir.png)

io.Writer and io.Reader are complementary interfaces for streaming information. An io.Reader's job is to take data from some source and write it to a transfer buffer. An io.Writer's job is to take a transfer buffer and write its contents to some destination.

There are a couple of benefits to the streamed approach:

Because we're only siphoning the data through a small transfer buffer, there's no need to load the entire source file into memory
You don't need to wait for the whole file to be read before you can start encrypting it. As soon as you've read some data into the transfer buffer, you can start writing that to the destination file.

...

In the examples so far we've directly given our readers/writers buffers, but io.Copy is the function which allows you to link up a reader with a writer so that you don't need to manually handle buffers. io.Copy uses a 32kb buffer and siphons data from the reader through the buffer and to the writer. In each iteration the buffer is given to the reader's Read method, then however much of the buffer gets populated is passed on to the writer's Write method.

...

Unlike io.Copy, whose job is to send data from a reader to a writer, the job of a pipe is to make it possible to send data from a writer to a reader (typically enlisting the help of io.Copy).

...

### Streaming IO in Go

[source](https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185)

In Go, input and output operations are achieved using primitives that model data as streams of bytes that can be read from or written to. To do this, the Go io package provides interfaces io.Reader and io.Writer, for data input and output operations respectively.

Go comes with many APIs that support streaming IO from resources like in-memory structures, files, network connections, to name a few. This writeup focuses on creating Go programs that are capable of streaming data using interfaces io.Reader and io.Writer using custom implementations as well as those from the standard library.

...

This writeup shows how to use the io.Reader and io.Writer interfaces to implement streaming IO in your program. After reading this write up you should be able to understand how to create programs that use the io package to stream data for IO. There are plenty of examples and the writeup shows you how to create your own io.Reader and io.Writer types for custom functionalities.
This is an introductory discussion and barely scratches the surface of the breath of and scope of Go packages that support streaming IO. We did not go into file IO, buffered IO, network IO, or formatted IO, for instance (saved for future write ups). I hope this gives you an idea of what is possible with the Go’s streaming IO idiom.