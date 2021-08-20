# 🖼️ Ykio
Simple image hosting server for personal use.

## 🚀 Deploying
### 🐳 Docker
This application is intended to be used with Docker for convenience and flexibility. An image is published on the [GitHub container registry](https://github.com/Romitou/Ykio/pkgs/container/ykio) for each new commit on the main branch. In order to start using Ykio, it is necessary to first set up a [PostgreSQL database](https://hub.docker.com/_/postgres). Additional network configuration of your containers **may be required** to connect your Ykio container to PostgreSQL. In order for your images to be persistent on every restart of the container, we recommend that you set up a **volume** to mount on your container.
```bash
docker run --volume ykio_images:/app/images \
    --publish 80:8080 \
    --restart unless-stopped \
    --env DB_DSN=host=<host> user=<user>  password=<password> dbname=<dbname> port=5432 sslmode=disable TimeZone=Europe/Paris \
    --env SEND_TOKEN=<generate your token> \
    --detach \
    ghcr.io/romitou/ykio:latest
```
### ⚙️ Standalone
Just like with Docker, you need to set up a [PostgreSQL database](https://www.postgresql.org/download/). We don't provide any binary application ready to run, so you'll have to build Ykio yourself. To do this, you will need to download [Git](https://git-scm.com/downloads) and [Go 1.17+](https://golang.org/doc/install) on your machine. Here are the steps to follow to build Ykio.
```bash
git clone https://github.com/Romitou/Ykio.git
cd Ykio
go get
CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o ykio
chmod +x ./ykio
./ykio
```
