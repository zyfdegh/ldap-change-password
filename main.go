package main

import (
	"fmt"

	"github.com/zyfdegh/ldap-change-password/ldapasswd"
)

func main() {
	host := "192.168.10.103"
	port := "8389"
	//	dn := "cn=admin,dc=linkernetworks,dc=com"

	dn := "uid=linker@linkernetworks.com,dc=linkernetworks,dc=com"
	ldapasswd.ChangePasswd(host, port, dn, "password", "newpassword")

	fmt.Println("End")
}
