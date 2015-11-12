package main

import (
	"fmt"

	"github.com/zyfdegh/ldap-change-password/ldap_pswd"
)

func main() {
	//Method 1:
	//	fmt.Println("Start")
	//	ldap_pswd.ChangePasswd("linker2", "password", "newpassword")
	//	fmt.Println("End ChangePasswd()")

	fmt.Println("")

	//Method 2:
	ldap_pswd.ChangePasswd2("linker2", "password", "newpassword")
}
