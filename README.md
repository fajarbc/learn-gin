# Learn Gin Framework (Golang)

## Description
Learning golang framework gin.
Implements:
   - Gin
   - Gorm MySQL CRUD (currently: create, read)
   - JWT Authentication
   - Docker using Dockerfile

There is nothing in http://localhost:8080 yet, but you can do :
   - Register/Login as an Author
   - Save/Get Your Articles

Use Postman for completes endpoints: https://www.getpostman.com/collections/e66fcfef85e70524e3ac

## Endpoint examples
An example of `POST` request to save new Article with its Author to http://localhost:8080/api/articles 
```json
{
    "title": "First Post Title",
    "content": "Lorem ipsum this is article content",
    "slug": "post-1",
    "status": 1
}
```
Make sure you are already authenticated, otherwise it will respond with 400.
To authenticate, you need to :
1. Register, make request to http://localhost:8080/author/register
   ```json
   {
      "name": "Fajar",
      "email": "user@mail.com",
      "status": 1,
      "username": "user",
      "password": "user"
   }
   ```
2. Login, make request to http://localhost:8080/author/login
   ```json
   {
      "username": "user",
      "password": "user"
   }
   ```

## Install
After you clone this repo
```bash
git clone https://github.com/fajarbc/learn-gin.git && cd learn-gin
```
Then, before you could run the app
1. install/downloads the dependecies first.
   ```
   go mod download
   ```
2. make database `go_articles`
3. (optional) For development, hot reloading using [Air](https://github.com/cosmtrek/air)
   1. Install Air
      ```bash
      go get github.com/cosmtrek/air@latest
      ```
   2. Add alias to `.bashrc`
      1. File check where `.bashrc` should be
         ```bash
         echo ~
         // output: C:/Users/fajarbc
         ```
      2. Check your air executable, in my case i found it on `C:\Users\fajarbc\go\bin\bin`
      3. Next, go to `C:/Users/fajarbc` and add this line to file `.bashrc` (create it yourself if it's not created yet) :
         ```bash
         alias air='~/go/bin/bin/air'
         ```
      Note: `~` equals `C:\Users\fajarbc`
   3. Check in your terminal by type ```air -v```


## Run
There are 4 ways to run the app. Once it runs, you could reach `http://localhost:8080`
1. Using go run, just run this in terminal
   ```
   go run server.go
   ```
2. Using air (Install step 3), it will automatically hot reload your golang app, i found this the most comfortable way for developing process for me so far. Just type this to start
   ```bash
   air
   ```
3. Using go build, there are 2 steps :
   1. Build the binary artifact
      - Windows
         ```
         go build -o server.exe
         ```
      - Linux
         ```
         go build -o server
         ```
   2. Run the binary artifact
4. Using Dockerfile, there are 2 steps :
   1. Build docker image
      ```
      docker build --tag learn-gin:v0.1 .
      ```
   2. Run docker image as container (also expose port 8080 and detach mode)
      ```
      docker run -d -p 8080:8080 learn-gin:v0.1
      ```
