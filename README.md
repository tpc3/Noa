# Noa
[![Docker Image CI](https://github.com/tpc3/Noa/actions/workflows/docker-image.yml/badge.svg)](https://github.com/tpc3/Noa/actions/workflows/docker-image.yml)  
Misskey Bot to post notes on the Markov chain.

# Use
## Docker
1. [Dowload config.yml](https://raw.githubusercontent.com/tpc3/Noa/main/config.yml)
2. Enter your token and host server to config.yml
3. `docker run --rm -it -v $(PWD)/config.yml:/go/src/Noa/config.yml ghcr.io/tpc3/noa`

## Build
1. Clone thos repository
2. `CGO_LDFLAGS -L/usr/lib/x86_64-linux-gnu -lmecab -lstdc++`
3. `CGO_CFLAGS -I/usr/include`
4. `go build`

### required
- golang
- mecab

# Contribution
Contributions are always welcome. (Please make issue or PR with English or Japanese)
