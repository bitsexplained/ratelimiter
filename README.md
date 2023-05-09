# rate limiter Limiter

rate limiter Limiter is a customizable rate limiter for Go applications that allows you to limit the rate of incoming requests based on a threshold and duration.

## Features

- ```Customizable threshold```: Set the maximum number of requests allowed within a specified duration.
- ```Customizable duration```: Set the time duration within which the threshold is applied.
- ```Reset```: Reset the rate limiter's internal state.

## Installation

To use rate limiter Limiter in your Go project, simply import it into your code:

```go
import "github.com/bitsexplained/ratelimiter"
```

## Usage
- Here's an example of how you can use rate limiter in your Go application:

```go
   package main

   import (
      "fmt"
      "github.com/bitsexplained/ratelimiter"
      "time"
   )

   func main() {
      // Create a new rate limiter with a threshold of 500 requests per second
      rl := ratelimiter.NewRateLimiter(
         ratelimiter.WithThreshold(500),
         ratelimiter.WithDuration(time.Second),
      )

      // Allow a request
      if rl.Allow() {
         // Process the request
         fmt.Println("Request allowed")
      } else {
         // Reject the request
         fmt.Println("Request rejected")
      }
   }
```

## Public Interface
rate limiter provides the following public interface:

- ```go Allow() bool```: Check if a request is allowed based on the rate limiter's threshold and duration.
- ```Reset() ```: Reset the rate limiter's internal state.

## Configuration
- rate limiter provides the following configuration options:

- ```WithThreshold(threshold int)```: Set the threshold for the rate limiter.
- ```WithDuration(duration time.Duration)```: Set the duration for the rate limiter.


## Initializing the Rate Limiter
You can initialize the rate limiter using the ```NewRateLimiter``` function and chaining the configuration options:

```go
   rl := ratelimiter.NewRateLimiter(
	ratelimiter.WithThreshold(500),
	ratelimiter.WithDuration(time.Second),
)
```

## Resetting the Rate Limiter
You can reset the rate limiter's internal state using the ```Reset()``` method:

```go
   rl.Reset()
```
## Contributing
Contributions to this library is welcomed! If you find a bug, have a feature request, or want to contribute code, please open an issue or submit a pull request to the GitHub repository.

## License
rate limiter library is released under the MIT License.
See  [LICENSE](LICENSE) for details.
