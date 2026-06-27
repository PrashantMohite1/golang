

**channel does not create a loop**.
**channel is just a communication pipe** between goroutines.


A channel simply transfers a value from one goroutine to another.
After one send and one receive match, that communication is complete.

Once a goroutine starts a send (or receive) operation on an unbuffered channel, that goroutine cannot continue until another goroutine performs the matching receive (or send).

-  1. Send data to the channel
```
channelName <- data
```

- Receive data from the channel
```
<- channelName
```

Let's look at a simple example.

```go
func sender(ch chan string) {
    fmt.Println("Step 1")

    ch <- "Hello"

    fmt.Println("Step 2")
}

func receiver(ch chan string) {
    msg := <-ch
    fmt.Println(msg)
}
```

Execution:

```
sender
|
| Step 1
|
| ch <- "Hello"      (waits)
|
------------------------------>
                       receiver
                       |
                       msg := <-ch
                       |
                       gets "Hello"
```

Once the receiver receives `"Hello"`,

the channel's work for **that send/receive operation** is finished.

Then both goroutines continue executing the remaining code.

```
sender                    receiver

Step 2                    Print Hello

Function ends             Function ends
```
