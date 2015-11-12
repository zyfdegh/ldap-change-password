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

Get mqu/openldap if you do not have one in your local gopath:

```sh
go get github.com/mqu/openldap
```

## Usage
* Step 1:
Modify **host** and **port** in ldap_pswd.go,which point to a runningOpenLDAP server

* Step 2:
Change **baseDN** in main.go, change userId,oldPassword,newPassword argument of the function **ChangePassword**

* Step 3:
**go build** and run the executable file.
Check if any error occured in the log.

## Common errors

* Invalid credentials
Checkout userId and password

