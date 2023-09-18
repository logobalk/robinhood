# syntax=docker/dockerfile:1
FROM golang:1.21.1 AS build

# Set destination for COPY
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY app ./app
COPY domain ./domain
COPY repo ./repo
COPY usecase ./usecase
COPY utils ./utils
WORKDIR /app/app

# Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /app/publish/robinhood 
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/publish/godocker 

FROM scratch AS final

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8081
WORKDIR /app/app
COPY --from=build /app/publish .
COPY --from=build /app/app /app/app
COPY --from=build /app/domain /app/domain
COPY --from=build /app/repo /app/repo
COPY --from=build /app/usecase /app/usecase
COPY --from=build /app/utils /app/utils



# Run
ENTRYPOINT ["/app/app/godocker"]