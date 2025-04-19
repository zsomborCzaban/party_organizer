- loginnal ha rossz a jelszo akkor valami felre megy a backenden. (internal server errort kapunk amikor nem kene)
- hashelve tarolni a jelszot a registration requestben, és kivenni a confirm jelszot

- error kezelesek kipofozasa (ha loginnal nincs backend akkor ne azt adjuk vissza, hogy invalid username or password)
- scroll to the top on tab valtas frontenden

auto delete on parties after they have been hosted
delete parties (maybe)


# doksi
oldal terkep
modellek kozotti kapcsolat

tanulsag: model-t masik folderbe tenni mint, hogy lehessen back referenceket csinalni

concurrent mapwrite errort fixelni backendben

when inviting nonexistong user give back normal error message





tovabbi fejlesztések:
- refresh tokenek oauth bejelntkezéshez.
- 3rd party user providerrel integrálódni
- misc contributions ahol nicns requirement
- endpointok atirasa querry parameterekre
- ci/cl setup githubon
- hostolásnál monitorozas kiepitese
- egyed kapcsolat diagram javitasa a dokumentacioban (kiegeszitese a regisztracios requesttel)
- link beviteli mezok validalasa phising linkek ellen
