# docker run --rm --name gogo --workdir /app --volume ./:/app -p 3000:3000 golang:latest go run main.go
FROM golang:latest

WORKDIR /app

EXPOSE 3000

#CMD ["go", "run" , "main.go"]