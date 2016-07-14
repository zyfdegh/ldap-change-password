package ldapasswd

import (
	"log"

	"github.com/mqu/openldap"
)

//This function is used to change password in openldap.
//Params:
//	host: Openldap IP or hostname.
//	port: Openldap port.
//	dn:	similar to user ID.
//	oldPasswd: Current password.
//	newPasswd: New password.
//The argument 'dn' looks like this,
//	dn := "uid=linker@linkernetworks.com,dc=linkernetworks,dc=com"
//	dn := "cn=admin,dc=linkernetworks,dc=com"
//Returns:
//	err: error
func ChangePasswd(host string, port string, dn string, oldPasswd string,
	newPasswd string) (err error) {
	//connect and bind
	ldap, err := connectOpenldap(host, port, dn, oldPasswd)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	defer func() {
		if ldap != nil {
			//ldap.Unbind()
			ldap.Close()
		}
	}()

	//modify userPassword
	attrs := make(map[string][]string)
	attrs["userPassword"] = []string{newPasswd}
	err = ldap.Modify(dn, attrs)
	if err != nil {
		log.Printf("Fail to modify.Reason:%v", err)
		return
	}
	log.Printf("Modified!")
	return
}

func connectOpenldap(host string, port string, dn string,
	passwd string) (*openldap.Ldap, error) {
	//init
	url := "ldap://" + host + ":" + port + "/"
	ldap, err := openldap.Initialize(url)
	if err != nil {
		log.Printf("Cannot connect to openldap.Reason:%v", err)
		return nil, err
	}
	log.Printf("Connected")

	//setup
	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	//connect
	err = ldap.Bind(dn, passwd)
	if err != nil {
		log.Printf("Cannot bind to openldap.Reason:%v", err)
		return nil, err
	}
	log.Printf("Binded")
	return ldap, err
}
