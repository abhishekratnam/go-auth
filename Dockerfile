FROM golang:latest
RUN mkdir /build
ADD . /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go mod tidy
RUN go get github.com/abhishekratnam/go-auth
RUN cd /build && git clone https://github.com/abhishekratnam/go-auth.git
RUN cd /build/go-auth && go build -o main .
EXPOSE 8080
EXPOSE 3306
ENTRYPOINT ["/build/go-auth/main"]

