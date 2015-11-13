package main

import (
	"github.com/zyfdegh/ldap-change-password/ldap_pswd"
)

func main() {
	host := "192.168.10.197"
	port := "8389"
	//	dn := "cn=admin,dc=linkernetworks,dc=com"

	dn := "uid=linker@linkernetworks.com,dc=linkernetworks,dc=com"
	ldap_pswd.ChangePasswd(host, port, dn, "password", "newpassword")
}
