# Rogue: CSGO Hack
A mega basic Counter-Strike: Go radar and near to mid range wallhack that works even after the update in 2018. The idea is to give the user more map awareness without tipping that they are hacking to make it much harder to detect. 

This is a working hack but I'm developing it primarily as a way to practice interacting with process memory via the Linux `/proc/*` subsystem creating a hack with ideally no dependencies at all beyond Go standard libraries. The design is intentionally subtle, will be be toggle-able, YAML configured, and very light-weight (resource-wise). Being subtle will enable hackers to go undetected longer, perhaps it should even watch match details and only turn on when its needed, like when you are the last player left or similar circumstances.  


### Install
Start by installing the dependencies, this obviously requires Linux as it relies on `/proc/*` subsystem. Switch out dependencies with the package names in your favorite distro of Linux. 

```
sudo apt-get install xorg-dev libxnvctrl-dev make build-essential git
```

After the dependencies are installed then just use make:

```
make
```

And run:

```
sudo ./rogue-cg
```


### Future Development
Still needs more of the bulk cut out of it, switched it to a much simpler `make` build system over the pre-existing `cmake` option. One can even use `make clean` when developing to simplify build and rebuilding. 

It would be cool to convert this to Go.

**Potential Libraries**
Need a library for interacting with `/proc/*` linux subsystem, here is a list of potential options:

  * https://github.com/prometheus/procfs

  * https://github.com/c9s/goprocinfo

  * https://github.com/ncabatoff/process-exporter

  * https://github.com/intelsdi-x/snap-plugin-collector-processes/blob/master/processes/procstat.go

  
**Most interesting candidates**
The problem with the above is that they are pretty bloated, so it would be nice to find one that is just right, then we can include it so that there are no outside dependencies beyond standard libraries. 

**Improving ticks**

```
  // Timing channels for polling interval
  collectionInterval := time.NewTicker(time.Second * time.Duration(config.TickerInterval))
```
https://github.com/bellycard/procd/blob/master/procd.go

