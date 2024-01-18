# Usa una imagen de Go para construir tu aplicación
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de tu aplicación al contenedor
COPY . .

# Descarga las dependencias de Go
RUN go mod download

# Compila tu aplicación
RUN go build -o myapp main.go

# Expon el puerto en el que tu aplicación escuchará
EXPOSE 8080

# Comando para ejecutar tu aplicación al iniciar el contenedor
CMD ["./myapp"]
