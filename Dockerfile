# Usamos una imagen base de Go ligera
# FROM golang:latest as builder
FROM --platform=linux/amd64 golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

ARG VERSION="development"
ARG BUILD="1999-01-01T00:00:00Z"

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos necesarios para la descarga de dependencias
ADD go.mod .
# ADD go.sum .
# descargamos las dependencias

RUN go mod download

# copiamos el resto de archivos
COPY . .

# Compilamos el binario
RUN go build -ldflags="-X 'main.Version=${VERSION}' -X main.Build=${BUILD}" -o /tmp/rps ./main.go

ENTRYPOINT ["./app"]

# Usamos alpine para mantener nuestro contenedor lo más ligero posible
FROM scratch

WORKDIR /app

# Copiamos el binario compilado desde el primer paso
COPY --from=builder /tmp/rps ./rps

# Exponemos el puerto
EXPOSE 8080

# Ejecutamos el microservicio
ENTRYPOINT ["/app/rps"]