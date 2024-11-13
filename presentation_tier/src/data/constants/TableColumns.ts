export const partyTableColumns = [
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

export const partyInviteTableColumns = [
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

export const invitedByTableColumns =[
    {
        title: 'Invited by',
        dataIndex: ['invitor', 'username'],
        key: 'username',
    },
];
export const invitedTableColumns =[
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

export const userTableColumns = [
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

export const requirementTableColumns = [
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