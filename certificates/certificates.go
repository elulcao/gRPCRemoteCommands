package certificates

import (
	"embed"
	"io/fs"
	"strings"
)

var (
	//go:embed assets/ca-cert.pem
	cacertByte []byte
	//go:embed assets/ca-key.pem
	cakeyByte []byte
	//go:embed assets/client-cert.pem
	clientcertByte []byte
	//go:embed assets/client-key.pem
	clientkeyByte []byte
	//go:embed assets/client-req.pem
	clientreqByte []byte
	//go:embed assets/server-cert.pem
	servercertByte []byte
	//go:embed assets/server-key.pem
	serverkeyByte []byte
	//go:embed assets/server-req.pem
	serverreqByte []byte
	//go:embed assets/ca-cert.pem
	cacertString string
	//go:embed assets/ca-key.pem
	cakeyString string
	//go:embed assets/client-cert.pem
	clientcertString string
	//go:embed assets/client-key.pem
	clientkeyString string
	//go:embed assets/client-req.pem
	clientreqString string
	//go:embed assets/server-cert.pem
	servercertString string
	//go:embed assets/server-key.pem
	serverkeyString string
	//go:embed assets/server-req.pem
	serverreqString string
	//go:embed assets/ca-cert.pem
	cacert embed.FS
	//go:embed assets/ca-key.pem
	cakey embed.FS
	//go:embed assets/client-cert.pem
	clientcert embed.FS
	//go:embed assets/client-key.pem
	clientkey embed.FS
	//go:embed assets/client-req.pem
	clientreq embed.FS
	//go:embed assets/server-cert.pem
	servercert embed.FS
	//go:embed assets/server-key.pem
	serverkey embed.FS
	//go:embed assets/server-req.pem
	serverreq embed.FS
	//go:embed assets
	allcerts embed.FS
)

// GetCertByte read the embeded file and return its content in []byte, valid for go >= 1.6
func GetCertByte(name string) []byte {
	var ans []byte

	switch strings.ToLower(name) {
	case "ca-cert":
		ans = cacertByte
	case "ca-key":
		ans = cakeyByte
	case "client-cert":
		ans = clientcertByte
	case "client-key":
		ans = clientkeyByte
	case "client-req":
		ans = clientreqByte
	case "server-cert":
		ans = servercertByte
	case "server-key":
		ans = serverkeyByte
	case "server-req":
		ans = serverreqByte
	}

	return ans
}

// GetCertString read the embeded file and return its content in string, valid for go >= 1.6
func GetCertString(name string) string {
	var ans string

	switch strings.ToLower(name) {
	case "ca-cert":
		ans = cacertString
	case "ca-key":
		ans = cakeyString
	case "client-cert":
		ans = clientcertString
	case "client-key":
		ans = clientkeyString
	case "client-req":
		ans = clientreqString
	case "server-cert":
		ans = servercertString
	case "server-key":
		ans = serverkeyString
	case "server-req":
		ans = serverreqString
	}

	return ans
}

// GetCertFile read the embeded file and return its content in fs.File, valid for go >= 1.6
func GetCertFile(name string) (fs.File, error) {
	var ans fs.File
	var err error

	switch strings.ToLower(name) {
	case "ca-cert":
		ans, err = cacert.Open("assets/ca-cert.pem")
		if err != nil {
			return nil, err
		}
	case "ca-key":
		ans, err = cakey.Open("assets/ca-key.pem")
		if err != nil {
			return nil, err
		}
	case "client-cert":
		ans, err = clientcert.Open("assets/client-cert.pem")
		if err != nil {
			return nil, err
		}
	case "client-key":
		ans, err = clientkey.Open("assets/client-key.pem")
		if err != nil {
			return nil, err
		}
	case "client-req":
		ans, err = clientreq.Open("assets/client-req.pem")
		if err != nil {
			return nil, err
		}
	case "server-cert":
		ans, err = servercert.Open("assets/server-cert.pem")
		if err != nil {
			return nil, err
		}
	case "server-key":
		ans, err = serverkey.Open("assets/server-key.pem")
		if err != nil {
			return nil, err
		}
	case "server-req":
		ans, err = serverreq.Open("assets/server-req.pem")
		if err != nil {
			return nil, err
		}
	}
	defer ans.Close()

	return ans, nil
}

// GetAllCertNames return all the template names from the embeded FS
func GetAllCertNames() []string {
	s := make([]string, 0)
	m, _ := allcerts.ReadDir(".") // Ignore the error, the FS is embebed

	for _, f := range m {
		s = append(s, f.Name())
	}

	return s
}
