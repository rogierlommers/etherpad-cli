# etherpad-cli
A small CLI tool to interact with your [etherpad](https://github.com/ether/etherpad-lite) instance. List, add and delete pads.

## how to install?

### MacOS:

```
curl -L "https://github.com/rogierlommers/etherpad-cli/releases/download/1/etherpad-cli-macos" -o /usr/local/bin/etherpad-cli && chmod +x /usr/local/bin/etherpad-cli
```

### Linux

```
curl -L "https://github.com/rogierlommers/etherpad-cli/releases/download/1/etherpad-cli-linux" -o /usr/local/bin/etherpad-cli && chmod +x /usr/local/bin/etherpad-cli
```

### Windows
download [binary](https://github.com/rogierlommers/etherpad-cli/releases/download/1/etherpad-cli.exe) and put somewhere in path.

## first time use
Run etherpad-cli with `setconfig` parameter. It will ask for the etherpad `hostname` and `token`. Values will be stored in your home folder, `$HOME/.etherpad-cli.json`.

After this, you can start using etherpad-cli.

## options

```
List, add and delete pads on your etherpad instance

Usage:
  etherpad-cli [command]

Available Commands:
  add         add a new pad
  setconfig   write mandatory settings in config file
  delete      delete a pad
  help        Help about any command
  list        list al pads
```
