# URLShortener
URLShortener is a go application for generating a short url for any given url.

Currently 2 endpoints are available for the API.
 
 /generate --> generates a short url for an entered url  
 /getall --> retreives all the generated short urls
 
 API spec :  
 POST:http://localhost:8000/generate  
 Request Body :  
 {
  "origUrl": "input url"
 }
 
 Response Body:  
 Generated short url
 
 GET:http://localhost:8000/getall  
 Response Body:  
 [
  {
    "id": "generated hash id",  
    "origUrl": "input url",  
    "shortUrl": "generated short url"  
  }
]
 

# Getting started
1. Download/clone the repo
2. Download the 2 dependencies - gorilla/mux and go-hashids  
  (go mod download)
3. Build the code after setting the GOOS  
  (go build)
  
 # Running the URLShortener
 1. Run the executable file from your terminal   
  (.\URLShortener.exe)
 2. Once the application and the server is up and running, launch any API client
 3. Reach the endpoints
