# Ldap-change-password
This project is used to change user password in ldap

## Compile & Install
After you git clone this project, try **go build** and you may get an error like this:
> fatal error: ldap.h: No such file or directory

This is caused by lack of ldap library.Install it in the command line

* RedHat/CentOS

```sh
yum install -y openldap-devel
```

* Debian/Ubuntu

```sh
apt-get install openldap-devel
```

## Usage

