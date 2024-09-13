# Simple Go Server

[![](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/pkg/)

## Use

```bash
bash run.sh
```

## Endpoints

1. [localhost:8888/heartbeat](http://localhost:8888/heartbeat)
    * Simple String to HTML writeout representing a dummy Heart Beat / Status / Test endpoint.
1. [http://localhost:8888/public/](http://localhost:8888/public/)
    * Also: [http://localhost:8888/public/index.html](http://localhost:8888/public/index.html)
    * Simple HTML Static File Serve
1. [http://localhost:8888/stop](http://localhost:8888/stop)
    * Immediately terminate the application.
    * Note that calling `os.Exit(0)` in the main Process will also terminate each Go Routine.
