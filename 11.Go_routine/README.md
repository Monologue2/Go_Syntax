# GoRoutine

### Summary
- Basic Goroutines: Use the go keyword to start a new goroutine.
- Anonymous Functions: Run inline anonymous functions as goroutines.
- Synchronization: Use sync.WaitGroup to wait for goroutines to complete.
- Channels: Use channels for communication between goroutines.
- Unbuffered Channels: Block until both the sender and receiver are ready.
- Buffered Channels: Allow a fixed number of values to be sent without blocking.
- Select Statement: Wait on multiple channel operations.
- Context for Cancellation: Manage the lifecycle of goroutines with the context package.
- Worker Pool: Process tasks concurrently with a pool of worker goroutines.
<br>
Goroutine이 언제 종료되는지 모른다면 실행을 시키지 말 것.<br>
Contaxt를 사용하여 확실하게 Goroutine의 실행 시점과 종료 시점을 설계할 것