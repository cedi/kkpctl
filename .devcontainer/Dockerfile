FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.16

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends python3.7 python3-pip

RUN python3.7 -m pip install -U setuptools
RUN python3.7 -m pip install -U checkov

RUN apt install wget
RUN wget https://golang.org/dl/go1.17.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
	&& tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz

RUN go get -u github.com/cweill/gotests/...
RUN go get -u golang.org/x/tools/cmd/benchcmp
RUN go get -u golang.org/x/lint/golint
RUN go get -u github.com/fzipp/gocyclo/cmd/gocyclo
RUN go get -u github.com/amitbet/gorename

RUN export GO111MODULE=on
