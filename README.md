# fud

Feel Up Daemon(fud) is a simple web server.

It enables:

- checking if your LAN allows contact;
- checking if a given port is exposed and little more.

## Usage

```sh
Usage of fud:
  -msg string
    	message to show (default "This is up")
  -port int
    	port to serve (default 1337)
  -v	display version
```

## Sample

```sh
fud
```

Output:

```log
2026/02/09 16:24:29 port: 1337
2026/02/09 16:24:29 msg: "This is up"
2026/02/09 16:24:29 Listening in:
2026/02/09 16:24:29 - http://127.0.0.1:1337
2026/02/09 16:24:29 - http://::1:1337
2026/02/09 16:24:29 - http://XX.XXX.XX.XXX:1337
2026/02/09 16:24:29 - http://xxxx::xxxx:xxxx:xxxx:xxxx:1337
```

## Install

```sh
go install github.com/cpmachado/fud/cmd/fud@latest
```
