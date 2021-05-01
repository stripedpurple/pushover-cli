# Pushover CLI
> **THIS IS NOT YET READY FOR USE BY THE PUBLIC**

`po` or pushover-cli is a simple cli to send messages to [pushover.net](https://pushover.net). 

## Installation
Currently the only way to install `po` is via go get
```shell script
go get github.com/viruscmd/pushover-cli

echo alias po=pushover-cli >> $HOME/.zshrc # Add alias for po to zsh
echo alias po=pushover-cli >> $HOME/.bashrc # Add alias for po to bash
```

## todos
- [ ] Add setup functionally to implement default config file
- [ ] Test highest priority functionality
- [ ] Rewrite descriptions to make sense
- [ ] Add attachment flag to send command
- [ ] Figure out better installation process  
- [ ] get added to brew  
- [ ] Setup CI/CD  
- [ ] prepare first release