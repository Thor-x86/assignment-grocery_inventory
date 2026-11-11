/***************************************************************************************************
 * Grocery Inventory Demo
 * Copyright ⓒ 2026 Athaariq Ardhiansyah
 *
 * This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under
 * certain conditions. See LICENSE file.
 */
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';

createRoot(document.getElementById('root')!).render(
	<StrictMode>
		<App />
	</StrictMode>
);
