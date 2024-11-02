

# Todos
 - validalni az a request adatokat (pl user email)
 - user registraciot kiboviteni megerosito emaillel es elfelejtett jelszoval


dto-k removolasa a useren kivul

auto delete on parties after they have been hosted
delete parties (maybe)

rename party_invites to partyAttandanceManager


# doksi
oldal terkep
modellek kozotti kapcsolat

tanulsag: model-t masik folderbe tenni mint, hogy lehessen back referenceket csinalni

usernek visszjelzees timeout utan

sikeres register utan egy timer, hogy "you will be navigated to the login page in 5... 4, 3, 2, 1 bumm"

validate party on backend

a party sliceokat osszetenni 1 sliceá és fronenden szurni az eppen szuksegesekre
pros: most epp jobb megoldas lenne, cons: ha hasznalunk graphql-t es egy requesttel lejon minden akkopr a mostani hatekonyabb megoldas mert ugyanannyi request, viszont amikor volatoztatunk 1 adatot akkor nem kell az egeszet ujra requestelni hanem csak azt a reszt es frontenden sem kell majd kulon szurni(altough picivel nagyobb maintain)

profilkép a usereknek, és megjeleníteni a contribution és hall of fame oldalakon

dupla requestek javítása a overView frontenden

concurrent mapwrite errort fixelni backendben