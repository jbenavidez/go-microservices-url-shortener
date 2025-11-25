# go-microservices-url-shortener(wip)

## Concepts
*   **url-shortener-service**: grpc server to  implement url shortener 
*   **client-service**: allow the user to perform CRUD operations, and  perform those operation on the `url-shortener-service` via gRPC
*   **postgres**: database

 

### Stack
<ul>
<li>Go</li>
<li>gRPC</li>
<li>PostgreSQL </li>
</ul>


### How to run the project 
*   **Start containers**: cd to project -> then run the following commands 
  ```bash
$ make up_build
```
The command will build the binary files for the `client-service` and `url-shortener-services`, and   it will start the containers


## URL Shortener APIs
*   **Desc**: Create URL Shortener 
*   **EndPoint**:{domain-name}/url-shortener
*   **Method**: Post
*   **Payload**:
  ```json
    {
        "full_path": "https://www.google/a/pretty/long/url",
    }
  ```
*   **Response**:
  ```json
{
    "error": false,
    "message": "URL Shortener successfully",
    "data": {
        "full_path": "https://www.google/a/pretty/long/url",
        "shortcut": "9024e8"
    }
}
    
```


*   **Desc**: Get All
*   **EndPoint**:{domain-name}/all-url-shortener
*   **Method**: get
 
*   **Response**:
  ```json
{
  "error": false,
  "message": "retrieved URL Shortener successfully",
  "data": {
    "result": [
      {
        "url_path": "adsdsds",
        "shortcut": "adsdsds22"
      },
      {
        "url_path": "https://www.google.comdsaasa/adadsasdas",
        "shortcut": "9024e8"
      }
    ]
  }
}
    
```