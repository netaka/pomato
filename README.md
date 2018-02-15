# pomato
Potarble on messaging app tiny organizer

## Feature
- Posting a text
- Supporting Slack and LINE
- Built-in api token

## How to build
1. Get packages

    ```
    $ go get github.com/netaka/pomato
    $ go get -u github.com/jteeuwen/go-bindata/...
    $ go get github.com/kazuph/go-binenv
    ```

1. Edit .env file
1. Convert .env files to bindata.go

    ```
    $ go-bindata .env
    ```

1. Build go files
   
   ```
   $ go build -o pomato main.go bindata.go
   ```

## How to use
```
$ pomato -m おやすみなさい
```
