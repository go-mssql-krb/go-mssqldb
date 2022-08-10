//go:build windows
// +build windows

package mssql

func getAuthN(p msdsn.Config) (auth auth, authOk bool) {
	auth, authOk = getAuth(p.User, p.Password, p.ServerSPN, p.Workstation)
	return
}
