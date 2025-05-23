## process-compose process

Execute operations on the available processes

### Options

```
  -a, --address string   address of the target process compose server (default "localhost")
  -h, --help             help for process
```

### Options inherited from parent commands

```
  -L, --log-file string      Specify the log file path (env: PC_LOG_FILE) (default "/tmp/process-compose-<user>.log")
      --no-server            disable HTTP server (env: PC_NO_SERVER)
      --ordered-shutdown     shut down processes in reverse dependency order
  -p, --port int             port number (env: PC_PORT_NUM) (default 8080)
      --read-only            enable read-only mode (env: PC_READ_ONLY)
  -u, --unix-socket string   path to unix socket (env: PC_SOCKET_PATH) (default "/tmp/process-compose-<pid>.sock")
  -U, --use-uds              use unix domain sockets instead of tcp
```

### SEE ALSO

* [process-compose](process-compose.md)	 - Processes scheduler and orchestrator
* [process-compose process get](process-compose_process_get.md)	 - Get a process state
* [process-compose process list](process-compose_process_list.md)	 - List available processes
* [process-compose process logs](process-compose_process_logs.md)	 - Fetch the logs of a process(es). For multiple processes, separate them with a comma (proc1,proc2)
* [process-compose process ports](process-compose_process_ports.md)	 - Get the ports that a process is listening on
* [process-compose process restart](process-compose_process_restart.md)	 - Restart a process
* [process-compose process scale](process-compose_process_scale.md)	 - Scale a process to a given count
* [process-compose process start](process-compose_process_start.md)	 - Start a process
* [process-compose process stop](process-compose_process_stop.md)	 - Stop running processes

