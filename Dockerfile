FROM golang:1.17-alpine

WORKDIR /app



#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download

#COPY *.go ./

COPY . .
RUN go mod download
#RUN go get "github.com/go-sql-driver/mysql"


#RUN go mod tidy
#RUN go get
RUN go build -o /pcb_backend

EXPOSE 8083

CMD [ "/pcb_backend" ]