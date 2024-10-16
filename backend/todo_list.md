# Todos
 - on deleting an entity create a callback for the otther entities (example: when deleting a party, delete every contribution to it too) - (maybe not needed, bc when a party is deletet were not gonna access any of its entities. altough it might be needed for other entites. also a trash collector would be good if were only going to delete parties)
 - validalni az a request adatokat (pl user email)
 - user registraciot kiboviteni megerosito emaillel es elfelejtett jelszoval


dto-k removolasa a useren kivul

admin user should have id 1, otherwise it will create buggs

auth party service (description in comments in api_party_service)

At update check of original record's organizer/creator ids have been changed

refactor party_invites tabel to be the join table for party_participants

implement omitupdates?

contributions doesnt need party objectum