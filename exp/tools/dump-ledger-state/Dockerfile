FROM ubuntu:18.04

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates wget curl gnupg
RUN wget -qO - https://apt.stellar.org/SDF.asc | apt-key add -
RUN echo "deb https://apt.stellar.org xenial stable" | tee -a /etc/apt/sources.list.d/SDF.list
RUN apt-get update -y

RUN apt-get install -y stellar-core jq
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ $(env -i bash -c '. /etc/os-release; echo $VERSION_CODENAME')-pgdg main" | tee /etc/apt/sources.list.d/pgdg.list && \
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    apt-get update && \
    DEBIAN_FRONTEND="noninteractive" apt-get install -y postgresql-9.6 postgresql-contrib-9.6 postgresql-client-9.6

# Create a PostgreSQL role named `circleci` and then create a database `core` owned by the `circleci` role.
RUN  su - postgres -c "/etc/init.d/postgresql start && psql --command \"CREATE USER circleci WITH SUPERUSER;\" && createdb -O circleci core"

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
RUN echo "host all all all trust" > /etc/postgresql/9.6/main/pg_hba.conf

# And add `listen_addresses` to `/etc/postgresql/9.6/main/postgresql.conf`
RUN echo "listen_addresses='*'" >> /etc/postgresql/9.6/main/postgresql.conf

RUN curl -sL https://storage.googleapis.com/golang/go1.15.3.linux-amd64.tar.gz | tar -C /usr/local -xz
RUN ln -s  /usr/local/go/bin/go /usr/local/bin/go
WORKDIR /go/src/github.com/AnneNamuli/go-stellar
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

ENV PGPORT=5432
ENV PGUSER=circleci
ENV PGHOST=localhost

WORKDIR /go/src/github.com/AnneNamuli/go-stellar/exp/tools/dump-ledger-state

ARG GITCOMMIT
ENV GITCOMMIT=${GITCOMMIT}

ENTRYPOINT ["./docker-entrypoint.sh"]