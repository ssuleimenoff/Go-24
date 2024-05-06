FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o football_teams cmd/main.go

CMD ["./football_teams"]