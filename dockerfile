FROM golang:alpine AS build
WORKDIR /go/src/myapp
ENV PROJECT_ROOT=/go/src/myapp
COPY . .
RUN go build -o /go/bin/myapp cmd/main.go


FROM scratch
COPY --from=build /go/bin/myapp /go/bin/myapp
ENTRYPOINT ["/go/bin/myapp"]
