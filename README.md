# learn-gin

Learning golang framework gin.
Focus only on build api (no database connection).
There is nothing in http://localhost:8080 yet, but you can do Login and Save/Get Videos using these endpoints here with Postman https://www.getpostman.com/collections/e66fcfef85e70524e3ac

## Endpoint examples 
An example of `POST` request to save new Article with its Author to http://localhost:8080/api/articles 
```json
{
    "title": "First Post Title",
    "content": "Lorem ipsum this is article content",
    "slug": "post-1",
    "status": 1,
    "author": {
        "name": "John Doe",
        "email": "john.doe@gmail.com",
        "status": 1
    }
}
```
Make sure you are already authenticated, otherwise it will respond with 400.

# install
Before you could run the app
1. install/downloads the dependecies first.
   ```
   go mod download
   ```
2. make database `go_articles`

# running
There are 3 ways to run the app. Once it runs. You could reach `http://localhost:8080`
1. Using go run, just run this in terminal
   ```
   go run server.go
   ```
2. Using go build, there are 2 steps :
   1. Build the binary artifact

    (in windows)
   ```
   go build -o server.exe
   ```
    (in linux)
   ```
   go build -o server
   ```
   2. Run the binary artifact
3. Using Dockerfile, there are 2 steps :
   1. Build docker image
    ```
    docker build --tag learn-gin:v0.1 .
    ```
   2. Run docker image as container (also expose port 8080 and detach mode)
    ```
    docker run -d -p 8080:8080 learn-gin:v0.1
    ```
