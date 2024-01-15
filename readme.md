# To Run
```
go mod tidy
go run main.go

```
This builds an image and serves it via the go file server. Intended to be consumed by https://github.com/sweeneyb/png-to-epaper

# To change the states
```curl http://localhost:8090/clear```

```curl http://localhost:8090/onAir```

```curl http://localhost:8090/offAir```


# Testing
Sample curl:
### To test what the device will pull
```curl -i -H "If-Modified-Since: Sun, 15 Oct 2023 00:59:39 GMT" 'http://localhost:8090/static/go-black.bmp'```

```curl -i --head -H "If-Modified-Since: Sun, 15 Oct 2023 00:59:39 GMT" 'http://localhost:8090/static/go-black.bmp'```


#Other details
## How it works
On startup, the server listens and serves files in the last state.  Hit one of the URLs above to change the state of the images.

When one of the above URLs is hit, an image (defined in code) is generated and written to disk.  The go server sets the last-modified header according to the disk modified time.  The client should come along with a modified-since header, and grab the images if they're new.  This should preserve state if the server reset and should avoid unnecessary fetching/painting of data. 

## Issues:
???

## Status
* The server works and generates images.  The math on the boarder is a bit odd at the moment, as the epaper seems to "eat" 20 pixels on the left and right.
* The format that the e-paper wants is an 800\*480 bit array, or a byte array of size (800*480/8) where each bit represents a pixel. 
* The current structure reflects my learning path and is probably inefficient.  Right now, I'm generating a full-color image, converting it to grayscale, then "compressing" into XBM.  The grayscale/convert to bit should be done at the same time.

## TODO
* Generating images in go is a pain and inflexible.  Since the code has to apply transforms to get into greyscale and compress to XBM, I should have some config that picks up files from a source area, and writes them to the static serving area after the transforms.  That would let me build better text/graphics/etc.  
