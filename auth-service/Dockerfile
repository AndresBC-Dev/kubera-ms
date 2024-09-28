# Etapa 1: Build
FROM golang:bookworm AS builder

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos del servicio
COPY . .

# Compilamos el binario
RUN go mod tidy
RUN go build -o auth-service

# Etapa 2: Imagen de producción
FROM debian:bookworm-slim

# Creamos un nuevo directorio para la aplicación
WORKDIR /app

# Copiamos el binario compilado desde la imagen builder
COPY --from=builder /app/auth-service /app/auth-service

# Definir el comando por defecto
CMD ["/app/auth-service"]
