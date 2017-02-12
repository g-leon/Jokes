# Jokes
A Buffalo based API that applies a Chuck Norris joke to a random person.

## Quick setup
    mkdir -p $GOPATH/src/github.com/g-leon
    mv ./Jokes $GOPATH/src/github.com/g-leon
    go get ./...

## Documentation

To view generated docs for Jokes, run the below command and point your browser to http://127.0.0.1:6060/pkg/

    godoc -http=:6060 2>/dev/null &

## Run Tests

    buffalo test    

## Run in dev

    buffalo dev
    
## Using the web service

     curl ‘http://localhost:3000’

    


