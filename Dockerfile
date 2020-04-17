FROM golang:1.14 as build

WORKDIR /go/src/github.com/alantang888/custom-error-map-backend
COPY . .
WORKDIR /go/src/github.com/alantang888/custom-error-map-backend/cmd/custom-error-map-backend
RUN go build -o /go/bin/app


FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
EXPOSE 8080

CMD ["/app"]
