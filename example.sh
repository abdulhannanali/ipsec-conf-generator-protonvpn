### Generates an Example Configuration. Requires a Linux AMD64 Binary
ARCH=amd64
OS=linux
NAME=ipsec-conf-generator-${ARCH}-${OS}

mkdir -p releases

./releases/ipsec-protonvpn-cli-linux-amd64 \
    -certPath /etc/ipsec.d/cacerts/protonvpn.der \
    -eapIdentity YOUR_IDENTITY \
    -writeToFile releases/protonvpn.conf
