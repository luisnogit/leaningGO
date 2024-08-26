FROM golang:alpine3.20 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
RUN air init
COPY go.mod ./
RUN go mod download

COPY . .

CMD [ "air" ]

FROM golang:alpine3.20 AS deploy

WORKDIR /app

# Copy only go.mod and go.sum to cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

CMD ["make run"]
