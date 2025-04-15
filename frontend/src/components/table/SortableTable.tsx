import React, { useState } from 'react';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    TableSortLabel,
    Paper,
    TableFooter,
    TablePagination,
    IconButton,
    Button,
    Box
} from '@mui/material';

export interface Column<T> {
    field: keyof T;
    headerName: string;
    width?: number;
    render?: (row: T) => React.ReactNode;
}

export interface ActionButton<T> {
    label?: string;
    icon?: React.ReactNode;
    color?: 'inherit' | 'primary' | 'secondary' | 'success' | 'error' | 'info' | 'warning';
    onClick: (row: T) => void;
}

interface SortableTableProps<T> {
    columns: Column<T>[];
    data: T[];
    actionButtons?: ActionButton<T>[];
    defaultSortField?: keyof T;
    rowsPerPageOptions?: number[];
    defaultRowsPerPage?: number;
}

export function SortableTable<T extends { id: string | number }>({
                                                                     columns,
                                                                     data,
                                                                     actionButtons = [],
                                                                     defaultSortField = columns[0]?.field,
                                                                     rowsPerPageOptions = [5, 10, 25],
                                                                     defaultRowsPerPage = 5,
                                                                 }: SortableTableProps<T>) {
    const [order, setOrder] = useState<'asc' | 'desc'>('asc');
    const [orderBy, setOrderBy] = useState<keyof T>(defaultSortField);
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(defaultRowsPerPage);

    const handleSort = (property: keyof T) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };

    const handleChangePage = (event: unknown, newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    // Sort and paginate the data
    const sortedData = [...data].sort((a, b) => {
        if (a[orderBy] < b[orderBy]) return order === 'asc' ? -1 : 1;
        if (a[orderBy] > b[orderBy]) return order === 'asc' ? 1 : -1;
        return 0;
    });

    const paginatedData = sortedData.slice(
        page * rowsPerPage,
        page * rowsPerPage + rowsPerPage
    );

    return (
        <TableContainer component={Paper}>
            <Table>
                <TableHead>
                    <TableRow>
                        {columns.map((column) => (
                            <TableCell key={column.field.toString()}>
                                <TableSortLabel
                                    active={orderBy === column.field}
                                    direction={orderBy === column.field ? order : 'asc'}
                                    onClick={() => handleSort(column.field)}
                                >
                                    {column.headerName}
                                </TableSortLabel>
                            </TableCell>
                        ))}
                        {actionButtons.length > 0 && (
                            <TableCell align="right" width={actionButtons.length * 50}>
                                Actions
                            </TableCell>
                        )}
                    </TableRow>
                </TableHead>
                <TableBody>
                    {paginatedData.map((row) => (
                        <TableRow key={row.id}>
                            {columns.map((column) => (
                                <TableCell key={`${row.id}-${column.field.toString()}`}>
                                    {column.render ? column.render(row) : String(row[column.field])}
                                </TableCell>
                            ))}
                            {actionButtons.length > 0 && (
                                <TableCell align="right">
                                    <Box display="flex" gap={1} justifyContent="flex-end">
                                        {actionButtons.map((button, index) => (
                                            button.icon ? (
                                                <IconButton
                                                    key={index}
                                                    color={button.color || 'primary'}
                                                    onClick={() => button.onClick(row)}
                                                >
                                                    {button.icon}
                                                </IconButton>
                                            ) : (
                                                <Button
                                                    key={index}
                                                    variant="outlined"
                                                    color={button.color || 'primary'}
                                                    onClick={() => button.onClick(row)}
                                                    size="small"
                                                >
                                                    {button.label}
                                                </Button>
                                            )
                                        ))}
                                    </Box>
                                </TableCell>
                            )}
                        </TableRow>
                    ))}
                </TableBody>
                <TableFooter>
                    <TableRow>
                        <TablePagination
                            rowsPerPageOptions={rowsPerPageOptions}
                            count={data.length}
                            rowsPerPage={rowsPerPage}
                            page={page}
                            onPageChange={handleChangePage}
                            onRowsPerPageChange={handleChangeRowsPerPage}
                            colSpan={columns.length + (actionButtons.length > 0 ? 1 : 0)}
                        />
                    </TableRow>
                </TableFooter>
            </Table>
        </TableContainer>
    );
}