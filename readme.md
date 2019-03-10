## how to build

```
$ git clone git@github.com:mangoqiqi/MiniHttpFileServer.git
$ cd MiniHttpFileServer
$ go build .

or
$ docker build -t MiniHttpFileServer .
$ docker run MiniHttpFileServer
```
## usage
`MiniHttpFileServer server -d {path} -p {port}` command use to start a server to expose path.
`MiniHttpFileServer upload -p {image path}` command use to upload a image to https://sm.ms and return a Markdown style link.

## thanks
Thanks to the platform provided by [https://sm.ms](https://sm.ms).



