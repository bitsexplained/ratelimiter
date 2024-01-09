# Rate Limiter
## A customizable rate limiter for Go applications

Rate Limiter empowers you to control the rate of incoming requests in your Go applications, safeguarding against excessive traffic and potential overload. It offers a flexible configuration to tailor its behavior to your specific needs.

## Key features

- ```Customizable threshold```: Define the maximum number of requests permitted within a specified duration.
- ```Customizable duration```: Set the time window within which the threshold is applied.
- ```Reset```: Clear the rate limiter's internal state, enabling a fresh start.

## Installation

To integrate Rate Limiter into your project, execute the following command in your terminal:

```go
go get github.com/bitsexplained/ratelimiter
```

## Usage
- Here's a practical example showcasing Rate Limiter's usage:

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
Rate Limiter exposes two essential methods for interaction:

- ```go Allow() bool```: Determines whether a request is permissible based on the configured threshold and duration.

- ```Reset() ```: Clears the rate limiter's internal state, effectively starting anew.

## Configuration
- Customize Rate Limiter's behavior using these configuration options:

- ```WithThreshold(threshold int)```: Sets the maximum allowable requests within the specified duration.
- ```WithDuration(duration time.Duration)```: Establishes the time window for applying the threshold.


## Initializing the Rate Limiter
Construct a Rate Limiter instance using the ```NewRateLimiter``` function, seamlessly chaining configuration options:

```go
   rl := ratelimiter.NewRateLimiter(
  		ratelimiter.WithThreshold(500),
  		ratelimiter.WithDuration(time.Second),
	)
```

## Resetting the Rate Limiter
To clear the rate limiter's internal state, employ the Reset() method:

```go
   rl.Reset()
```
## Contributing
Your contributions to Rate Limiter are warmly welcomed! If you encounter bugs, have feature requests, or wish to contribute code, please create an issue or submit a pull request to the GitHub repository.

## License
rate limiter library is released under the MIT License.
See  [LICENSE](LICENSE) for details.
