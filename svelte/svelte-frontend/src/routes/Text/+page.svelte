<script>
	import { onDestroy } from 'svelte';

	import { quill } from 'svelte-quill';

	let options = { placeholder: 'Write something from outside...' };

	let content = { html: '', text: '' };
	onDestroy(() => {});
</script>

<svelte:head>
	<link href="//cdn.quilljs.com/1.3.6/quill.snow.css" rel="stylesheet" />
</svelte:head>
<body>
	<div class="grid lg:grid-cols-5 gap-4 min-w-max p-5 max-h-5 min-h-screen">
		<div />
		<div
			class="min-h-none col-span-3 grid-cols-3 gap-4 py-12 px-4 content-center inline-block sm:px-6 lg:px-8 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
		>
			<div class="col-span-3 p-4 rounded-lg">
				<!-- <h1>SVELTE-QUILL</h1> -->

				<div
					class="editor min-h-screen "
					use:quill={options}
					on:text-change={(e) => (content = e.detail)}
				/>

				<!-- <br /> -->
				<!-- Resulting HTML:
		
	{@html content.html} -->
			</div>
		</div>
	</div>
</body>

<style>
	:global(body) {
		background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.com/svgjs' width='1980' height='1080' preserveAspectRatio='none' viewBox='0 0 1980 1080'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1008%26quot%3b)' fill='none'%3e%3crect width='1980' height='1080' x='0' y='0' fill='%230e2a47'%3e%3c/rect%3e%3cpath d='M1758.14 1007.3a47.31 47.31 0 1 0-23.79-91.58z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M741.51 698.53L820.29 698.53L820.29 777.31L741.51 777.31z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1635.94 355.16L1651.53 355.16L1651.53 456.73L1635.94 456.73z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M137.6 471.44a55.5 55.5 0 1 0 101.4-45.14z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M111.05 104.25a107.76 107.76 0 1 0 117.43-180.72z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1282.09 400.16a57.41 57.41 0 1 0 112.9-20.88z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M128.3 814.79 a89.78 89.78 0 1 0 179.56 0 a89.78 89.78 0 1 0 -179.56 0z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M348.38 905.1L420.07 905.1L420.07 944.43L348.38 944.43z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M300.06 31.26L305.12 31.26L305.12 36.32L300.06 36.32z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1642.5 796.03a29.21 29.21 0 1 0 19.03 55.23z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1100.98 552.17L1198.92 552.17L1198.92 552.84L1100.98 552.84z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1534.08 517.77L1614.1 517.77L1614.1 597.79L1534.08 597.79z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M396.05 760.81 a26.35 26.35 0 1 0 52.7 0 a26.35 26.35 0 1 0 -52.7 0z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M582.41 258.74L652.36 258.74L652.36 354.65L582.41 354.65z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M27.1 96.53a41.94 41.94 0 1 0-31.96-77.55z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1585.84 476.77a78.03 78.03 0 1 0 121.35 98.13z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1139.43 508.17L1139.94 508.17L1139.94 578.58L1139.43 578.58z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M659.63 59.8a67.89 67.89 0 1 0 65.92 118.7z' stroke='%23037b0b'%3e%3c/path%3e%3cpath d='M449.84 98.1 a43.74 43.74 0 1 0 87.48 0 a43.74 43.74 0 1 0 -87.48 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M746.29 258.6 a81.09 81.09 0 1 0 162.18 0 a81.09 81.09 0 1 0 -162.18 0z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1370.87 526.01 a81.86 81.86 0 1 0 163.72 0 a81.86 81.86 0 1 0 -163.72 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M147.8 755.43 a78.97 78.97 0 1 0 157.94 0 a78.97 78.97 0 1 0 -157.94 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1794.71 84.39L1898.8 84.39L1898.8 96.33L1794.71 96.33z' fill='%23e73635'%3e%3c/path%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1008'%3e%3crect width='1980' height='1080' fill='white'%3e%3c/rect%3e%3c/mask%3e%3c/defs%3e%3c/svg%3e");
	}
</style>
