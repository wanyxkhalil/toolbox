# toolbox
toolbox: mkpasswd, https-expired, mysql-to-gostruct ...


## Install

```shell
go install github.com/wanyxkhalil/toolbox@latest
toolbox -h
```

### Completion

If you use zsh, add this to ~/.zshrc & source.
Also support bash, fish, powershell 
```shell
if [ $commands[toolbox] ]; then
	source <(toolbox completion zsh)
	compdef _toolbox toolbox
fi
```

## Usage

### mkpasswd

Generate password, Inspired by [gaiustech/MkPasswd](https://github.com/gaiustech/MkPasswd).
```shell
$ toolbox mkpasswd -h # help message
A tool for generating random passwords

Usage:
  toolbox mkpasswd [flags]

Flags:
  -d, --digit uint     Number of digits (default 2)
  -h, --help           help for mkpasswd
  -l, --length uint    Length in chars (default 9)
  -c, --lower uint     Number of lowercase chars (default 2)
  -s, --special uint   Number of special chars
  -C, --upper uint     Number of uppercase chars (default 2)
```

Sample
```shell
toolbox mkpasswd -l 17 # length is 17
toolbox mkpasswd -l 17 -C 4 -d 4 -s 3 # length is 17, include 4 upper char, 4 digit, 3 special char, 6 lower char
```

### https-expired

Show the cert expiration time. Just like
```shell
alias ,https-expired='function _as() {echo | openssl s_client -servername $1 -connect $1:443 2>/dev/null | openssl x509 -noout -dates;};_as'
```

Sample
```shell
toolbox https-expired github.com
```