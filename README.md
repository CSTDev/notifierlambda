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
or wrapped up as a zip and deployed to AWS

It requires the following environment variables to be set:</br>
REGION - AWS region</br>
URL - Web URL to search for text</br>
SEARCH_TEXT - Text to be searched for</br>
PHONE_NUMBER - Phone number to send the SMS message to</br>
ALERT_IF_PRESENT - (optional) true if you want to alert if a page has text</br>
                                false if you want to alert if text is not present</br>
