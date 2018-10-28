# go-yq

Golang version of [yq](https://github.com/kislyuk/yq)

## Rationale

go-yq was created to prevent that pip has to be installed in order to install yq

## Usage

### help

```
[user@localhost go-yq]$ ./go-yg -h
Usage of ./go-yg:
  -debug
    	Whether debugging should be enabled
  -key string
    	Specify the key (default "key")
  -yamlFile string
    	Path to a yaml file (default "file.yaml")
```

### example

```
[user@localhost go-yq]$ ./go-yg -yamlFile /home/user/dev/ansible-firefox/defaults/main.yml -key firefox_version
```

returns:

```
62.0.3
```