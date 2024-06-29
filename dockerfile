FROM golang:1.21
LABEL "authors"="Chibby-k Ezeiru"

# Create app directory
WORKDIR /usr/local/app

COPY . .

RUN go build -o bin ./cmd/app

ENTRYPOINT [ "/usr/local/app/bin" ]
