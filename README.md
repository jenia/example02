# Image processing server
## How to run
You can run this server by cloning this repository and running in the terminal:
```
docker build -t bimg-server .
docker run -p 8080:8080 bimg-server
```

Usually, you'd want to run this server in as a Docker container due to ease of library 

## API Endpoints Documentation
### Resize Image
`POST /resize`

Description: Resizes the uploaded image to the specified dimensions.

Request Parameters:
- image (file, required): The image file to be resized.
- x (string, required): The desired width for the resized image.
- y (string, required): The desired height for the resized image.

Response:
- Returns the resized image as a binary stream.

Example:
```
curl -X POST http://localhost:8080/resize \
  -F "image=@path/to/image.jpg" \
  -F "x=800" \
  -F "y=600" \
  -o resized.png
```

### Convert Image
`POST /convert`

Description: Converts the uploaded image to a different format (e.g., JPEG to PNG).

Request Parameters:
- image (file, required): The image file to be converted.

Response:
- Returns the converted image as a binary stream.

Example:
```
curl -X POST http://localhost:8080/convert \
  -F "image=@path/to/image.jpg" \
  -o converted.png
```
### Compress Image
`POST /compress`

Description: Compresses the uploaded image, reducing its file size.

Request Parameters:
- image (file, required): The image file to be compressed.

Response:
- Returns the compressed image as a binary stream.

Example:
```
curl -X POST http://localhost:8080/compress \
  -F "image=@path/to/image.jpg" \
  -o compressed.png
```

## Run tests
To run the tests do:
```
go test ./...
```

You will notice that the test fails with:
```
--- FAIL: Test_Given_Image_When_Resize_Then_Dimensions_Are_Correct (0.01s)
    image_test.go:27: image size want 400 but is: x 20 and y 20
```

For some reason the bimg library does not perform any of the transformations. I'm not 100% sure why but the problem is isolated to how I use the bimg library itself. To resolve the problem for a production system, we can raise an issue on the bimg Github repository with the sample code and test.
