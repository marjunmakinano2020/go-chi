# go-chi

## Overview

RESTful API with Go and chi

## Setup

1. Run the following command to install the project's dependencies and generate a go.sum file that contains checksums of the modules that this project depends on

```
	go mod tidy
```

2. Run the project

```
	go run .
```

3. If you would like to run the service on a different port, then set the environment variable PORT to a different port number:

```
	PORT=8888 go run .
```

4. Test the service by sending a POST request to /posts:

```
	curl -X POST -d '{"userId":1,"title":"Lorem Ipsum","body":"Foo Bar"}' http://localhost:8080/posts
```

## References

https://www.newline.co/@kchan/building-a-simple-restful-api-with-go-and-chi--5912c411