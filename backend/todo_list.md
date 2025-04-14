-endpointokat (party, friendManager, partyAttendanceManager), atirni hogy ne-automatikusan szedjuk ki az id-t hanem parameterben kelljen megadni
-loginnal ha rossz a jelszo akkor valami felre megy a backenden. (internal server errort kapunk amikor nem kene)
-create partynal az accesscode elejere az id 0 lesz uj partnal
- hashelve tarolni a jelszot a registration requestben, és kivenni a confirm jelszot


-publikus endpoint lehet a adiscover és a party home page



# Todos
 - validalni az a request adatokat (pl user email)
 - user registraciot kiboviteni megerosito emaillel es elfelejtett jelszoval
 - paginizáció


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





tovabbi fejlesztések:
- refresh tokenek oauth bejelntkezéshez.
- 3rd party user providerrel integrálódni
- misc contributions ahol nicns requirement
- endpointok atirasa querry parameterekre
- ci/cl setup githubon
- hostolásnál monitorozas kiepitese
- egyed kapcsolat diagram javitasa a dokumentacioban (kiegeszitese a regisztracios requesttel)