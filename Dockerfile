FROM golang:1.14 as build

WORKDIR /go/src/github.com/alantang888/always-error-401
COPY . .
WORKDIR /go/src/github.com/alantang888/always-error-401/cmd/always-error-401
RUN go build -o /go/bin/app


FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
EXPOSE 8080

CMD ["/app"]
