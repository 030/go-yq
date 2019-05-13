# go-yq

[![Build Status](https://travis-ci.org/030/go-yq.svg?branch=master)](https://travis-ci.org/030/go-yq)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/go-yq)](https://goreportcard.com/report/github.com/030/go-yq)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=bugs)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=coverage)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_go-yq&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_go-yq)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2811/badge)](https://bestpractices.coreinfrastructure.org/projects/2811)

jq-style golang equivalent of [yq](https://github.com/kislyuk/yq). [Another yq tool that is written in golang](https://github.com/mikefarah/yq) could be used if one requires more features.

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
