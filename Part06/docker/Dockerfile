FROM ubuntu:20.04
RUN apt-get update && apt-get upgrade -y 
RUN apt-get install software-properties-common apt-transport-https apt-utils wget -y
RUN wget -O- https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor | tee /usr/share/keyrings/postgresql.gpg
RUN echo deb [arch=amd64,arm64,ppc64el signed-by=/usr/share/keyrings/postgresql.gpg] http://apt.postgresql.org/pub/repos/apt/ focal-pgdg main | tee /etc/apt/sources.list.d/postgresql.list
RUN apt-get update
RUN apt-get install postgresql-client postgresql -y
RUN apt-get update && apt-get install -y git golang vim
USER postgres
RUN /etc/init.d/postgresql start && psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" && createdb -O docker docker
RUN echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/14/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/14/main/postgresql.conf
USER root
RUN mkdir files
COPY webServer.go files
WORKDIR files
RUN go get github.com/lib/pq
RUN go build webServer.go
RUN ls -l
USER postgres
CMD ["/usr/lib/postgresql/14/bin/postgres", "-D", "/var/lib/postgresql/14/main", "-c", "config_file=/etc/postgresql/14/main/postgresql.conf"]
