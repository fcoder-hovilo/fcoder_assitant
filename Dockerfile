FROM golang:latest

WORKDIR /app

COPY ./app .

# Print directory structure for debugging
RUN ls -l

RUN go mod download
RUN go mod tidy

# Linux version
# RUN groupadd -r myuser && useradd -r -g myuser -G audio,video myuser \
#     && mkdir -p /home/myuser/Downloads \
#     && chown -R myuser:myuser /home/myuser

# USER myuser

USER chrome

# Autorun chrome
ENTRYPOINT [ "chrome" ]

# Command to run the application
CMD ["go", "run", "."]
