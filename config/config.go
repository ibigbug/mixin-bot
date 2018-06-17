package config

import "os"

// GetEnv gets the environment variable
func GetEnv(key, dft string) (v string) {
	if v = os.Getenv(key); v == "" {
		v = dft
	}
	return
}

var (
	ClientID     = GetEnv("ClientId", "62278bf5-04e0-4499-89e6-51f085be0b06")
	ClientSecret = GetEnv("ClientSecret", "2e9e1dbd5277f89cf853e02478d26657fbe33617bc872547291ec7bef9c01081")
	PinCode      = GetEnv("PinCode", "264659")
	SessionID    = GetEnv("SessionId", "1ee09259-4c46-4d7e-88af-047729aeede2")
	PinTOken     = GetEnv("PinToken", "eEWAR6w6aiaNGTqV7kz3HZxYYGRfC3aIZpKlS8nXueOwqmQcU7F/oiSqXc6C7x5tTh1xdG5TFlhJtfGMJmscPhRcKoqbPqGFKOsHQyTuvIs0L1nOoTPZ07lyWnEm3R5qpasTfwXLFNexaQGKC+xzyIlOeqjvkCPvSWDuSMUMmos=")
	PrivateKey   = GetEnv("PrivateKey", `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDPVlaUm64012ttmhOXoPTdMPTrsQLCYi+mupI/LMQv+N5KVsmn
25Fa9ep3+GaqY8eKxTfG/qdgFy7taY4pA04aLZ147lffAt/DtWc0BmO6f2PYa4Yr
czmeiKDddQbqoUSfUqM3bgLH8bC9dwoYP47SR8vJ5Pm54ABFGljSKf7zAQIDAQAB
AoGAGrx0b7dfkHrS5JBAxIXB7Z/5hUcPJwfIQ0S9xR29sr3x1D46xMWAie06Lfw/
KOuy+IcT+TovVSnIKF87MLi3YldHjZ4A5Zi9GjxZ6JJDN/WOR0zz7YMhbPFC7BvG
LjU9RfpeSamXws1F9EswwpffnXhLC3Y2ieysfKeRuhCpqxECQQD0hbpB++9KmF9T
lFqkMfda/PsEV0mR9qHh6JO9l0z5VJe+jne7Ei05SlKfxovqWSGBsHBChqtkeEhT
6ZMpusfdAkEA2RHHxj5SmtgJ2uXA/jx6oRufigrpTL2T2Xap51mAxlYjZclR9CVz
0vGIV1aEh06M0dg62Y0SztKOmbYr3/PXdQJAWmN8oJuiry54Posenoeh6k6N8+LY
XU6QNNQmmPTvDVPJ7DT56XskoILq3Akm+3ALRMeWF+F0ReV4xxwWvFxBOQJAUIq9
YnxrcBLrDTKkunC8W50BRFoqGJRbEJxsDHB0TKNXfQQHCZz/7ew615U9lWr66z3d
EC29JcqQqM4kV3OFpQJBANFbgZRKJoIAJRKpgqr7MJa6oYKY2q+GV9s5/3bSfSNm
0ft2s17sX0KN92V5NIpNZcVu5lBw0nk8uk/Awxo2j/s=
-----END RSA PRIVATE KEY-----
		`)
)
