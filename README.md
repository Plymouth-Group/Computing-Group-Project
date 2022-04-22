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
      - [Download, Install and Setup PostgreSQL Server](#download-install-and-setup-postgresql-server)
      - [Setup Software for manage Local Database](#setup-software-for-manage-local-database)
      - [Setup Codinoc's Main Databases](#setup-codinocs-main-databases)
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

Follow [this tutorial](https://golangdocs.com/install-go-linux) or [this tutorial](https://linuxtect.com/how-to-install-go-golang-in-linux/) to install GO Language in any linux distribution

```sh
# Enable go modules

go env -w GO111MODULE=auto
```

```sh
# External GO Modules

go get github.com/gorilla/handlers &&
go get github.com/gorilla/mux &&
go get github.com/gorilla/handlers &&
go get github.com/lib/pq &&
go get github.com/gorilla/schema &&
go get github.com/prastuvwxyz/argon2 &&
go get golang.org/x/crypto/bcrypt &&
go get github.com/mailjet/mailjet-apiv3-go
```

### PostgreSQL Installation

#### Download, Install and Setup PostgreSQL Server

Follow [this tutorial](https://docs.fedoraproject.org/en-US/quick-docs/postgresql/) or [this tutorial](https://www.linuxshelltips.com/install-postgresql-in-fedora-linux/) to install PostgreSQL on Fedora

```sh
# Install PostgreSQL on Fedora

sudo dnf install postgresql-server postgresql-contrib

# Install PostgreSQL on Debian

sudo apt install postgresql postgresql-contrib
```

Following steps are common to any linux distribution

```sh
# Initialize Data Directory

sudo postgresql-setup --initdb --unit postgresql
```

```sh
# Configure PostgreSQL for multi application use

sudo nano /var/lib/pgsql/data/pg_hba.conf

# Now, edit
# host all all 127.0.0.1/32 ident -> host all all 127.0.0.1/32 md5
```

```sh
# System Enable PostgreSQL Server (Only for Production Server)

sudo systemctl enable postgresql

# System Enable PostgreSQL Server (Only for Production Server)

sudo systemctl disable postgresql

# Start PostgreSQL Server

sudo systemctl start postgresql

# Restart PostgreSQL Server

sudo systemctl restart postgresql

# stop PostgreSQL Server

sudo systemctl stop postgresql
```

```sh
# Login to PostgreSQl Server

sudo -u postgres psql
```

```sh
# Change PostgreSQL default user's (postgres) password

\password postgres

# Set Password into '1234'
```

```sh
# Create user role for our project

CREATE USER admin_codinoc WITH PASSWORD '1234';
```

**Analyzed Details**

| User            | Password | Host        | Port   | Details              |
| --------------- | -------- | ----------- | ------ | -------------------- |
| `postgres`      | `1234`   | `localhost` | `5432` | Default User         |
| `admin_codinoc` | `1234`   | `localhost` | `5432` | Used by this Project |

#### Setup Software for manage Local Database

In this project we use software called [DBeaver CE](https://dbeaver.io/) and follow [this link](https://www.tipsonunix.com/2022/02/install-dbeaver-ce-on-ubuntu-almalinux-fedora/) to install it on your Linux.

| Detail   | Value           |
| -------- | --------------- |
| Host     | `localhost`     |
| Port     | `5432`          |
| Database | `db_site`       |
| Username | `admin_codinoc` |
| Password | `1234`          |

#### Setup Codinoc's Main Databases

```sh
# Login to PostgreSQL

sudo -u postgres psql

# Create main site's database

CREATE DATABASE db_site OWNER admin_codinoc;

# Delete default 'public' schema

\c db_site
DROP SCHEMA public CASCADE;

# Create Server Template Database

CREATE DATABASE db_server OWNER admin_codinoc;
\c db_server
DROP SCHEMA public CASCADE;
```

Use DBeaver software to insert PostgreSQL database tables

1. Right click on database name in Database Navigator
2. Goto `SQL Editor -> Open SQL Script`
3. Paste PostgreSQL code
4. Press `Execute SQL Script` button

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
