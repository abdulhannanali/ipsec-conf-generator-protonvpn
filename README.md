# ipsec-conf-generator-protonvpn


This tool is used for generation for IPSec configuration for ProtonVPN, which makes it similar to the usecase of downloading a directory of all OpenVPN files for ProtonVPN. Can expose the full power of ProtonVPN ecosystem at your hands with an alternative protocol

## But why?

So it turns out OpenVPN is blocked by default, once I bought the subscription I was 
baffled to see that my money was being wasted since their official Linux client didn't really work as it only seems to support OpenVPN, however the ProtonVPN Android client seemed to work just fine, on a deeper inspection I discovered that ProtonVPN uses IPSec/IKEv2 for connection and follow some very unclear tutorial to do the same on Linux machine.

This worked out very well and I was able to exploit the full power of ProtonVPN global server network. However, I don't really like using just one configuration and modification of default `ipsec.conf` is a pain as there are approximate 1400+ servers at the time of writing this doc, and for the biggest for buck, making use of different servers should be important. Hence, the need for this tool.

It helps you simply generate an ipsec.conf file with all the ProtonVPN server configuration, making it similar to the usecase where you download a zip directory of all the open vpn configuration and can just pick one or more to connect. 


It relies on legacy ipsec.conf file, and although I have now come to know that Strongswan now supports a new configuration file and also a `vici` protocol, I might implement that building on this tool if there's more demand.

## Architecture/OS Support

Checkout [releases](./releases) for more information about Architecture/OS Support. All three major desktop platforms: Ubuntu, MacOS, Windows are supported.

## Usage


### Automated 

More can be learnt about the usage by using the command command. You'd need a client and related credentails for it to work, please refer to this tutorial for instructions on how to do that, and later use this tool for an easy way to generate full configuration.

`./ipsec-protovpn-cli-linux-amd64 -help`

For an example checkout [example.sh](./example.sh)

### Example config file

An example config file is pushed each time a new release is done. You can check it out in [releases](/releases), it is named protonvpn.conf.



### Feature Requests/Feedback

If you have any future request/feedback, please go to Issues tab and open an issue. I'd love to hear from you.

### Contribution

Contributions and Bug Fixes are more than welcome. I'd be creating more issues about things that new contributors can work on.


## License

See [LICENSE](./LICENSE) for more details.