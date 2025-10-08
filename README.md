# demo-lernia-golang

Enkelt demo som visar hur vi kan använda multistage build med podman.

Containerfile kommer bygga en container som installerar golang och bygger `src/main.go` filen och sedan körs ett nytt bygge där vi kopierar filen `src/main` från `build`-containern in i den container vi bygger.

För att bygga:

```bash
$ podman build -t webserver:latest .
```

För att köra:

```bash
$ podman run -d -p 8080:8080 webserver
```

Testa med:

```
$ curl localhost:8080
```


