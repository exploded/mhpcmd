This program allows command line control of a HiTechAstro Mount Hub Pro.

https://i.imgur.com/6VYarDZ.jpeg

<img src="https://i.imgur.com/6VYarDZ.jpeg" alt="MHP">

Commands can then be sent from N.I.N.A to the Mount Hub Pro using this tool.

Windows 64bit binary is supplied, but code can be recompiled for other operating systems. This has not been tested.

Focuser commands are not supported - use the MHP 32bit Ascom driver (file name: ASCOM_MHP_Setup.exe) and the ASCOM device hub to bridge to 64bit N.I.N.A.

<h2>Example commands:</h2>
<code>
Turn switch #1 on
mhpcmd switch -num=1 -state=1

Turn switch #1 off
mhpcmd.exe switch -num=1 -state=0

Turn switch #4 on
mhpcmd switch -num=4 -state=1

Turn on all switches:
mhpcmd switch -num=9 -state=1

Turn dew heater #1 to max (i.e. 100%)
mhpcmd.exe dew -num=1 -level=100

Turn dew heater #1 to 75%
mhpcmd.exe dew -num=1 -level=75

Turn dew heater #1 off (i.e. 0%)
mhpcmd.exe dew -num=1 -level=0
</code>

Default switch is #1, default action is On. This command will turn on switch #1:
mhpcmd switch

for dew heaters, default is dew heater #1, and default level is 100. This command will set dew heater #1 to 100%:
<code>
mhpcmd dew
</code>

To get help:
<code>
mhpcmd switch -h
mhpcmd dew -h
</code>