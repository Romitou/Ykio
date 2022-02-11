FROM alpine:3.15 AS build

RUN apk add go

WORKDIR /app/go/
ADD . .
ENV GOPATH /app

RUN go get
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o ykio

FROM alpine:3.15

WORKDIR /app
COPY --from=build /app/go/ykio /app/ykio
RUN chmod +x ./ykio

CMD ["./ykio"]
