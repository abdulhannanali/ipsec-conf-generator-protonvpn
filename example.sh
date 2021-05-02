### Generates an Example Configuration. Requires a Linux AMD64 Binary
mkdir -p releases

./releases/ipsec-protonvpn-cli-linux-amd64 \
    -certPath /etc/ipsec.d/cacerts/protonvpn.der \
    -eapIdentity YOUR_IDENTITY \
    -writeToFile releases/protonvpn.conf
