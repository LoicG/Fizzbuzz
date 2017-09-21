# FizzBuzz Server

```
./fizzbuzz.exe -h

Simple fizz-buzz REST server [OPTIONS]
routers & frameworks:
        - gorilla (default): https://github.com/gorilla/mux
        - goji: https://github.com/zenazn/goji/
        - emicklei: https://github.com/emicklei/go-restful
  -port int
        server port (default 8084)
  -router string
        router (default "gorilla")
```

# FizzBuzz Stress Client

```
./fizzbuzzclient.exe -h

client sends http requests to the fizz-buzz server
  -error
        if true sends invalid request
  -jobs int
        concurrent jobs (default 1)
  -limit int
        fizzbuzz limit (default 16)
  -number int
        number request per jobs (default 1)
  -port uint
        server port (default 8084)
```
