package ldap_pswd

import (
	"log"

	"github.com/mqu/openldap"
)

func ChangePasswd(baseDn string, userId string, oldPasswd string, newPasswd string) {
	//connect and bind
	ldap, err := connectLdap(baseDn, userId, oldPasswd)
	if err != nil {
		log.Printf("%v", err)
	}
	defer func() {
		if ldap != nil {
			//ldap.Unbind()
			ldap.Close()
		}
	}()

	//modify userPassword
	dn := "cn=" + userId + "," + baseDn
	attrs := make(map[string][]string)

	attrs["userPassword"] = []string{newPasswd}
	err = ldap.Modify(dn, attrs)
	if err != nil {
		log.Printf("Fail to modify.Reason:%v", err)
		return
	}
	log.Printf("Modified!")
}

func connectLdap(baseDN string, userId string, passwd string) (*openldap.Ldap, error) {
	//TODO caution,hard code
	//change the host and port
	host := "192.168.10.194"
	port := "8389"

	//init
	url := "ldap://" + host + ":" + port + "/"
	ldap, err := openldap.Initialize(url)
	if err != nil {
		log.Fatalf("Cannot connect to openldap.Reason:%v", err)
		return nil, err
	}
	log.Printf("Connected")

	//setup
	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	//connect
	userDn := "cn=" + userId + "," + baseDN
	err = ldap.Bind(userDn, passwd)
	if err != nil {
		log.Fatalf("Cannot bind to openldap.Reason:%v", err)
		return nil, err
	}
	log.Printf("Binded")
	return ldap, err
}
