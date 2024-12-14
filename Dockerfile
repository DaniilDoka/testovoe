FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

CMD ["go", "run", "cmd/api/main.go"]
