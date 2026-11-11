/***************************************************************************************************
 * Grocery Inventory Demo
 * Copyright ⓒ 2026 Athaariq Ardhiansyah
 *
 * This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under
 * certain conditions. See LICENSE file.
 */
import './App.sass';

import { useState, useEffect } from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Autocomplete from '@mui/material/Autocomplete';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Table from '@mui/material/Table';
import TableHead from '@mui/material/TableHead';
import TableBody from '@mui/material/TableBody';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';

interface WarehouseData {
	id: number;
	name: string;
}

interface InventoryData {
	id: number;
	supplier: string;
	warehouse: string;
	warehouse_code: string;
	received_at: string;
	last_order_at: string;
	expired_at: string;
	name: string;
	status: string;
	unit_price_usd: number;
	percentage: number;
	stock_qty: number;
	reorder_level: number;
	reorder_qty: number;
	sales_volume: number;
	inventory_turnover_ratio: number;
}

export default function App() {
	const [warehouseData, setWarehouseData] = useState<Partial<WarehouseData>[] | null>(null);
	const [warehouseID, selectWarehouse] = useState<number>(0);
	const [bestSales, setBestSales] = useState<Partial<InventoryData>[]>([]);
	const [worstSales, setWorstSales] = useState<Partial<InventoryData>[]>([]);

	useEffect(() => {
		let isCanceled = false;

		fetch('http://localhost:8000/warehouses')
			.then((response) => {
				if (isCanceled) return;
				return response.json();
			})
			.then((warehouses) => {
				if (isCanceled || !warehouses) return;
				setWarehouseData([{ id: 0, name: 'All Warehouses' }, ...warehouses]);
			});

		return () => {
			isCanceled = true;
		};
	}, []);

	useEffect(() => {
		let isCanceled = false;

		const bestSalesURL = warehouseID
			? `http://localhost:8000/best_sales/${warehouseID}`
			: 'http://localhost:8000/best_sales';

		const worstSalesURL = warehouseID
			? `http://localhost:8000/worst_sales/${warehouseID}`
			: 'http://localhost:8000/worst_sales';

		fetch(bestSalesURL)
			.then((response) => {
				if (isCanceled) return;
				return response.json();
			})
			.then((data) => {
				if (isCanceled || !data) return;
				setBestSales(data);
			});

		fetch(worstSalesURL)
			.then((response) => {
				if (isCanceled) return;
				return response.json();
			})
			.then((data) => {
				if (isCanceled || !data) return;
				setWorstSales(data);
			});

		return () => {
			isCanceled = true;
		};
	}, [warehouseID, setBestSales, setWorstSales]);

	return (
		<Container>
			<Typography
				align="center"
				component="h1"
				variant="h4"
				sx={{ padding: '32px 16px 0px 16px' }}
			>
				Localhost Supermarket Inventory Demo
			</Typography>
			<Typography align="center" style={{ padding: '0px 16px 16px 16px' }}>
				by Athaariq Ardhiansyah R.
			</Typography>
			<Autocomplete
				disablePortal
				disableClearable
				value={warehouseID}
				defaultValue={0}
				options={warehouseData || [{ id: 0, name: 'All Warehouses' }]}
				getOptionKey={(item) => item.id!}
				getOptionValue={(item) => item.id!}
				getOptionLabel={(item) => item.name || '{Unlabeled}'}
				renderInput={(params) => <TextField {...params} label="Filter by warehouse" />}
				sx={{ width: '320px', maxWidth: '100%', margin: '16px auto' }}
				onChange={(_, value) => {
					if (typeof value === 'number') selectWarehouse(value);
				}}
			/>
			<Box
				sx={{
					margin: '16px 0px',
					display: 'flex',
					padding: '16px',
					gap: '16px',
					flexDirection: 'column',
					justifyContent: 'center',
					'@media (min-width: 960px)': { flexDirection: 'row' }
				}}
			>
				<Card sx={{ borderRadius: '16px', padding: '16px' }}>
					<Typography align="center" component="h2" variant="h5">
						Most Purchased Items
					</Typography>
					<Table>
						<TableHead>
							<TableRow>
								<TableCell>Name</TableCell>
								<TableCell>Supplier</TableCell>
								<TableCell>Reorder Ratio (%)</TableCell>
							</TableRow>
						</TableHead>
						<TableBody>
							{bestSales
								.filter(
									(each) =>
										each.id &&
										each.name &&
										each.supplier &&
										typeof each.reorder_level === 'number'
								)
								.map((each) => (
									<TableRow key={each.id!}>
										<TableCell>{each.name!}</TableCell>
										<TableCell>{each.supplier!}</TableCell>
										<TableCell>{each.reorder_level!}</TableCell>
									</TableRow>
								))}
						</TableBody>
					</Table>
				</Card>
				<Card sx={{ borderRadius: '16px', padding: '16px' }}>
					<Typography align="center" component="h2" variant="h5">
						Least Purchased Items
					</Typography>
					<Table>
						<TableHead>
							<TableRow>
								<TableCell>Name</TableCell>
								<TableCell>Supplier</TableCell>
								<TableCell>Reorder Ratio (%)</TableCell>
							</TableRow>
						</TableHead>
						<TableBody>
							{worstSales
								.filter(
									(each) =>
										each.id &&
										each.name &&
										each.supplier &&
										typeof each.reorder_level === 'number'
								)
								.map((each) => (
									<TableRow key={each.id!}>
										<TableCell>{each.name!}</TableCell>
										<TableCell>{each.supplier!}</TableCell>
										<TableCell>{each.reorder_level!}</TableCell>
									</TableRow>
								))}
						</TableBody>
					</Table>
				</Card>
			</Box>
		</Container>
	);
}
