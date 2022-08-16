//go:build !windows && go1.13
// +build !windows,go1.13

package mssql

import (
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/keytab"
)

func getAuthN(user, password, serverSPN, workstation string, Kerberos map[string]interface{}) (auth auth, authOk bool) {
	if Kerberos != nil && Kerberos["Config"] != nil {
		auth, authOk = getKRB5Auth(user, serverSPN, Kerberos["Config"].(*config.Config), Kerberos["Keytab"].(*keytab.Keytab), Kerberos["Cache"].(*credentials.CCache))
	} else {
		auth, authOk = getAuth(user, password, serverSPN, workstation)
	}
	return
}
