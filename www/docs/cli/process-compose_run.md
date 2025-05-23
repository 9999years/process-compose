## process-compose run

Run PROCESS in the foreground, and its dependencies in the background

### Synopsis

Run selected process with std(in|out|err) attached, while other processes run in the background.
Command line arguments, provided after --, are passed to the PROCESS.

```
process-compose run PROCESS [flags] -- [process_args]
```

### Options

```
  -f, --config stringArray   path to config files to load (env: PC_CONFIG_FILES)
      --disable-dotenv       disable .env file loading (env: PC_DISABLE_DOTENV=1)
  -h, --help                 help for run
      --no-deps              don't start dependent processes
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

