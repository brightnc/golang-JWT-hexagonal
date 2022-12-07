FROM golang:1.19
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go build -o main .
RUN ls -al
CMD ["/app/main"]