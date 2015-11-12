package ldap_pswd

import (
	"log"

	"github.com/mqu/openldap"
)

func ChangePasswd2(userId string, oldPasswd string, newPasswd string) {
	connectLdap2()
}

func connectLdap2() {
	host := "192.168.10.194"
	port := "8389"
	//admin
	//	userDn := "cn=admin,dc=linkernetworks,dc=com"
	//	passwd := "changeit"

	//linker2
	userDn := "cn=linker2,dc=linkernetworks,dc=com"
	passwd := "password"

	//	baseDN := "dc=linkernetworks,dc=com"

	url := "ldap://" + host + ":" + port + "/"

	//init
	ldap, err := openldap.Initialize(url)
	if err != nil {
		log.Fatalf("Cannot connect to openldap.Reason:%v", err)
		return
	}
	log.Printf("Connected")

	//setup
	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)
	//	ldap.SetOption(openldap.)

	//connect
	err = ldap.Bind(userDn, passwd)
	if err != nil {
		log.Fatalf("Cannot bind to openldap.Reason:%v", err)
		return
	}
	log.Printf("Binded")

	defer func() {
		//		ldap.Unbind()
		ldap.Close()
	}()

	//do modify
	dn := userDn
	attrs := make(map[string][]string)

	//remove old password
	attrs["userPassword"] = []string{"password"}
	err = ldap.ModifyDel(dn, attrs)
	if err != nil {
		log.Printf("Fail to modify.Reason:%v", err)
		return
	}
	log.Print("Removed!")

	//add new password
	attrs["userPassword"] = []string{"newpassword4"}
	err = ldap.ModifyAdd(dn, attrs)
	if err != nil {
		log.Printf("Fail to modify.Reason:%v", err)
		return
	}
	log.Print("Added!")

	//modify
	attrs["userPassword"] = []string{"password"}
	err = ldap.Modify(dn, attrs)
	if err != nil {
		log.Printf("Fail to modify.Reason:%v", err)
		return
	}
	log.Printf("Modified!")
}
