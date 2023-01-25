# Go: Waiting after killing a process

A minimal example showing what happens when you call `exec.Cmd.Kill()` on a process that you're waiting for via `exec.Cmd.Wait()`.
This also applies when the passed context is closed when using `exec.CommandContext()` as it uses `Kill()` under the hood.

## Usage

```shell
make run
```

## Results

Aside from probably getting myself on some kind of government watchlist for _killing children_ (processes), the results showed that using `exec.CommandContext()` sends `SIGKILL` to the command process, not giving it a chance to clean up after itself.

It's also not configurable with which signal is sent to the process -- it's always kill.
Perhaps this is because it's the only signal supported across all OSs -- as of the time of writing, Go doesn't support sending interrupt signals on Windows.

```
manager: Running interrupt child
child: Awaiting signal
child: Received signal: interrupt
child: Exiting
manager: child interrupted



manager: Running kill child
child: Awaiting signal
manager: child killed
```
