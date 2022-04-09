
# CODINOC IDE

> A Simple Cloud IDE for Web Designing Tasks by GO Language

**Table of Contents**

- [CODINOC IDE](#codinoc-ide)
  - [Introduction](#introduction)
  - [Build Test](#build-test)
  - [Requirements](#requirements)
  - [Directory Structure](#directory-structure)
    - [Repository Directory Structure](#repository-directory-structure)
    - [Server Root Directory Structure](#server-root-directory-structure)
  - [Server Setup and Installation](#server-setup-and-installation)
    - [Main Directory Creation](#main-directory-creation)
    - [Go Language Installation](#go-language-installation)
    - [PostgreSQL Installation](#postgresql-installation)
    - [PostgreSQL Databases Setup](#postgresql-databases-setup)
    - [Restore Server](#restore-server)
  - [Documentation](#documentation)
  - [License and Copyrights](#license-and-copyrights)
  - [Development Team](#development-team)

## Introduction

> TODO

## Build Test

> TODO

## Requirements

> TODO

## Directory Structure

### Repository Directory Structure

```
Repository
  |- Database     -> Database Designed Files
  |- Document     -> Documents
  |- Interface    -> UI Designed Files
  |- Resources    -> Resources Files
  |- Server       -> Server Root
  |- Theme        -> Developed User Interface
```

### Server Root Directory Structure

```
Server
  |- adm          -> Administrator Section
  |- cdn          -> Storage
  |- cmd          -> Application Source
  |- thm          -> User Interface Theme
```

## Server Setup and Installation

### Main Directory Creation

1. Create temporary global variable in `.bashrc` file, `export c_path=/server_codinoc`
2. Test it using `echo $c_path`
3. Create Main directory using `sudo mkdir $c_path`
4. Add linux Read and Write permissions to to everyone using `sudo chmod 777 $c_path`
5. Copy server content into that `$c_path` directory

### Go Language Installation

1. Follow [this tutorial](https://golangdocs.com/install-go-linux) or [this tutorial](https://linuxtect.com/how-to-install-go-golang-in-linux/) to install GO Language in any linux distribution
2. Setup GO Packages used by the project

```sh
go env -w GO111MODULE=auto
```

```
go get github.com/gorilla/handlers &&
go get github.com/gorilla/mux &&
go get github.com/gorilla/handlers &&
go get github.com/lib/pq
```

### PostgreSQL Installation

1. Follow [this tutorial](https://docs.fedoraproject.org/en-US/quick-docs/postgresql/) or [this tutorial](https://www.linuxshelltips.com/install-postgresql-in-fedora-linux/) to install PostgreSQL on Fedora

```sh
sudo dnf install postgresql-server
sudo dnf install postgresql-contrib

# Initialize Data Directory for PostgreSQL

sudo postgresql-setup --initdb --unit postgresql

# Start PostgreSQL Server

sudo systemctl start postgresql

# Stop PostgreSQL Server

sudo systemctl stop postgresql
```

### PostgreSQL Databases Setup

> We didn't create another user and we use default PostgreSQL user role here

1. Login to PostgreSQL and [change password](https://serverfault.com/questions/406606/postgres-error-message-fatal-ident-authentication-failed-for-user)

```sh
sudo su postgres
psql

\password

# Then enter and re-enter new password

# Restart postgresql server

sudo systemctl restart postgresql
```

_hint: use `1234` as password_

2. Find [postgresql details](https://stackoverflow.com/questions/5598517/find-the-host-name-and-port-using-psql-commands)

```sh
\conninfo
```

The default values for local server is, `user = postgres` and `port = 5432`

3. [Create](https://www.tutorialspoint.com/postgresql/postgresql_create_database.htm) codinoc site's database

```sql
CREATE DATABASE db_site;
```

4. [Upload database](https://www.a2hosting.com/kb/developer-corner/postgresql/import-and-export-a-postgresql-database) that we're created already

_hint: Open terminal in the location that we're stored our database files_

```sql
sudo su postgres
psql -U postgres db_site < db_site.psql
```

5. Show uploaded Database

```sql
```

**Additional Setup only for Development Server**

Add these content into the `hosts` file using `sudo nano /etc/hosts`

```sh
# Codinoc IDE

127.0.0.1:8080 codinoc.com
```

So after that, we can access to site site using `codinoc.com:8080/` in web browser

### Restore Server

```sh
# Remove PostgreSQL

sudo dnf remove postgresql-server postgresql-contrib
```

## Documentation

> TODO

## License and Copyrights

> TODO

## Development Team

| Member Name       | Plymouth Index | Leader Position          | Github Profile                                         |
| ----------------- | -------------- | ------------------------ | ------------------------------------------------------ |
| L.M. Nishshanka   |                | Group Leader             | [@las-nish](https://www.github.com/las-nish)           |
| E.C.N. Nandasiri  |                | Testing Leader           | [@ITxChana](https://www.github.com/ITxChana)           |
| M.P.M. Abeyrathne |                | Programming Leader       | [@mpmabeyrathne](https://www.github.com/mpmabeyrathne) |
| A.B. Navodya      |                | Technical Leader         | [@ABNavodya](https://www.github.com/ABNavodya)         |
|                   |                | Quality Assurance Leader |                                                        |
| E.L.P. Prasandika |                | Planning Leader          | [@Prasandika](https://www.github.com/Prasandika)       |
