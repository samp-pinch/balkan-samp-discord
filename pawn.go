{{$n := .StrippedMsg}}
{{ if eq (lower $n) (lower "OnIncomingConnection")}}
**- OnIncomingConnection(playerid, ip_address[], port) -**
> **playerid** - id igrača koji zahteva konekciju sa serverom.
> **ip_address** - ip adresa igrača koji zahteva konekciju sa serverom.
> **port** - port igrača koji zahteva konekciju sa serverom.
Poziva se pre OnPlayerConnect-a ( kada igrač vidi poruku Connecting **IP** ).

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnIncomingConnection

{{ else if eq (lower $n) (lower "OnPlayerConnect")}}
**- OnPlayerConnect(playerid) -**
> **playerid** - id igrača koji se konektuje.
Poziva se kada se igrač konektuje na server.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerConnect

{{ else if eq (lower $n) (lower "OnPlayerFinishedDownloading")}}
**- OnPlayerFinishedDownloading(playerid, virtualworld) -**
> **playerid** - id igrača koji je završio preuzimanje modela.
> **virtualworld** - id vw-a u kom je igrač koji je završio preuzimanje modela.
Poziva se kada se igrač preuzme modele u cache folder (0.3DL).

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerFinishedDownloading

{{else if eq (lower $n) (lower "OnPlayerDisconnect")}}
**- OnPlayerDisconnect(playerid, reason) -**
> **playerid** - id igrača koji se diskonektovao.
> **reason**:
```0 - Timeout/Crash
1 - Quit
2 - Kick/Ban```
Poziva se kada se igrač diskonektuje sa servera.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerDisconnect

{{else if eq (lower $n) (lower "OnGameModeInit")}}
**- OnGameModeInit() -**
> Ovaj callback nema parametre.
Poziva se kada se GameMode pokrene.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnGameModeInit

{{else if eq (lower $n) (lower "OnGameModeExit")}}
**- OnGameModeExit() -**
> Ovaj callback nema parametre.
Poziva se kada se GameMode ugasi (rcon exit).

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnGameModeExit

{{else if eq (lower $n) (lower "OnPlayerRequestSpawn")}}
**- OnPlayerRequestSpawn(playerid) -**
> **playerid** - id igrača koji zahteva spawn.
Poziva se kada igrač pritisnite 'Spawn' dugme (Class Selection).

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerRequestSpawn

{{else if eq (lower $n) (lower "OnPlayerSpawn")}}
**- OnPlayerSpawn(playerid) -**
> **playerid** - id igrača koji je spawnovan.
Poziva se kada se igrač spawnuje.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerSpawn

{{else if eq (lower $n) (lower "OnPlayerDeath")}}
**- OnPlayerDeath(playerid, killerid, reason) -**
> **playerid** - id igrača koji je umro.
> **killerid** - id igrača koji je izvršio ubistvo (može biti INVALID_PLAYER_ID).
> **reason** - oružje kojim je igrač ubijen (wiki/Weapons za IDove).
Poziva se kada igrač umre.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerDeath

{{else if eq (lower $n) (lower "OnDialogResponse")}}
**- OnDialogResponse(playerid, dialogid, response, listitem, inputtext[]) -**
> **playerid** - id igrača koji odgovara na dialog.
> **dialogid** - id dialoga na koji igrač odgovara.
> **response** - id dugmeta/gumba koje je igrač pritisnuo (**1** levo : **2** desno).
> **listitem** - id itema koji je igrač izabrao iz liste (kreće od **0**, DIALOG_STYLE_LIST).
> **inputtext** - text koji je igrač uneo u dialog (DIALOG_STYLE_INPUT).
Poziva se kada igrač odgovori na dialog.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnDialogResponse

{{else if eq (lower $n) (lower "OnPlayerText")}}
**- OnPlayerText(playerid, text[]) -**
> **playerid** - id igrača koji je poslao poruku.
> **text[]** - text poruke koju je igrač poslao.
Poziva se kada igrač napiše poruku u chat.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerText

{{else if eq (lower $n) (lower "OnPlayerCommandText")}}
**- OnPlayerCommandText(playerid, cmdtext[]) -**
> **playerid** - id igrača koji je poslao komandu.
> **cmdtext[]** - komanda koju je igrač poslao.
Poziva se kada igrač napiše komandu u chat.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerCommandText

{{else if eq (lower $n) (lower "OnRconCommand")}}
**- OnRconCommand(cmd[]) -**
> **cmd[]** - komanda koja je izvršena u konzolu servera.
Poziva se kada rcon komanda bude procesuirana u konzolu servera.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnRconCommand

{{else if eq (lower $n) (lower "OnRconLoginAttempt")}}
**- OnRconLoginAttempt(ip[], password[], success) -**
> **ip[]** - ip igrača koji je pokušao da se poveže na rcon.
> **password[]** - lozinka sa kojom je igrač pokušao da se poveže.
> **success** - Da li je lozinka tačna (**0** ako je pogrešna, **1** ako je tačna).
Poziva se igrač pokuša da se prijavi kao RCON administrator.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnRconLoginAttempt

