# GoUploadSvr

A very simple server for upload files


# Usage

Use curl to upload `./go_test.ipynb`:
```
curl -F "file=@./go_test.ipynb" http://127.0.0.1:9090/upload
```

# TODO

One click to start 3 server: upload server, static database server, download support server.