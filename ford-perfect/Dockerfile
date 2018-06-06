FROM golang AS build
ARG package=eu.gcr.io/hacker-playground/ford-perfect
WORKDIR /go/src/$package
COPY . .
# -d: Download only without install
# -t: Download test dependencies
RUN go get -d -t ./...
RUN go test ./...
# CGO_ENABLED=0: Disables CGO tools (Go packages that call C code)
# -o: Write executable
RUN CGO_ENABLED=0 GOOS=linux go build -o /service .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /service .
ENTRYPOINT ["/service"]
