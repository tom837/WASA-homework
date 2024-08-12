# [Istruzioni]
# Creazione : docker build -t wasaphoto-backend:latest -f Dockerfile.backend . 
# Esecuzione : docker run -it --rm -p 3000:3000 wasaphoto-backend:latest (-p fa un link della porta 3000 dell'host alla 3000 del container,esposta)

#--------------------------------------------------------------------------------------------------------#
# [GOLANG]
# 1) Setting di golang 1.20 come immagine di base da cui partire
# 2) Creazione della cartella src in cui lavorarec
# 3) Copia dei files dell'intera directory WasaProject2022 all'interno della directory di lavoro (src)
# 4) Esecuzione del comando "go build .. " per creare il file da eseguire
#--------------------------------------------------------------------------------------------------------#

FROM golang:1.20 as backend_compiler
WORKDIR /src/
COPY . /src/
RUN go build ./cmd/webapi

#--------------------------------------------------------------------------------------------------------#
# [DEBIAN]
# 1) Setting di debian (versione stable) come immagine di base da cui partire
# 2) Creazione della cartella backend_binary in cui lavorare
# 3) Copia del file webapi dall'immagine "backend_compiler" nella directory di lavoro (backend_binary)
# 4) Esposizione della porta 3000 (il container sarà in ascolto sulla porta a runtime)
# 5) Setting istruzione di default quando viene eseguito il container da questa immagine
#--------------------------------------------------------------------------------------------------------#

FROM debian:stable
WORKDIR /backend_binary/
COPY --from=backend_compiler /src/webapi .
EXPOSE 3000
CMD ["./webapi"]