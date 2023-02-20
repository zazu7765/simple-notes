<script module lang="ts">
	import { redirect } from '@sveltejs/kit';
	import Sidebar from '../../../lib/Sidebar.svelte';
	import Notebook from '../../../lib/notes.svelte';
	import { write } from '../../../lib/notes';
	export let data: any;
	let login = data.bool;
	$write = data['responseNote'];
	let pp;
	write.subscribe((value) => {
		pp = value;
	});
	console.log(pp['Data'] + ' note');
	// console.log(data['responseNote'])

	let num = 5;
	console.log(data);
	let content = data['response']['Data'];
	let arr: Array<string>[] = [];
	for (const item in content) {
		arr.push([content[item]['ID'], content[item]['Title']]);
	}
	console.log(arr);
	let go = () => {
		throw redirect(302, '/login');
	};
</script>

<button
	data-drawer-target="logo-sidebar"
	data-drawer-toggle="logo-sidebar"
	aria-controls="logo-sidebar"
	type="button"
	class="inline-flex items-center p-2 mt-2 ml-3 text-sm text-white rounded-lg md:hidden hover:bg-gray-500 bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
>
	<span class="sr-only">Open sidebar</span>
	<svg
		class="w-6 h-6"
		aria-hidden="true"
		fill="currentColor"
		viewBox="0 0 20 20"
		xmlns="http://www.w3.org/2000/svg"
	>
		<path
			clip-rule="evenodd"
			fill-rule="evenodd"
			d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"
		/>
	</svg>
</button>

<aside id="sidebar-multi-level-sidebar" class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0" aria-label="Sidebar">
	<div class="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
		<a href="https://flowbite.com/" class="flex items-center pl-2.5 mb-5">
			<img
				src="https://flowbite.com/docs/images/logo.svg"
				class="h-6 mr-3 sm:h-7"
				alt="Flowbite Logo"
			/>
			<span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
				>Flowbite</span
			>
		</a>
		
		<ul class="space-y-2">
			<li>
				<a
					href="#"
					class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
				>
					<svg
						aria-hidden="true"
						class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
							clip-rule="evenodd"
						/></svg
					>
					<span class="flex-1 ml-3 whitespace-nowrap">Users</span>
				</a>
			</li>
			<ul class="pt-4 mt-4 space-y-2 border-t border-gray-500 dark:border-gray-700" />

			{#each arr as name, index}
			<li>
				<Sidebar {name} />
			</li>
			{/each}

			<ul class="pt-4 mt-4 space-y-2 border-t border-gray-500 dark:border-gray-700" />
			<li>
				<!-- weird shit target="_self" -->
				<a
					href="/logout"
					target="_self"
					class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
				>
					<svg
						aria-hidden="true"
						class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z"
							clip-rule="evenodd"
						/></svg
					>
					<span class="flex-1 ml-3 whitespace-nowrap">Logout</span>
				</a>
			</li>
		</ul>
	</div>
</aside>
<div class="grid lg:grid-cols-3 md:grid-cols-2 xl:grid-cols-4 gap-4 grid-flow-row p-5" />

<style>
	:global(body) {
		background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.com/svgjs' width='1980' height='1080' preserveAspectRatio='none' viewBox='0 0 1980 1080'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1008%26quot%3b)' fill='none'%3e%3crect width='1980' height='1080' x='0' y='0' fill='%230e2a47'%3e%3c/rect%3e%3cpath d='M1758.14 1007.3a47.31 47.31 0 1 0-23.79-91.58z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M741.51 698.53L820.29 698.53L820.29 777.31L741.51 777.31z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1635.94 355.16L1651.53 355.16L1651.53 456.73L1635.94 456.73z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M137.6 471.44a55.5 55.5 0 1 0 101.4-45.14z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M111.05 104.25a107.76 107.76 0 1 0 117.43-180.72z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1282.09 400.16a57.41 57.41 0 1 0 112.9-20.88z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M128.3 814.79 a89.78 89.78 0 1 0 179.56 0 a89.78 89.78 0 1 0 -179.56 0z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M348.38 905.1L420.07 905.1L420.07 944.43L348.38 944.43z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M300.06 31.26L305.12 31.26L305.12 36.32L300.06 36.32z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1642.5 796.03a29.21 29.21 0 1 0 19.03 55.23z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1100.98 552.17L1198.92 552.17L1198.92 552.84L1100.98 552.84z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1534.08 517.77L1614.1 517.77L1614.1 597.79L1534.08 597.79z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M396.05 760.81 a26.35 26.35 0 1 0 52.7 0 a26.35 26.35 0 1 0 -52.7 0z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M582.41 258.74L652.36 258.74L652.36 354.65L582.41 354.65z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M27.1 96.53a41.94 41.94 0 1 0-31.96-77.55z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1585.84 476.77a78.03 78.03 0 1 0 121.35 98.13z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1139.43 508.17L1139.94 508.17L1139.94 578.58L1139.43 578.58z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M659.63 59.8a67.89 67.89 0 1 0 65.92 118.7z' stroke='%23037b0b'%3e%3c/path%3e%3cpath d='M449.84 98.1 a43.74 43.74 0 1 0 87.48 0 a43.74 43.74 0 1 0 -87.48 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M746.29 258.6 a81.09 81.09 0 1 0 162.18 0 a81.09 81.09 0 1 0 -162.18 0z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1370.87 526.01 a81.86 81.86 0 1 0 163.72 0 a81.86 81.86 0 1 0 -163.72 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M147.8 755.43 a78.97 78.97 0 1 0 157.94 0 a78.97 78.97 0 1 0 -157.94 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1794.71 84.39L1898.8 84.39L1898.8 96.33L1794.71 96.33z' fill='%23e73635'%3e%3c/path%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1008'%3e%3crect width='1980' height='1080' fill='white'%3e%3c/rect%3e%3c/mask%3e%3c/defs%3e%3c/svg%3e");
	}
</style>
