# go-microservices-url-shortener(wip)

## Concepts
*   **url-shortener-service**: grpc server to  implement url shortener 
*   **client-service**: allow the user to perform CRUD operations, and  perform those operations on the `url-shortener-service` via gRPC
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
*   **Method**: Get
 
*   **Response**:
  ```json
{
  "error": false,
  "message": "retrieved URL Shortener successfully",
  "data": {
    "result": [
      {
        "url_path": "https://www.google.com/some/path",
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


*   **Desc**: Update URL
*   **EndPoint**:{domain-name}/all-url-shortener/{id}
*   **Method**: Put
*   **Payload**:
  ```json
    {
        "url_path": "https://www.google/a/new/pretty/long/url",
    }
  ```
*   **Response**:
  ```json
{
    "error": false,
    "message": " URL was updated successfully",
    "data": {
        "id": 1,
        "url_path": "https://www.google.comdsaasa/adadsasdas",
        "shortcut": "de141f"
    }
}
    
```



*   **Desc**: Get URL by Shortcut
*   **EndPoint**:{domain-name}/all-url-shortener/{shortcut}
*   **Method**: Get
*   **Response**:
  ```json
{
    "error": false,
    "message": " URL was retrieved successfully",
    "data": {
        "id": 1,
        "url_path": "https://www.google.comd/long/path/1/3/4/3",
        "shortcut": "f5a9e9"
    }
}
    
```



*   **Desc**: Delete Url
*   **EndPoint**:{domain-name}/url-shortener/{id}
*   **Method**: Delete
*   **Response**:
  ```json
{
    "error": false,
    "message": " URL was deleteted successfully",
  
}
    
```