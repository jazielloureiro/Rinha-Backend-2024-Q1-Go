FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server

FROM scratch

COPY --from=build /app/server /

CMD ["/server"]