{{else if eq (lower $n) (lower "OnPlayerWeaponShot")}}
**- OnPlayerWeaponShot(playerid, weaponid, hittype, hitid, Float:fX, Float:fY, Float:fZ) -**
> **playerid** - id igrača koji je pucao.
> **weaponid** - id oružja koje je igrač koristio da puca.
> **hittype** - Vrsta objekta koji je igrač pogodio:
```
BULLET_HIT_TYPE_NONE            0 //Nikoga
BULLET_HIT_TYPE_PLAYER          1 //igrača
BULLET_HIT_TYPE_VEHICLE         2 //Vozilo
BULLET_HIT_TYPE_OBJECT          3 //Object
BULLET_HIT_TYPE_PLAYER_OBJECT   4 //PObject
```
> **hitid** - ID objekta kojeg je igrač upucao.
> **fX** - X kooridnate koje je hitac pogodio.
> **fY** - Y kooridnate koje je hitac pogodio.
> **fZ** - Z kooridnate koje je hitac pogodio.
Poziva se kada igrač puca iz oružja.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerWeaponShot

{{else if eq (lower $n) (lower "OnPlayerGiveDamage")}}
**- OnPlayerGiveDamage(playerid, damagedid, Float:amount, weaponid, bodypart) -**
> **playerid** - id igrača koji je naneo štetu.
> **damagedid** - ID igrača koji je primio štetu.
> **amount** - Količina štete koju je igrač naneo.
> **weaponid** - id oružja koje je igrač koristio da nanese štetu.
> **bodypart** - deo tela koji je igrač pogodio:
```
3	Torso
4	Groin
5	Left arm
6	Right arm
7	Left leg
8	Right leg
9	Head
```
Poziva se kada igrač zada štetu igraču.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerGiveDamage

{{else if eq (lower $n) (lower "OnPlayerTakeDamage")}}
**- OnPlayerTakeDamage(playerid, issuerid, Float:amount, weaponid, bodypart) -**
> **playerid** - id igrača koji je primio štetu.
> **issuerid** - ID igrača koji je naneo štetu.
> **amount** - Količina štete koju je igrač primio.
> **weaponid** - id oružja koje je issuerid koristio da nanese štetu.
> **bodypart** - deo tela koji je issuerid pogodio:
```
3	Torso
4	Groin
5	Left arm
6	Right arm
7	Left leg
8	Right leg
9	Head
```
Poziva se kada igrač primi štetu od issuerid-a.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerTakeDamage

{{else if eq (lower $n) (lower "OnPlayerEnterVehicle")}}
**- OnPlayerEnterVehicle(playerid, vehicleid, ispassenger) -**
> **playerid** - id igrača koji je ušao u vozilo.
> **vehicleid** - id vozila u koje je igrač ušao.
> **ispassenger** - da li je igrač suvozač (**0** ako je vozač : **1** ako je suvozač).
Poziva se kada igrač uđe u vozilo (kada pritisne ENTER, ne kada sedne!).

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerEnterVehicle

{{else if eq (lower $n) (lower "OnPlayerExitVehicle")}}
**- OnPlayerExitVehicle(playerid, vehicleid) -**
> **playerid** - id igrača koji izlazi iz vozila.
> **vehicleid** - id vozila iz kog igrač izlazi.
Poziva se kada igrač započne proces izlaženja iz vozila.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerExitVehicle

{{else if eq (lower $n) (lower "OnPlayerStateChange")}}
**- OnPlayerStateChange(playerid, newstate, oldstate) -**
> **playerid** - id igrača koji menja stanje.
> **newstate** - novi id stanja u koje igrač prelazi.
> **oldstate** - staro stanje u kom je igrač bio:
```
0	PLAYER_STATE_NONE		Nema
1	PLAYER_STATE_ONFOOT		Igrač je na nogama
2	PLAYER_STATE_DRIVER		Igrač je vozač
3	PLAYER_STATE_PASSENGER	Igrač je suvozač
7	PLAYER_STATE_WASTED		Igrač je mrtav ili bira klasu
8	PLAYER_STATE_SPAWNED	Igrač je spawnovan
9	PLAYER_STATE_SPECTATING	Igrač spectatuje
```
Poziva se kada igrač promeni 'stanje' u kom se nalazi.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerStateChange

{{else if eq (lower $n) (lower "OnPlayerKeyStateChange")}}
**- OnPlayerKeyStateChange(playerid, newkeys, oldkeys) -**
> **playerid** - id igrača koji stiska/pušta gumb na tastaturi.
> **newstate** - novi id key-a koji igrač stiska/pušta.
> **oldstate** - stari id key-a koji igrač stiska/pušta.
Poziva se kada igrač promeni 'stanje' gumbova na tastaturi.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerKeyStateChange

{{else if eq (lower $n) (lower "OnPlayerClickMap")}}
**- OnPlayerClickMap(playerid, Float:fX, Float:fY, Float:fZ) -**
> **playerid** - id igrača koji je označio kooridnate na mapi.
> **fX** - X kooridnate koje je igrač označio.
> **fY** - Y kooridnate koje je igrač označio.
> **fZ** - Z kooridnate koje je igrač označio.
Poziva se kada igrač označi kooridnate markerom na mapi.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerClickMap

{{else if eq (lower $n) (lower "OnPlayerClickPlayer")}}
**- OnPlayerClickPlayer(playerid, clickedplayerid, source) -**
> **playerid** - id igrača koji je kliknuo igrača.
> **clickedplayerid** - id igrača kojeg je playerid kliknuo.
> **source** - izvor igračevog klika (TAB).
Poziva se kada igrač označi kooridnate markerom na mapi.

**SA-MP Wiki:** https://wiki.sa-mp.com/wiki/OnPlayerClickPlayer
{{end}}
