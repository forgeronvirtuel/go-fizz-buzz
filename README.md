# fizzbuzz server

REST API to compute a fizz buzz program.

## Build

```sh
make build
```

Build the application in the `build` subdirectory.

```sh
make run
```

Run the application through `go run` on the `localhost:80` address.

## Use

By default, when you run the application, the listening address will be
localhost:80. You can give another address to the CLI, for instance:

```sh
./fizz-buzz-server 0.0.0.0:8080
```

## API

The main route aims to compute the fizz buzz:

```
http://localhost:9080/?int1=<int:min=1>&int2=<int:min=1>&limit=<int:min=1>&str1=<string>&str2=<string>
```

All parameters are required otherwise the API returns an error.

Examples:

```sh
curl http://localhost:9080/?int1=3&int2=5&limit=15&str1=fizz&str2=buzz
# Return "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz"

curl http://localhost:9080/?int1=3&int2=5&limit=15&str1=fizz
# Returns "{"error":"Key: 'fizzBuzzParam.Str2' Error:Field validation for 'Str2' failed on the 'required' tag"}"
```
