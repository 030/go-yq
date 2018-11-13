# go-yq

Golang version of [yq](https://github.com/kislyuk/yq)

## Rationale

go-yq was created to prevent that pip has to be installed in order to install yq

## Usage

### help

```
[user@localhost go-yq]$ ./go-yq -h
2018/11/13 12:07:39 Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>
exit status 1
```

### example

```
[user@localhost go-yq]$ ./go-yq go run main.go .firefox_version ~/dev/ansible-firefox/defaults/main.yml
```

returns:

```
63.0.1
```

## Dependencies

```
go get github.com/spf13/viper
```
