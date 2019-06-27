# Build indicator

[![Go Report Card](https://goreportcard.com/badge/github.com/egel/build-indicator)](https://goreportcard.com/report/github.com/egel/build-indicator)
[![Build Status](https://travis-ci.org/egel/build-indicator.svg?branch=master)](https://travis-ci.org/egel/build-indicator)
[![Coverage Status](https://coveralls.io/repos/github/egel/build-indicator/badge.svg?branch=master)](https://coveralls.io/github/egel/build-indicator?branch=master)
[![MIT LICENSE](http://img.shields.io/badge/license-MIT-yellowgreen.svg?style=square)](https://github.com/egel/dotfiles/blob/master/LICENSE)

> This small tool use [USB-MiniTrafficLight](http://www.cleware-shop.de/epages/63698188.sf/en_US/?ObjectPath=/Shops/63698188/Products/41/SubProducts/41-1) to display  build status from your repository branch.

We support:

- Gitlab v4 (official, private instances)

## Use docker image

```bash
docker pull egel/build-indicator:latest
docker run -i -t --privileged egel/build-indicator:latest /bin/bash
clewarecontrol -l
```

## Build from source

```bash
docker build -t egel/usb-mini-traffic-light:latest .
docker run -i -t --privileged egel/usb-mini-traffic-light:latest /bin/bash
```

