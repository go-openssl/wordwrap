package wordwrap_test

import (
	"testing"

	"github.com/go-openssl/wordwrap"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAtg/71nZw/qB/nWeaiYmO3xNWXr8jJWQ7Fv7RSeSRTItIP0bX
Nbbj7hPg5Fd0co3kFfJegg0S2faOSxsgSmwM7dz/BNfjWmnWjQZBR9bXGL3fWo6F
1ASxlYNcqKEo9oh0v4clZx9VjL3Lrq72GLpJU8LdUw2KvT9qP4iECPRizJmN4HF9
gRzz1LqkXn1Cd8yAwhQgjs70HTD8y2CK1+5TlzW3gTnl50uP5LbISlNo2/Kpgn4g
2nd2kUrHy4boejAFxIpNx0mV3xv9t2zr6lS/5r35nftmhoeAtJXMGOFd61OzkNla
8MwayNMKoteskjSeWhmdIntyfHu0QQjuo6ky0wIDAQABAoIBABIwCJWFIYaeXAFK
f0qsHkS0ttUiPQ2YNLr4W8oI9mbyJxaDKi6ohZ6cB8Rn5C8pc5aprr1JNw/lLJPe
MtycbLI6eJNwSbsXhaJy9ISLttM27KAsSrxCd8ca51/FMcVnHlqf0qG5pJ85uqZJ
vjX8LiUa+2c3nBuJwUB1j9MhLtmHdlVbuRgyGUdmJj6BcQtS1YtFTTWGMnpLq226
VGYM75KrVesD1VrfdWY3vLF3BO1EUy/IvWRQezox6Sm4zF5buFGD7IHVgOyWwxnk
bqUjwXxVBsOsDw6Fzg0jlu4sOPks90v+esqxdvr6dfijtwtn5KZBVQxyv7PdcTTc
vdHKsaECgYEA4F2LgtUsiFWn/GMRBwvwLPESIyMdWnAgfeM3NEdj6zfUUEZpFF/C
7o8jIoFLqa7mhlUzAOz7saFET3uENUzJ82oySckhodgZ8IlHE1KedZCCYkPFL2jz
50+cJ5UufEJFO1jLisYPeoVbGkBPQsomXfqdn+3M2s8VeLWpAi3VzfECgYEAz7uE
AuLepis4wV9rJ6UcTaJdgjesID7Umjk6p8OsPtygm6bTJm5cARuBT/LhaWAQxdYj
Y+x4xr+6oy68gi0zInQkPmukEnHswMnoq3JapnUabdaGCGe4GRBzcTjaDkyVPDz9
Ug5+x5Usl3aaualw+0YvpnyFI3tN3oxaskt/WQMCgYEAiYdNuQjKn1dB/WcMTPF4
a1Pp9jfUCleo0wGwGQ+Zo9k4/2vphV+dsXVz5/axVnWrQLSA6xRYw+1CXiYsSC+l
qttxr+DmCLraS6MaOjHuh8no4isAd6sxtpwJ8Al10R0eKt6nBY2ad1O/IDxDWYFo
Ozsf26R8abN+SduwmXFXGUECgYAH5U24Ol1SHZRzrSfKgvkXblN1jp4pP5ofHou/
Mq3KWeJ06Btge5Ndq2j32/h7Y95fVqtTsfpJO6Jhb3ZU0FkANz/la3v6A4CHN2Vz
ls4hQ5Q0lpHTofWaZkitBgcrwfduKbdLNifVeDMQsr5gzjLwKhPHlTYOSjKEgfs0
ibAWdwKBgQCSsbnCTNnXw4e6/b4b/1LiOC0ZtH0eZ1cyavumGuZJOHd5X0gH3StC
QUEKuiU+YcQDwlCPZK+a0ZQRLeq2HmlPU/sgXJKIp6oUWPNYTzama4YJPXGr30eb
1nybRhtcVUeD0QHeURgApSl/1/HAW+JT+jXtxuTEH36BKRyahlvB3g==
-----END RSA PRIVATE KEY-----`

const privateKeyStr = "MIIEpAIBAAKCAQEAtg/71nZw/qB/nWeaiYmO3xNWXr8jJWQ7Fv7RSeSRTItIP0bXNbbj7hPg5Fd0co3kFfJegg0S2faOSxsgSmwM7dz/BNfjWmnWjQZBR9bXGL3fWo6F1ASxlYNcqKEo9oh0v4clZx9VjL3Lrq72GLpJU8LdUw2KvT9qP4iECPRizJmN4HF9gRzz1LqkXn1Cd8yAwhQgjs70HTD8y2CK1+5TlzW3gTnl50uP5LbISlNo2/Kpgn4g2nd2kUrHy4boejAFxIpNx0mV3xv9t2zr6lS/5r35nftmhoeAtJXMGOFd61OzkNla8MwayNMKoteskjSeWhmdIntyfHu0QQjuo6ky0wIDAQABAoIBABIwCJWFIYaeXAFKf0qsHkS0ttUiPQ2YNLr4W8oI9mbyJxaDKi6ohZ6cB8Rn5C8pc5aprr1JNw/lLJPeMtycbLI6eJNwSbsXhaJy9ISLttM27KAsSrxCd8ca51/FMcVnHlqf0qG5pJ85uqZJvjX8LiUa+2c3nBuJwUB1j9MhLtmHdlVbuRgyGUdmJj6BcQtS1YtFTTWGMnpLq226VGYM75KrVesD1VrfdWY3vLF3BO1EUy/IvWRQezox6Sm4zF5buFGD7IHVgOyWwxnkbqUjwXxVBsOsDw6Fzg0jlu4sOPks90v+esqxdvr6dfijtwtn5KZBVQxyv7PdcTTcvdHKsaECgYEA4F2LgtUsiFWn/GMRBwvwLPESIyMdWnAgfeM3NEdj6zfUUEZpFF/C7o8jIoFLqa7mhlUzAOz7saFET3uENUzJ82oySckhodgZ8IlHE1KedZCCYkPFL2jz50+cJ5UufEJFO1jLisYPeoVbGkBPQsomXfqdn+3M2s8VeLWpAi3VzfECgYEAz7uEAuLepis4wV9rJ6UcTaJdgjesID7Umjk6p8OsPtygm6bTJm5cARuBT/LhaWAQxdYjY+x4xr+6oy68gi0zInQkPmukEnHswMnoq3JapnUabdaGCGe4GRBzcTjaDkyVPDz9Ug5+x5Usl3aaualw+0YvpnyFI3tN3oxaskt/WQMCgYEAiYdNuQjKn1dB/WcMTPF4a1Pp9jfUCleo0wGwGQ+Zo9k4/2vphV+dsXVz5/axVnWrQLSA6xRYw+1CXiYsSC+lqttxr+DmCLraS6MaOjHuh8no4isAd6sxtpwJ8Al10R0eKt6nBY2ad1O/IDxDWYFoOzsf26R8abN+SduwmXFXGUECgYAH5U24Ol1SHZRzrSfKgvkXblN1jp4pP5ofHou/Mq3KWeJ06Btge5Ndq2j32/h7Y95fVqtTsfpJO6Jhb3ZU0FkANz/la3v6A4CHN2Vzls4hQ5Q0lpHTofWaZkitBgcrwfduKbdLNifVeDMQsr5gzjLwKhPHlTYOSjKEgfs0ibAWdwKBgQCSsbnCTNnXw4e6/b4b/1LiOC0ZtH0eZ1cyavumGuZJOHd5X0gH3StCQUEKuiU+YcQDwlCPZK+a0ZQRLeq2HmlPU/sgXJKIp6oUWPNYTzama4YJPXGr30eb1nybRhtcVUeD0QHeURgApSl/1/HAW+JT+jXtxuTEH36BKRyahlvB3g=="

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtg/71nZw/qB/nWeaiYmO
3xNWXr8jJWQ7Fv7RSeSRTItIP0bXNbbj7hPg5Fd0co3kFfJegg0S2faOSxsgSmwM
7dz/BNfjWmnWjQZBR9bXGL3fWo6F1ASxlYNcqKEo9oh0v4clZx9VjL3Lrq72GLpJ
U8LdUw2KvT9qP4iECPRizJmN4HF9gRzz1LqkXn1Cd8yAwhQgjs70HTD8y2CK1+5T
lzW3gTnl50uP5LbISlNo2/Kpgn4g2nd2kUrHy4boejAFxIpNx0mV3xv9t2zr6lS/
5r35nftmhoeAtJXMGOFd61OzkNla8MwayNMKoteskjSeWhmdIntyfHu0QQjuo6ky
0wIDAQAB
-----END PUBLIC KEY-----`

const publicKeyStr = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtg/71nZw/qB/nWeaiYmO3xNWXr8jJWQ7Fv7RSeSRTItIP0bXNbbj7hPg5Fd0co3kFfJegg0S2faOSxsgSmwM7dz/BNfjWmnWjQZBR9bXGL3fWo6F1ASxlYNcqKEo9oh0v4clZx9VjL3Lrq72GLpJU8LdUw2KvT9qP4iECPRizJmN4HF9gRzz1LqkXn1Cd8yAwhQgjs70HTD8y2CK1+5TlzW3gTnl50uP5LbISlNo2/Kpgn4g2nd2kUrHy4boejAFxIpNx0mV3xv9t2zr6lS/5r35nftmhoeAtJXMGOFd61OzkNla8MwayNMKoteskjSeWhmdIntyfHu0QQjuo6ky0wIDAQAB"

func TestToPrivateKey(t *testing.T) {
	if wordwrap.New().ToPrivateKey(privateKeyStr) != privateKey {
		t.Error("ToPrivateKey failed")
		return
	}
}

func TestToPublicKey(t *testing.T) {
	if wordwrap.New().ToPublicKey(publicKeyStr) != publicKey {
		t.Error("ToPublicKey failed")
		return
	}
}
