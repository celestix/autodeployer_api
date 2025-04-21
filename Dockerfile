FROM golang:1.23-alpine AS builder
WORKDIR /builder
ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64
COPY . . 
RUN go mod tidy -e 
RUN go build -o server 

FROM gcr.io/distroless/cc AS runtime
WORKDIR /root 
COPY --from=builder /builder/server ./server
COPY ./config.yaml ./config.yaml
EXPOSE 8000
ENTRYPOINT [ "/root/server" ]
