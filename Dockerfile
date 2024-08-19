FROM golang:1.22.5

WORKDIR "/project"

COPY . .

EXPOSE 8000

CMD ["./run"]