package mssql

func getAuthN(user, password, serverSPN, workstation string, _ map[string]interface{}) (auth auth, authOk bool) {
	auth, authOk = getAuth(user, password, serverSPN, workstation)
	return
}
