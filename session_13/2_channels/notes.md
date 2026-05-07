Main Goroutine
    |
    |---- creates channel
    |
    |---- starts another goroutine
                    |
                    |---- sends "ping" into channel
    |
    |---- waits to receive from channel
    |
    |---- receives "ping"
    |
    |---- prints "ping"