# Notifier
Searches a given website for a text string and sends a text when it's found or not depending on the set environment variable.

## Build
Building is done using make, so requires you have that installed.

### Local Build
Then run:
```
BINARY_NAME=notifierlambda make build
```

### Build for deploying
To build it for deploying it needs to be built for linux and packaged in a zip.
Either of these will do it:
```
BINARY_NAME=notifierlambda make build-linux
make build-all-linux
```

## Run
Once built it can be run locally using:
```
BINARY_NAME=notifierlambda make build
```

Note: although it's worth swapping out the following lines in notifierlambda.go's main
method before building it if running locally.
```
lambda.Start(handleRequest)
// handleRequest()
```

