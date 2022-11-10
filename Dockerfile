FROM golang:latest AS build
COPY src/ src/
WORKDIR bin/
RUN go build ../src/executor.go
RUN go build -tags netgo -a ../src/viewer.go

FROM alpine:latest AS deploy
COPY --from=build /go/bin/ /opt/test/
CMD ["/opt/test/viewer"]
EXPOSE 80