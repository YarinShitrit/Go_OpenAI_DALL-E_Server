# Go_OpenAI_DALL-E_Server

This is a simple Go server written using the Gin framework for retrieving images from DALL-E using a free text query to the DALL-E API. The server returns a JSON object containing a link to the generated image.

## Usage

To retrieve an image, send a GET request to the /image endpoint with the query parameter set to your desired text query:
```
http://localhost:8080/image?query=my%20text%20query
```
