# syntax=docker/dockerfile:1

##
## build stage
##
FROM golang:1.17-buster AS builder

WORKDIR /app

# copy source files
COPY . ./

# downloading go modules
RUN go mod download


# building artifact
RUN go build -o /my-golang-app

##
## running app
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

# copy templates (avoid issue loading templates)
COPY --from=builder /app/templates /templates

# copy artifacts to runner
COPY --from=builder /my-golang-app /my-golang-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/my-golang-app" ]
