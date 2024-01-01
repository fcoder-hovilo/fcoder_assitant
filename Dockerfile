FROM golang:latest

WORKDIR /app

COPY ./app .

# Print directory structure for debugging
RUN ls -l

RUN go mod download
RUN go mod tidy

# Command to run the application
CMD ["go", "run", "."]
