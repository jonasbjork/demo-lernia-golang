FROM ubi9:latest AS build

RUN dnf -y install golang
RUN mkdir /src
COPY src/main.go /src/
RUN go build -C /src/ main.go

FROM ubi9-minimal
RUN mkdir /app
COPY --from=build /src/main /app/webserver

ENTRYPOINT ["/app/webserver"]
