# Ldap-change-password
This project is used to change user password in openldap

## Compile and build source code
After you git clone this project, try **go build** 
If you get an error like this:

> fatal error: ldap.h: No such file or directory

You need to install the missing ldap library.
Install it in the command line according to your linux distributin.

* RedHat/CentOS

```sh
yum install -y openldap-devel
```

* Debian/Ubuntu

```sh
apt-get install openldap-devel
```

## Install

```sh
go get github.com/zyfdegh/
```

## Usage
Simply modify
