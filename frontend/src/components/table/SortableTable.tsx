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
} from '@mui/material';

interface Column<T> {
    field: keyof T;
    headerName: string;
    width?: number;
    render?: (row: T) => React.ReactNode;
}

interface SortableTableProps<T> {
    columns: Column<T>[];
    data: T[];
    defaultSortField?: keyof T;
    rowsPerPageOptions?: number[];
    defaultRowsPerPage?: number;
}

export function SortableTable<T extends { id: string | number }>({
                                                                     columns,
                                                                     data,
                                                                     defaultSortField = columns[0].field,
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
                        />
                    </TableRow>
                </TableFooter>
            </Table>
        </TableContainer>
    );
}