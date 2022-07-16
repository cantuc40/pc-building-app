FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go mod tidy
RUN go build -o /pcb_backend

EXPOSE 8083

CMD [ "/pcb_backend" ]