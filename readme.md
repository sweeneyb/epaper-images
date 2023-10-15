```
go mod tidy
go run main.go

```
This builds an image and serves it via the go file server.

Sample curl:
```curl -i -H "If-Modified-Since: Sun, 15 Oct 2023 00:59:39 GMT" 'http://localhost:8090/static/go-black.bmp'```

## Issues:
* Color depth is 24-bit, which is insane for 1 B/W images.  

## generate images from python
```
.venv/Scripts/activate
python3 generateImages.py
```

this generates much smaller files than the golang image gen