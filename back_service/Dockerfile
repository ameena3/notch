FROM golang
COPY . /go/src/github.com/ameena3/notch_project/back_service
WORKDIR /go/src/github.com/ameena3/notch_project/back_service
RUN go mod init 
RUN go mod tidy 
RUN go build ./cmd/main.go
EXPOSE 8080
ENTRYPOINT [ "./main" ]