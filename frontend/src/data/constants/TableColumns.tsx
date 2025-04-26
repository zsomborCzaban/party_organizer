import { Column } from "../../components/table/SortableTable";

export interface PartyTableRow {
    id: number;
    name: string;
    place: string;
    time: string;
    organizerName: string;
    organizerProfilePicture: string;
}

export const partyTableColumns: Column<PartyTableRow>[] = [
    {
        headerName: 'Name',
        field: 'name',
    },
    {
        headerName: 'Place',
        field: 'place',
    },
    {
        headerName: 'Time',
        field: 'time',
        render: (row) => {
            const date = new Date(row.time);
            return date.toLocaleString();
        }
    },
    {
        headerName: 'Organizer',
        field: 'organizerName',
        render: (row) => (
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                {row.organizerProfilePicture && (
                    <img 
                        src={row.organizerProfilePicture} 
                        alt={row.organizerName}
                        style={{ width: '24px', height: '24px', borderRadius: '50%' }}
                    />
                )}
                <span>{row.organizerName}</span>
            </div>
        )
    },
];

export interface PartyInviteTableRow {
    id: number;
    invitedBy: string;
    partyName: string;
    partyPlace: string;
    partyTime: string;
    invitorProfilePicture: string;

}

export const partyInviteTableColumns: Column<PartyInviteTableRow>[] = [
    {
        headerName: 'Invited by',
        field: 'invitedBy',
        render: (row) => (
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                <img
                    src={row.invitorProfilePicture}
                    alt={row.invitedBy}
                    style={{ width: '24px', height: '24px', borderRadius: '50%' }}
                />
                <span>{row.invitedBy}</span>
            </div>
        ),
    },
    {
        headerName: 'To party',
        field: 'partyName',
    },
    {
        headerName: 'Place',
        field: 'partyPlace',
    },
    {
        headerName: 'Time',
        field: 'partyTime',
        render: (row) => {
            const date = new Date(row.partyTime);
            return date.toLocaleString();
        }
    },
];

export interface FriendTableRow {
    id: number;
    username: string;
    friendProfilePicture: string;
    email: string;
}

export const FriendTableColumns: Column<FriendTableRow>[] = [
    {
        headerName: 'Username',
        field: 'username',
        render: (row) => (
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                <img
                    src={row.friendProfilePicture}
                    alt={row.username}
                    style={{ width: '24px', height: '24px', borderRadius: '50%' }}
                />
                <span>{row.username}</span>
            </div>
        ),
    },
    {
        headerName: 'Email',
        field: 'email',
    },
]

export interface FriendInviteTableRow {
    id: number;
    invitedBy: string;
    invitorProfilePicture: string;
}

export const FriendInviteTableColumns: Column<FriendInviteTableRow>[] = [
    {
        headerName: 'Invited by',
        field: 'invitedBy',
        render: (row) => (
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                <img
                    src={row.invitorProfilePicture}
                    alt={row.invitedBy}
                    style={{ width: '24px', height: '24px', borderRadius: '50%' }}
                />
                <span>{row.invitedBy}</span>
            </div>
        ),
    },
]

export const invitedTableColumnsLegacy =[
    {
        title: 'Invited',
        dataIndex: ['invited', 'username'],
        key: 'username',
    },
    {
        title: 'State',
        dataIndex: 'state',
        key: 'state',
    },
];

export const userTableColumnsLegacy = [
    {
        title: 'Username',
        dataIndex: 'username',
        key: 'username',
    },
    {
        title: 'Email',
        dataIndex: 'email',
        key: 'email',
    },
];

export const requirementTableColumnsLegacy = [
    {
        title: 'Type',
        dataIndex: 'type',
        key: 'type',
    },
    {
        title: 'Target Quantity',
        dataIndex: 'target_quantity',
        key: 'target_quantity',
    },
    {
        title: 'Quantity Mark',
        dataIndex: 'quantity_mark',
        key: 'quantity_mark',
    },
];