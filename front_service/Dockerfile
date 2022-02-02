FROM golang
COPY . /go/src/github.com/ameena3/notch_project/front_service
COPY ./gen /go/src/github.com/ameena3/notch_project/back_service/gen
WORKDIR /go/src/github.com/ameena3/notch_project
RUN go mod init 
RUN go mod tidy 
RUN go build ./front_service/cmd/main.go
EXPOSE 9000
ENTRYPOINT [ "./main" ]