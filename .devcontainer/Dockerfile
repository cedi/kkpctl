FROM qmcgaw/godevcontainer

RUN apk add make

RUN go get -u github.com/cweill/gotests/...
RUN go get -u golang.org/x/tools/cmd/benchcmp
RUN go get -u golang.org/x/lint/golint
RUN go get -u github.com/fzipp/gocyclo/cmd/gocyclo
RUN go get -u github.com/amitbet/gorename
RUN go get -u dlv-dap
RUN go get -u staticcheck

RUN export GO111MODULE=on
