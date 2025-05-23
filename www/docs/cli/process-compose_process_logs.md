## process-compose process logs

Fetch the logs of a process(es). For multiple processes, separate them with a comma (proc1,proc2)

```
process-compose process logs [PROCESS] [flags]
```

### Options

```
  -f, --follow             Follow log output
  -h, --help               help for logs
  -N, --namespace string   Logs all the processes in the given namespace
      --raw-log            If set, don't format the multi process log output to include the process name
  -n, --tail int           Number of lines to show from the end of the logs (default 9223372036854775807)
```

### Options inherited from parent commands

```
  -a, --address string       address of the target process compose server (default "localhost")
  -L, --log-file string      Specify the log file path (env: PC_LOG_FILE) (default "/tmp/process-compose-<user>.log")
      --no-server            disable HTTP server (env: PC_NO_SERVER)
      --ordered-shutdown     shut down processes in reverse dependency order
  -p, --port int             port number (env: PC_PORT_NUM) (default 8080)
      --read-only            enable read-only mode (env: PC_READ_ONLY)
  -u, --unix-socket string   path to unix socket (env: PC_SOCKET_PATH) (default "/tmp/process-compose-<pid>.sock")
  -U, --use-uds              use unix domain sockets instead of tcp
```

### SEE ALSO

* [process-compose process](process-compose_process.md)	 - Execute operations on the available processes
* [process-compose process logs truncate](process-compose_process_logs_truncate.md)	 - Truncate the logs for a running or stopped process

