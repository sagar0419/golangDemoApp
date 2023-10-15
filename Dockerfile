FROM golang:1.20-alpine
 
WORKDIR /demo/github.com

# ENV MYSQL_USER=root
# ENV MYSQL_PASSWORD=password
# ENV MYSQL_DATABASE=db
# ENV DBNAME=citadel
# ENV HSTNAME="192.168.1.10:3306"

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o github.com

EXPOSE 3000

CMD [ "./github.com" ]