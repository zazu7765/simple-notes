<script lang="ts">
	import { get } from 'svelte/store';
	import icon from '$lib/assets/icon.png'
	import { onMount } from 'svelte';

	// import { onMount } from "svelte";

	// onMount(() => {

	//    return document.querySelector('#menu')
	let menu = 'hidden';
	export let login: boolean;


	

	// });
	console.log(login);
</script>

<div>
	<body
		class="
        antialiased
 
      "
	>
		<header>
			<nav
				class="

          fixed-top
            flex flex-wrap
            items-center
            justify-between
            w-full
            py-4
            md:py-0
            px-4
            text-lg text-gray-700
            bg-white

          "
			>
				<div class=" h-15 w-12 object-contain">
					<a href="#">
					<img src='./icon.jpg' class="object-contain">
					</a>
				</div>

				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<svg
					on:click={() => {
						if (menu === 'hidden') {
							menu = 'visible';
						} else {
							menu = 'hidden';
						}
					}}
					xmlns="http://www.w3.org/2000/svg"
					id="menu-button"
					class="h-6 w-6 cursor-pointer md:hidden block"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h16M4 18h16"
					/>
				</svg>

				<div class="{menu} w-full md:flex md:items-center md:w-auto" id="menu">
					<ul
						class="
                pt-4
                text-base text-gray-700
                md:flex
                md:justify-between 
                md:pt-0"
					>
						{#if !login}
							<li>
								<a class="md:p-4 py-2 block hover:text-purple-400" href="/login">Login</a>
							</li>
							<li>
								<a class="md:p-4 py-2 block hover:text-purple-400 text-purple-500" href="/sign"
									>Sign Up</a
								>
							</li>
						{:else}
							<li>
								<a class="md:p-4 py-2 block hover:text-purple-400 text-purple-500" href="/test"
									>Open Notes</a
								>
              </li>
              <li>
								<a class="md:p-4 py-2 block hover:text-purple-400 text-purple-500" href="/logout"
									>Logout</a
								>
							</li>
						{/if}
					</ul>
				</div>
			</nav>
		</header>
	</body>
</div>

<style>
	:global(body) {
		background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.com/svgjs' width='1980' height='1080' preserveAspectRatio='none' viewBox='0 0 1980 1080'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1008%26quot%3b)' fill='none'%3e%3crect width='1980' height='1080' x='0' y='0' fill='%230e2a47'%3e%3c/rect%3e%3cpath d='M1758.14 1007.3a47.31 47.31 0 1 0-23.79-91.58z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M741.51 698.53L820.29 698.53L820.29 777.31L741.51 777.31z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1635.94 355.16L1651.53 355.16L1651.53 456.73L1635.94 456.73z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M137.6 471.44a55.5 55.5 0 1 0 101.4-45.14z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M111.05 104.25a107.76 107.76 0 1 0 117.43-180.72z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1282.09 400.16a57.41 57.41 0 1 0 112.9-20.88z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M128.3 814.79 a89.78 89.78 0 1 0 179.56 0 a89.78 89.78 0 1 0 -179.56 0z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M348.38 905.1L420.07 905.1L420.07 944.43L348.38 944.43z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M300.06 31.26L305.12 31.26L305.12 36.32L300.06 36.32z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1642.5 796.03a29.21 29.21 0 1 0 19.03 55.23z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1100.98 552.17L1198.92 552.17L1198.92 552.84L1100.98 552.84z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1534.08 517.77L1614.1 517.77L1614.1 597.79L1534.08 597.79z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M396.05 760.81 a26.35 26.35 0 1 0 52.7 0 a26.35 26.35 0 1 0 -52.7 0z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M582.41 258.74L652.36 258.74L652.36 354.65L582.41 354.65z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M27.1 96.53a41.94 41.94 0 1 0-31.96-77.55z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1585.84 476.77a78.03 78.03 0 1 0 121.35 98.13z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1139.43 508.17L1139.94 508.17L1139.94 578.58L1139.43 578.58z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M659.63 59.8a67.89 67.89 0 1 0 65.92 118.7z' stroke='%23037b0b'%3e%3c/path%3e%3cpath d='M449.84 98.1 a43.74 43.74 0 1 0 87.48 0 a43.74 43.74 0 1 0 -87.48 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M746.29 258.6 a81.09 81.09 0 1 0 162.18 0 a81.09 81.09 0 1 0 -162.18 0z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1370.87 526.01 a81.86 81.86 0 1 0 163.72 0 a81.86 81.86 0 1 0 -163.72 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M147.8 755.43 a78.97 78.97 0 1 0 157.94 0 a78.97 78.97 0 1 0 -157.94 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1794.71 84.39L1898.8 84.39L1898.8 96.33L1794.71 96.33z' fill='%23e73635'%3e%3c/path%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1008'%3e%3crect width='1980' height='1080' fill='white'%3e%3c/rect%3e%3c/mask%3e%3c/defs%3e%3c/svg%3e");
	}
</style>
