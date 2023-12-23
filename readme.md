```
go mod tidy
go run main.go

```
This builds an image and serves it via the go file server. Intended to be consumed by https://github.com/sweeneyb/png-to-epaper

Sample curl:

```curl -i -H "If-Modified-Since: Sun, 15 Oct 2023 00:59:39 GMT" 'http://localhost:8090/static/go-black.bmp'```

```curl -i --head -H "If-Modified-Since: Sun, 15 Oct 2023 00:59:39 GMT" 'http://localhost:8090/static/go-black.bmp'```

## Issues:
???

## Status
A bit in flux at the moment.  The esp32 didn't have the memory to decode 2 tiny/compressed pngs (one worked fine). I'm re-working to generate the image data in this server in a vaguely xbm format.  

The format that the e-paper wants is an 800\*480 bit array, or a byte array of size (800*480/8) where each bit represents a pixel. 

The current structure reflects my learning path and is probably inefficient.  Right now, I'm generating a full-color image, converting it to grayscale, then "compressing" into XBM.  The grayscale/contert to bit should be done at the same time.

Also, these files should be written to disk so that the file server can preserve the last-modified time.

## generate images from python
```
.venv/Scripts/activate
python3 generateImages.py
```

this generates much smaller files than the golang image gen