package main

import (
	"github.com/zyfdegh/ldap-change-password/ldap_pswd"
)

func main() {
	//baseDN
	baseDN := "dc=linkernetworks,dc=com"
	ldap_pswd.ChangePasswd(baseDN, "userId", "oldpassword", "newpassword")
}
