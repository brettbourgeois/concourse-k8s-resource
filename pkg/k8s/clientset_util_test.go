package k8s

import (
	"github.com/brettbourgeois/concourse-k8s-resource/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const ca = `-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE5MTIzMTIzNTk1OFoXDTI5MTIyODIzNTk1OFowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJsh
iuGrvdI1YzuQM9yrREsRugS+JWBLYvlo4g+LKdt2YVHhGT6tyd4+FZxodq95kkn7
8DEpau8LLCMvXGTUPGSdWgbH70Jkts8mSXwoY1R+GFzLolKeEdgOcCn5HVAh0L/V
vZl60Iy+GI+3n/q0C+aSE+IQGebqr9/ripphJvufEw8Whdy3pQplhaZq8rBLNDr5
u9NOR+IQOvvf25nAbAvP+lH9CFcckb6nv2fV/Q8YybAA8/bFaNDo+mXQbGtJ9gkK
JBnSWD2n5LNX4G3yw2p/38UQTTCF4tSI/92UEiobbOqAP37mEQQOkcxQXhadtMJz
QWu8XGe483CtF0Nku0ECAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAAAJ/grFvkr52AVgF7VE1cvk4128
jFl+3owB9id+XFqeUfJIs2kQqG1IkSQhW4XRp2oZ+dwJ0Mf9852oCRtXNaJBhFcH
n1LWmTsqY/v3abryg3PezHDoxznCW2V0zHaeC/46oMWdoYV0Gruyk5VIuU0psmOI
K/O1O4BVhmmVXVEfpHriGgdcA+0Hbs3xIIFWluQ1BHWpq/EHA2QPsL3BfUsptEm4
iNE/Vaq7gOpHHSlHd11/lCzCb+Uwlegtjo66wiSoTGRMp/MOD7o81dSW3jhKtt+p
aeiNl12FygFE/YosX3nRGX1UoWeC/TmY/p6q+jeKoEXRA/RXAUmZYFXhQ2Q=
-----END CERTIFICATE-----
`
const clientCert = `-----BEGIN CERTIFICATE-----
MIICwjCCAaoCFDvXDRPyDqOGphLPhBWTSIF4svbGMA0GCSqGSIb3DQEBCwUAMBUx
EzARBgNVBAMTCmt1YmVybmV0ZXMwHhcNMjAwMTAxMDAwMDIzWhcNMjAxMjI2MDAw
MDIzWjAmMQ4wDAYDVQQDDAVhZG1pbjEUMBIGA1UECgwLbWFtZXpvdS1zcmUwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCrNsE6uf/mYviER93O7dmHVEje
7lmMSI8mBiFrGrbv0pnKy8sbTERc6LB2sONgoozeJAXbOjXn+N+S5EYanozFBigW
ovmBep6Nw/mkuDRxu3eDeDNkR/OrFdUM5+erA1d+NbdsglnIfRakCo90+hLr7suK
kPdJP4U5eJ5lnkqxY73hOoWdHMEcWCyCfqbQ+X7EQ17+rUErxMPoe3kasNPtAEkb
okSxzfHd5o1bs/6Pn3WwqSK3ssbh/tHHwP62rCg4DgNbReNlG29UrlIXtW21oAlb
H+VsoZORf0Gtv7wcmqtE2wGcz210ENhCqTZx0SgtQCoFMM60MvV+u2ZJ8i0DAgMB
AAEwDQYJKoZIhvcNAQELBQADggEBAJX9/Xa4VATH1KU83OPRLxQbB+M0er86VjwE
UGL4wHEma1XYDGDdwBTSFQXxa6dmvYtWZJ6bOWa0S+Io8WUmAQ+FE+LP7uohczt1
Ztul6ZE2vIXEt1cThRHjVpg1gJePqHiKazEclDbdhqnLSsd04NdL2vG9KZNS9Dwk
KBja3rmzJeyS1hAO3pu3s7bNulokAyR5a9VGZsK9xLDy/K9M6pApFXzRiVu7gNPL
cgBWXh98vKAj4F2jW5N2r92oHTRxY9vX7AcFV32RQdYD2cP+QWgISt2iqBBbgcWw
7sA/Mj2JCuqCIYV699CDo0Eu1acZZ5ygJn2yFgKQr/0LJR//M9w=
-----END CERTIFICATE-----
`
const clientKey = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCrNsE6uf/mYviE
R93O7dmHVEje7lmMSI8mBiFrGrbv0pnKy8sbTERc6LB2sONgoozeJAXbOjXn+N+S
5EYanozFBigWovmBep6Nw/mkuDRxu3eDeDNkR/OrFdUM5+erA1d+NbdsglnIfRak
Co90+hLr7suKkPdJP4U5eJ5lnkqxY73hOoWdHMEcWCyCfqbQ+X7EQ17+rUErxMPo
e3kasNPtAEkbokSxzfHd5o1bs/6Pn3WwqSK3ssbh/tHHwP62rCg4DgNbReNlG29U
rlIXtW21oAlbH+VsoZORf0Gtv7wcmqtE2wGcz210ENhCqTZx0SgtQCoFMM60MvV+
u2ZJ8i0DAgMBAAECggEAQx3806O0cEEeAOUXS5Yr7wQYaOPw0LBlBVfj49OeIRdi
2H/ZSAM2zWEeQ/kFuY0fQbnHXfBMz3ndUv0PikHbFyVZs74Bp0NFQnevtmXLkUYX
DL+jDc2y9L9jPGLwizaNJtmx5OSYg6KdrILDR+z8W+bJfbFkbx9qf2QMW/OYfj7i
PAXbdxF7/eRRk+rsu0trZ/sDQXMH8A3iTpGpoxk3zEzlGswCvvQrJWiupphDFoMt
gOwLgF/kzjSyBJZslLJ/+hWavSW3UiJMNNTcky7yvqj+p35X2Flgop3LEUjZdq19
tjdjSFrDHr3+odFcxixH3dkDWInJY5xhgZblRRErWQKBgQDS53OLK3ztrjzqMjMd
XP2n6DIYEQ0/cOi89ULt322TRPKbFtnI4mTqCEH4CaRS9sZWCsJs2OX1JL+mObZO
wfj1G1TjXTrGLZVPnOXae9IISJfFrthSQm32dFQ5+fahiVPUPIyuUhmLVaGFo6oS
PVhNXOYhkdroqP7iDooZ1AUbvwKBgQDP0rmkDTtQBgDgzET0LnWZdsDhZ4x7d2GR
n2Og781pmIMCNmVH5STU/91SLSxXcD/yrUoaVseWCluOlwkNLFyASDPpB7wCuPpB
9wH1c38iIQj3/LQnzz4i8DQhgoBnnb1vO/3Bu1mBxC88nwRB38hXf9bOuu4CNay8
YMX7R5OPvQKBgByBgBJ9bENL25vj8Ri06uv47FxoYZwDjNGNbOBt5IeVOB1SN1l5
kB45w4Dc/MLh6+jRR3oizuIVd3nmLwfyG841RYH9peYHXzkFgePH/Jl2Bl2HxmFH
7Uj0bDXx3S30O8ph7LnbCuzURCKl/mS8ueSq+8fpyObNgLXZNT1MdOxNAoGAO5tY
DXqSEYC3TcKo4FRW/H44Ei5t95elD2xk2esNwoSwxritUfKiHsmIRCKavjV+0e7r
+yP6uMkdu4cMXI/ltBGBegvy2+EMPlFHaYwH4dURynbbgTOKweCdQyM4CwAOLlJJ
lQBUSsjnN37wbKhvwND03nR1AYM9mQY0or7Dzw0CgYAIL2+8LI3u2zPzn2tPimN5
uFLU1i+k8TltqFs75zJxW//TOzaX6yxNKcrGwmimWNtEgSsJ+JicHaIQTzkEbJli
vefGyJne/451xEHW4jhhEiwa4rMw4fCwNZ1rX7l9+YZZqHQTs7xRV8dID3fcfAnq
Xkv0WCwnvAH0PCQP4ItqBw==
-----END PRIVATE KEY-----
`

const kubeconfig = `
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRFNU1USXpNVEl6TlRrMU9Gb1hEVEk1TVRJeU9ESXpOVGsxT0Zvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBSnNoCml1R3J2ZEkxWXp1UU05eXJSRXNSdWdTK0pXQkxZdmxvNGcrTEtkdDJZVkhoR1Q2dHlkNCtGWnhvZHE5NWtrbjcKOERFcGF1OExMQ012WEdUVVBHU2RXZ2JINzBKa3RzOG1TWHdvWTFSK0dGekxvbEtlRWRnT2NDbjVIVkFoMEwvVgp2Wmw2MEl5K0dJKzNuL3EwQythU0UrSVFHZWJxcjkvcmlwcGhKdnVmRXc4V2hkeTNwUXBsaGFacThyQkxORHI1CnU5Tk9SK0lRT3Z2ZjI1bkFiQXZQK2xIOUNGY2NrYjZudjJmVi9ROFl5YkFBOC9iRmFORG8rbVhRYkd0Sjlna0sKSkJuU1dEMm41TE5YNEczeXcycC8zOFVRVFRDRjR0U0kvOTJVRWlvYmJPcUFQMzdtRVFRT2tjeFFYaGFkdE1KegpRV3U4WEdlNDgzQ3RGME5rdTBFQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFBQUovZ3JGdmtyNTJBVmdGN1ZFMWN2azQxMjgKakZsKzNvd0I5aWQrWEZxZVVmSklzMmtRcUcxSWtTUWhXNFhScDJvWitkd0owTWY5ODUyb0NSdFhOYUpCaEZjSApuMUxXbVRzcVkvdjNhYnJ5ZzNQZXpIRG94em5DVzJWMHpIYWVDLzQ2b01XZG9ZVjBHcnV5azVWSXVVMHBzbU9JCksvTzFPNEJWaG1tVlhWRWZwSHJpR2dkY0ErMEhiczN4SUlGV2x1UTFCSFdwcS9FSEEyUVBzTDNCZlVzcHRFbTQKaU5FL1ZhcTdnT3BISFNsSGQxMS9sQ3pDYitVd2xlZ3RqbzY2d2lTb1RHUk1wL01PRDdvODFkU1czamhLdHQrcAphZWlObDEyRnlnRkUvWW9zWDNuUkdYMVVvV2VDL1RtWS9wNnEramVLb0VYUkEvUlhBVW1aWUZYaFEyUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
            server: https://k8s.test.local:6443
          name: local-k8s
      contexts:
        - context:
            cluster: local-k8s
            namespace: test
            user: tester
          name: local-k8s-tester
      current-context: local-k8s-tester
      preferences: {}
      users:
        - name: tester
          user:
            client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN3akNDQWFvQ0ZEdlhEUlB5RHFPR3BoTFBoQldUU0lGNHN2YkdNQTBHQ1NxR1NJYjNEUUVCQ3dVQU1CVXgKRXpBUkJnTlZCQU1UQ210MVltVnlibVYwWlhNd0hoY05NakF3TVRBeE1EQXdNREl6V2hjTk1qQXhNakkyTURBdwpNREl6V2pBbU1RNHdEQVlEVlFRRERBVmhaRzFwYmpFVU1CSUdBMVVFQ2d3TGJXRnRaWHB2ZFMxemNtVXdnZ0VpCk1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRQ3JOc0U2dWYvbVl2aUVSOTNPN2RtSFZFamUKN2xtTVNJOG1CaUZyR3JidjBwbkt5OHNiVEVSYzZMQjJzT05nb296ZUpBWGJPalhuK04rUzVFWWFub3pGQmlnVwpvdm1CZXA2TncvbWt1RFJ4dTNlRGVETmtSL09yRmRVTTUrZXJBMWQrTmJkc2dsbklmUmFrQ285MCtoTHI3c3VLCmtQZEpQNFU1ZUo1bG5rcXhZNzNoT29XZEhNRWNXQ3lDZnFiUStYN0VRMTcrclVFcnhNUG9lM2thc05QdEFFa2IKb2tTeHpmSGQ1bzFicy82UG4zV3dxU0szc3NiaC90SEh3UDYyckNnNERnTmJSZU5sRzI5VXJsSVh0VzIxb0FsYgpIK1Zzb1pPUmYwR3R2N3djbXF0RTJ3R2N6MjEwRU5oQ3FUWngwU2d0UUNvRk1NNjBNdlYrdTJaSjhpMERBZ01CCkFBRXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBSlg5L1hhNFZBVEgxS1U4M09QUkx4UWJCK00wZXI4NlZqd0UKVUdMNHdIRW1hMVhZREdEZHdCVFNGUVh4YTZkbXZZdFdaSjZiT1dhMFMrSW84V1VtQVErRkUrTFA3dW9oY3p0MQpadHVsNlpFMnZJWEV0MWNUaFJIalZwZzFnSmVQcUhpS2F6RWNsRGJkaHFuTFNzZDA0TmRMMnZHOUtaTlM5RHdrCktCamEzcm16SmV5UzFoQU8zcHUzczdiTnVsb2tBeVI1YTlWR1pzSzl4TER5L0s5TTZwQXBGWHpSaVZ1N2dOUEwKY2dCV1hoOTh2S0FqNEYyalc1TjJyOTJvSFRSeFk5dlg3QWNGVjMyUlFkWUQyY1ArUVdnSVN0MmlxQkJiZ2NXdwo3c0EvTWoySkN1cUNJWVY2OTlDRG8wRXUxYWNaWjV5Z0puMnlGZ0tRci8wTEpSLy9NOXc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
            client-key-data: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2QUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktZd2dnU2lBZ0VBQW9JQkFRQ3JOc0U2dWYvbVl2aUUKUjkzTzdkbUhWRWplN2xtTVNJOG1CaUZyR3JidjBwbkt5OHNiVEVSYzZMQjJzT05nb296ZUpBWGJPalhuK04rUwo1RVlhbm96RkJpZ1dvdm1CZXA2TncvbWt1RFJ4dTNlRGVETmtSL09yRmRVTTUrZXJBMWQrTmJkc2dsbklmUmFrCkNvOTAraExyN3N1S2tQZEpQNFU1ZUo1bG5rcXhZNzNoT29XZEhNRWNXQ3lDZnFiUStYN0VRMTcrclVFcnhNUG8KZTNrYXNOUHRBRWtib2tTeHpmSGQ1bzFicy82UG4zV3dxU0szc3NiaC90SEh3UDYyckNnNERnTmJSZU5sRzI5VQpybElYdFcyMW9BbGJIK1Zzb1pPUmYwR3R2N3djbXF0RTJ3R2N6MjEwRU5oQ3FUWngwU2d0UUNvRk1NNjBNdlYrCnUyWko4aTBEQWdNQkFBRUNnZ0VBUXgzODA2TzBjRUVlQU9VWFM1WXI3d1FZYU9QdzBMQmxCVmZqNDlPZUlSZGkKMkgvWlNBTTJ6V0VlUS9rRnVZMGZRYm5IWGZCTXozbmRVdjBQaWtIYkZ5VlpzNzRCcDBORlFuZXZ0bVhMa1VZWApETCtqRGMyeTlMOWpQR0x3aXphTkp0bXg1T1NZZzZLZHJJTERSK3o4VytiSmZiRmtieDlxZjJRTVcvT1lmajdpClBBWGJkeEY3L2VSUmsrcnN1MHRyWi9zRFFYTUg4QTNpVHBHcG94azN6RXpsR3N3Q3Z2UXJKV2l1cHBoREZvTXQKZ093TGdGL2t6alN5Qkpac2xMSi8raFdhdlNXM1VpSk1OTlRja3k3eXZxaitwMzVYMkZsZ29wM0xFVWpaZHExOQp0amRqU0ZyREhyMytvZEZjeGl4SDNka0RXSW5KWTV4aGdaYmxSUkVyV1FLQmdRRFM1M09MSzN6dHJqenFNak1kClhQMm42RElZRVEwL2NPaTg5VUx0MzIyVFJQS2JGdG5JNG1UcUNFSDRDYVJTOXNaV0NzSnMyT1gxSkwrbU9iWk8Kd2ZqMUcxVGpYVHJHTFpWUG5PWGFlOUlJU0pmRnJ0aFNRbTMyZEZRNStmYWhpVlBVUEl5dVVobUxWYUdGbzZvUwpQVmhOWE9ZaGtkcm9xUDdpRG9vWjFBVWJ2d0tCZ1FEUDBybWtEVHRRQmdEZ3pFVDBMbldaZHNEaFo0eDdkMkdSCm4yT2c3ODFwbUlNQ05tVkg1U1RVLzkxU0xTeFhjRC95clVvYVZzZVdDbHVPbHdrTkxGeUFTRFBwQjd3Q3VQcEIKOXdIMWMzOGlJUWozL0xRbnp6NGk4RFFoZ29Cbm5iMXZPLzNCdTFtQnhDODhud1JCMzhoWGY5Yk91dTRDTmF5OApZTVg3UjVPUHZRS0JnQnlCZ0JKOWJFTkwyNXZqOFJpMDZ1djQ3RnhvWVp3RGpOR05iT0J0NUllVk9CMVNOMWw1CmtCNDV3NERjL01MaDYralJSM29penVJVmQzbm1Md2Z5Rzg0MVJZSDlwZVlIWHprRmdlUEgvSmwyQmwySHhtRkgKN1VqMGJEWHgzUzMwTzhwaDdMbmJDdXpVUkNLbC9tUzh1ZVNxKzhmcHlPYk5nTFhaTlQxTWRPeE5Bb0dBTzV0WQpEWHFTRVlDM1RjS280RlJXL0g0NEVpNXQ5NWVsRDJ4azJlc053b1N3eHJpdFVmS2lIc21JUkNLYXZqViswZTdyCit5UDZ1TWtkdTRjTVhJL2x0QkdCZWd2eTIrRU1QbEZIYVl3SDRkVVJ5bmJiZ1RPS3dlQ2RReU00Q3dBT0xsSkoKbFFCVVNzam5OMzd3YktodndORDAzblIxQVlNOW1RWTBvcjdEencwQ2dZQUlMMis4TEkzdTJ6UHpuMnRQaW1ONQp1RkxVMWkrazhUbHRxRnM3NXpKeFcvL1RPemFYNnl4Tktjckd3bWltV050RWdTc0orSmljSGFJUVR6a0ViSmxpCnZlZkd5Sm5lLzQ1MXhFSFc0amhoRWl3YTRyTXc0ZkN3Tloxclg3bDkrWVpacUhRVHM3eFJWOGRJRDNmY2ZBbnEKWGt2MFdDd252QUgwUENRUDRJdHFCdz09Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
`

const url = "https://k8s.test.local:6443"

func TestBuildClientSetUsingClientCert(t *testing.T) {
	assert := assert.New(t)

	source := models.Source{
		ApiServerUrl: url,
		ApiServerCA:  ca,
		ClientCert:   clientCert,
		ClientKey:    clientKey,
		Namespace:    "test",
	}

	_, config := NewClientSet(&source)

	namespace, _, err := config.Namespace()
	if err != nil {
		assert.Fail("cannot get namespace")
	}
	assert.Equal(namespace, "test")
	rawConfig, err := config.RawConfig()
	if err != nil {
		assert.FailNow("cannot get raw config")
		return
	}
	cluster := rawConfig.Clusters["default"]
	if assert.NotNil(cluster) {
		assert.False(cluster.InsecureSkipTLSVerify)
		assert.Equal(cluster.CertificateAuthorityData, []byte(ca))
		assert.Equal(cluster.Server, url)
	}
	auth := rawConfig.AuthInfos["concourse"]
	if assert.NotNil(auth) {
		assert.Equal(auth.ClientCertificateData, []byte(clientCert))
		assert.Equal(auth.ClientKeyData, []byte(clientKey))
	}
}

func TestBuildClientSetUsingKubeconfig(t *testing.T) {
	assert := assert.New(t)

	source := models.Source{
		Kubeconfig: kubeconfig,
	}

	_, config := NewClientSet(&source)

	namespace, _, err := config.Namespace()
	if err != nil {
		assert.Fail("cannot get namespace")
	}
	assert.Equal(namespace, "test")
	rawConfig, err := config.RawConfig()
	if err != nil {
		assert.FailNow("cannot get raw config")
		return
	}
	cluster := rawConfig.Clusters["local-k8s"]
	if assert.NotNil(cluster) {
		assert.False(cluster.InsecureSkipTLSVerify)
		assert.Equal(cluster.CertificateAuthorityData, []byte(ca))
		assert.Equal(cluster.Server, url)
	}
	auth := rawConfig.AuthInfos["tester"]
	if assert.NotNil(auth) {
		assert.Equal(auth.ClientCertificateData, []byte(clientCert))
		assert.Equal(auth.ClientKeyData, []byte(clientKey))
	}
}

func TestBuildClientSetUsingToken(t *testing.T) {
	assert := assert.New(t)

	source := models.Source{
		ApiServerUrl: url,
		ApiServerCA:  ca,
		ClientToken:  "token",
		Namespace:    "test",
	}

	_, config := NewClientSet(&source)

	namespace, _, err := config.Namespace()
	if err != nil {
		assert.Fail("cannot get namespace")
	}
	assert.Equal(namespace, "test")
	rawConfig, err := config.RawConfig()
	if err != nil {
		assert.FailNow("cannot get raw config")
		return
	}
	cluster := rawConfig.Clusters["default"]
	if assert.NotNil(cluster) {
		assert.False(cluster.InsecureSkipTLSVerify)
		assert.Equal(cluster.CertificateAuthorityData, []byte(ca))
		assert.Equal(cluster.Server, url)
	}
	auth := rawConfig.AuthInfos["concourse"]
	if assert.NotNil(auth) {
		assert.Equal(auth.Token, "token")
	}
}
