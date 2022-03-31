<div style="text-align: center; padding: 16px; border: 1px solid #555555; margin-bottom: 32px;">
	<h1 style="font-weight: 200;">Project <span style="font-weight: 600;">CODINOC</span> IDE</h1>
	<ttS>A Simple Cloud IDE for Web Designing by GO Language</p>
</div>

**Table of Contents**

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
	|- Binary       -> Binary Tarball
	|- Database     -> Database Designed Files
	|- Design       -> UI Designed Files
	|- Document     -> Documents
	|- Interface    -> Developed User Interface
	|- Readme       -> README Document Assets
	|- Server       -> Server Root
```

### Server Root Directory Structure

```
Server
	|- adm          -> Administrator Section
	|- bin          -> Web Application Executable
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
go get github.com/gorilla/mux &&
go get github.com/gorilla/handlers
```

### PostgreSQL Installation

1. Follow [this tutorial](https://docs.fedoraproject.org/en-US/quick-docs/postgresql/) or [this tutorial](https://www.linuxshelltips.com/install-postgresql-in-fedora-linux/) to install PostgreSQL on Fedora

> Note:\
> On the first login, need to create a Database:

```sh
psql template1
CREATE DATABASE sample_db;
```

### PostgreSQL Databases Setup

> TODO

**Additional Setup only for Development Server**

Add these content into the `hosts` file using `sudo nano /etc/hosts`

```
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

| Member Name | Plymouth Index | Leader Position | Github Profile |
| -- | -- | -- | -- |
| L.M. Nishshanka | | Group Leader | [@las-nish](https://www.github.com/las-nish) |
| E.C.N. Nandasiri | | Testing Leader | [@ITxChana](https://www.github.com/ITxChana) |
| M.P.M. Abeyrathne | | Programming Leader | [@mpmabeyrathne](https://www.github.com/mpmabeyrathne) |
| A.B. Navodya | | Technical Leader | [@ABNavodya](https://www.github.com/ABNavodya) |
| | | Quality Assurance Leader | |
| E.L.P. Prasandika | | Planning Leader | [@Prasandika](https://www.github.com/Prasandika) |
