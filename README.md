## UrlShortner

A url shortner that is written in Go with gin framework:
it's a simple url shortner using hash algorithm (more accurate(sha256))
it is exposed via localhost:9808

##How To Test
you can use any tool to check the result :                                                       
it has 2 endpoints :

localhost/9808/create-short-url 
you should post a request like that :     
````
{                                                                                             
    "long_url": "",                                                                           
    "user_id" : ""                                                                    
}
````

localhost/9808/:shortUrl
it is a get request you should pass the short url as a parameter 

## Run Test 
you can run test in docker-compose

````
docker exec -it <container name> bash

cd store or shortener

go test -v . -cover

````

