# Simple Go Server

[![](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/pkg/)

## Use

```bash
bash run.sh
```

## Endpoints

1. [localhost:8888/print](http://localhost:8888/print)
    * Simple String to HTML writeout
2. [localhost:8888/write ](http://localhost:8888/write)
    * Simple Model String to HTML writeout
3. [http://localhost:8888/public/](http://localhost:8888/public/)
    * Also: [http://localhost:8888/public/index.html](http://localhost:8888/public/index.html)
    * Simple HTML Static File Serve
4. [http://localhost:8888/stop](http://localhost:8888/stop)
    * Immediate terminate the application
