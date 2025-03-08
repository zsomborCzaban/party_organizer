-endpointokat (party, friendManager, partyAttendanceManager), atirni hogy ne-automatikusan szedjuk ki az id-t hanem parameterben kelljen megadni
-test friendManager on bruno. on Removefriend we get: no such column: users.user_id
-loginnal ha rossz a jelszo akkor valami felre megy a backenden. (internal server errort kapunk amikor nem kene)

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

profilkép a usereknek, és megjeleníteni a contribution és hall of fame oldalakon

dupla requestek javítása a overView frontenden

navbar es profile kimozgatasa kivulre
todo: useSelector party-t is kivul hasznalni

concurrent mapwrite errort fixelni backendben

groupchat

when inviting nonexistong user give back normal error message