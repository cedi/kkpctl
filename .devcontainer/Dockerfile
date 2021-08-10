# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.137.0/containers/go/.devcontainer/base.Dockerfile
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends python3.7 python3-pip
#
# RUN python3.7 -m pip install -U setuptools
# RUN python3.7 -m pip install -U checkov


RUN export GO111MODULE=on
RUN go get -u github.com/cweill/gotests/...
RUN go get -u golang.org/x/tools/cmd/benchcmp
RUN go get -u golang.org/x/lint/golint
RUN go get -u github.com/fzipp/gocyclo/cmd/gocyclo
RUN go get -u github.com/amitbet/gorename