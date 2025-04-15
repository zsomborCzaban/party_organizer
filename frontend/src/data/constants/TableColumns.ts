import {Column} from "../../components/table/SortableTable.tsx";

export const partyInviteTableColumns: Column<PartyInviteTableRow>[] = [
    {
        headerName: 'Invited by',
        field: 'invitedBy',
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
    },
];

export interface PartyInviteTableRow {
    id: number;
    invitedBy: string;
    partyName: string;
    partyPlace: string;
    partyTime: Date;
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
    },
    {
        headerName: 'Organizer',
        field: 'organizerName',
    },
];

export interface PartyTableRow {
    id: number;
    name: string;
    place: string;
    time: string;
    organizerName: string;
}

export const partyTableColumnsLegacy = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
    },
    {
        title: 'Place',
        dataIndex: 'place',
        key: 'place',
    },
    {
        title: 'Time',
        dataIndex: 'start_time',
        key: 'time',
    },
    {
        title: 'Organizer',
        // dataIndex: ['organizer', 'username'],
        dataIndex: ['organizer', 'username'],
        key: 'organizer',
    },
    {
        // todo: to be done in backend
        title: 'Headcount',
        dataIndex: 'headcount',
        key: 'headcount',
    },
];

export const partyInviteTableColumnsLegacy = [
    {
        title: 'Invited by',
        dataIndex: ['invitor', 'username'],
        key: 'invited by',
    },
    {
        title: 'To party',
        dataIndex: ['party', 'name'],
        key: 'to party',
    },
    {
        title: 'Place',
        dataIndex: ['party', 'place'],
        key: 'place',
    },
    {
        title: 'Time',
        dataIndex: ['party', 'start_time'],
        key: 'time',
    },

    {
        // todo: to be done in backend
        title: 'Headcount',
        dataIndex: ['party', 'headcount'],
        key: 'headcount',
    },
];

export const invitedByTableColumnsLegacy =[
    {
        title: 'Invited by',
        dataIndex: ['invitor', 'username'],
        key: 'username',
    },
];
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