This program allows command line control of the origina; version of the HiTechAstro Mount Hub Pro. It looks like this:

<img src="https://i.imgur.com/6VYarDZ.jpeg" alt="MHP">

Commands can then be sent from <a href="https://nighttime-imaging.eu/">N.I.N.A.</a> to the Mount Hub Pro using this tool.

Windows 64bit binary is supplied. Code can be recompiled for other operating systems, but has not been tested.

Focuser commands are not supported - use the MHP 32bit Ascom driver (file name: ASCOM_MHP_Setup.exe) and the ASCOM device hub to bridge to 64bit N.I.N.A.

<h2>Example commands:</h2>

```
Turn switch #1 on
mhpcmd switch -num=1 -state=1

Turn switch #1 off
mhpcmd switch -num=1 -state=0

Turn switch #4 on
mhpcmd switch -num=4 -state=1

Turn on all switches:
mhpcmd switch -num=9 -state=1

Turn dew heater #1 to max (i.e. 100%)
mhpcmd dew -num=1 -level=100

Turn dew heater #1 to 75%
mhpcmd dew -num=1 -level=75

Turn dew heater #1 off (i.e. 0%)
mhpcmd dew -num=1 -level=0
```

Default switch is #1, default action is On. This command will turn on switch #1:

```
mhpcmd switch
```

for dew heaters, default is dew heater #1, and default level is 100. This command will set dew heater #1 to 100%:

```
mhpcmd dew
```

To get help:

```
mhpcmd switch -h
mhpcmd dew -h
```
