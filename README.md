# go-yq

[![Build Status](https://travis-ci.org/030/go-yq.svg?branch=master)](https://travis-ci.org/030/go-yq)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/go-yq)](https://goreportcard.com/report/github.com/030/go-yq)
![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/go-yq.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/go-yq.svg)
![Issues](https://img.shields.io/github/issues-raw/030/go-yq.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/go-yq.svg)
![Total downloads](https://img.shields.io/github/downloads/030/go-yq/total.svg)
![License](https://img.shields.io/github/license/030/go-yq.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/go-yq.svg)
![Contributors](https://img.shields.io/github/contributors/030/go-yq.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/go-yq.svg)
![Last commit](https://img.shields.io/github/last-commit/030/go-yq.svg)
![Release date](https://img.shields.io/github/release-date/030/go-yq.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/go-yq.svg)
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
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com/r/github.com/030/go-yq)

jq-style golang equivalent of [yq](https://github.com/kislyuk/yq). [Another yq tool that is written in golang](https://github.com/mikefarah/yq) could be used if one requires more features.

## rationale

go-yq was created to prevent that pip has to be installed in order to install yq

## installation

```
curl -L https://github.com/030/go-yq/releases/download/2.1.2/go-yq_2.1.2-0.deb -o go-yq.deb && \
sudo apt -y install ./go-yq.deb
```

## usage

### help

```
[user@localhost go-yq]$ ./go-yq -h
2018/11/13 12:07:39 Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>
exit status 1
```

### examples

```
[user@localhost go-yq]$ ./go-yq go run main.go .firefox_version ~/dev/ansible-firefox/defaults/main.yml
```

or

[![dockeri.co](https://dockeri.co/image/utrecht/go-yq)](https://hub.docker.com/r/utrecht/go-yq)

```
docker run -v /home/ben/dev/ansible-firefox:/ansible-firefox \
       -it utrecht/go-yq:2.1.0 .firefox_version \
       /ansible-firefox/defaults/main.yml
```

returns:

```
66.0.3
```
