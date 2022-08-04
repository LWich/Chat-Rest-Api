FROM golang

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . .

RUN go build -o . cmd/api/main.go

EXPOSE 8080
EXPOSE 4545

CMD [ "./main" ]