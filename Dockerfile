FROM golang:1.14 AS server-builder
WORKDIR /app
COPY . ./
RUN go mod vendor
RUN GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -a -o server cmd/web/main.go

FROM node:12.13.0-alpine AS web-builder
WORKDIR /app/web
ENV PATH /app/node_modules/.bin:$PATH
COPY web/anagram/ ./
RUN yarn global add react-scripts@4.0.3
RUN yarn install
RUN yarn build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=web-builder /app/web/build ./public
COPY --from=server-builder /app/server .
EXPOSE 4000
CMD ["./server"]
