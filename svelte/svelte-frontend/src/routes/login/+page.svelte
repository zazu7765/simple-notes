<script lang="ts">
	import { enhance } from '$app/forms';
	import NavBar from '../../lib/NavBar.svelte';
	// export let data:any;
	// console.log(data);
	let Email = 'email@example.com';
	let Password = 'password';
	let total: any;
	async function signIn() {
		const response = await fetch('http://localhost:81/login', {
			method: 'POST',
			body: JSON.stringify({ Email: Email, Password: Password }),
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			}
		});
		total = await response.json();
		console.log(total);
	}

	/** @type {import('./$types').ActionData} */
	export let form;
	if(form){
	console.log(form.success+"fasdfdasf");
	}
	

	let login = false;
</script>

<div>
	<NavBar {login} />
</div>

<!--
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/forms'),
    ],
  }
  ```
-->
<!--
  This example requires updating your template:

  ```
  <html class="h-full bg-gray-50">
  <body class="h-full">
  ```
-->
<div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8 ">
	<div
		class="w-full max-w-md space-y-8 p-4 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
	>
		<div>
			<img
				class="mx-auto h-12 w-auto"
				src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
				alt="Your Company"
			/>
			<h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">
				Sign in to your Note's account
			</h2>
			<!-- <p class="mt-2 text-center text-sm text-gray-600">
          Or
          <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">start your 14-day free trial</a>
        </p> -->
		</div>
		<form class="mt-8 space-y-6" method="POST">
			<input type="hidden" name="remember" value="true" />
			<div class="-space-y-px rounded-md shadow-sm">
				<div>
					<label for="email-address" class="sr-only">Email address</label>
					<input
						id="email-address"
						name="email"
						type="email"
						autocomplete="email"
						required
						class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
						placeholder="Email address"
					/>
				</div>
				<div>
					<label for="password" class="sr-only">Password</label>
					<input
						id="password"
						name="password"
						type="password"
						autocomplete="current-password"
						required
						class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
						placeholder="Password"
					/>
				</div>
			</div>

			<!-- <div class="flex items-center justify-between">
				<div class="flex items-center">
					<input
						id="remember-me"
						name="remember-me"
						type="checkbox"
						class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
					/> -->
					<!-- <label for="remember-me" class="ml-2 block text-sm text-gray-900">Remember me</label> -->
				<!-- </div> -->
<!-- 
				<div class="text-sm">
					<a href="../notes" class="font-medium text-indigo-600 hover:text-indigo-500"
						>Forgot your password?</a
					>
				</div>
			</div> -->

			<div>
				<button
					type="submit"
					class="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
				>
					<span class="absolute inset-y-0 left-0 flex items-center pl-3">
						<!-- Heroicon name: mini/lock-closed -->
						<svg
							class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400"
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z"
								clip-rule="evenodd"
							/>
						</svg>
					</span>
					Login
				</button>
				{#if form}
				{#if form.success == false}
				<p class="mt-2 text-sm text-red-600 dark:text-red-500"><span class="font-medium">Oh, snapp!</span> {form.message}</p>
				{/if}
				{/if}
			</div>
		</form>
	</div>
</div>
<style>
	:global(body) {
		background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.com/svgjs' width='1980' height='1080' preserveAspectRatio='none' viewBox='0 0 1980 1080'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1008%26quot%3b)' fill='none'%3e%3crect width='1980' height='1080' x='0' y='0' fill='%230e2a47'%3e%3c/rect%3e%3cpath d='M1758.14 1007.3a47.31 47.31 0 1 0-23.79-91.58z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M741.51 698.53L820.29 698.53L820.29 777.31L741.51 777.31z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1635.94 355.16L1651.53 355.16L1651.53 456.73L1635.94 456.73z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M137.6 471.44a55.5 55.5 0 1 0 101.4-45.14z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M111.05 104.25a107.76 107.76 0 1 0 117.43-180.72z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1282.09 400.16a57.41 57.41 0 1 0 112.9-20.88z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M128.3 814.79 a89.78 89.78 0 1 0 179.56 0 a89.78 89.78 0 1 0 -179.56 0z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M348.38 905.1L420.07 905.1L420.07 944.43L348.38 944.43z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M300.06 31.26L305.12 31.26L305.12 36.32L300.06 36.32z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1642.5 796.03a29.21 29.21 0 1 0 19.03 55.23z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1100.98 552.17L1198.92 552.17L1198.92 552.84L1100.98 552.84z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1534.08 517.77L1614.1 517.77L1614.1 597.79L1534.08 597.79z' stroke='%23d3b714'%3e%3c/path%3e%3cpath d='M396.05 760.81 a26.35 26.35 0 1 0 52.7 0 a26.35 26.35 0 1 0 -52.7 0z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M582.41 258.74L652.36 258.74L652.36 354.65L582.41 354.65z' fill='%23e73635'%3e%3c/path%3e%3cpath d='M27.1 96.53a41.94 41.94 0 1 0-31.96-77.55z' fill='%23037b0b'%3e%3c/path%3e%3cpath d='M1585.84 476.77a78.03 78.03 0 1 0 121.35 98.13z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1139.43 508.17L1139.94 508.17L1139.94 578.58L1139.43 578.58z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M659.63 59.8a67.89 67.89 0 1 0 65.92 118.7z' stroke='%23037b0b'%3e%3c/path%3e%3cpath d='M449.84 98.1 a43.74 43.74 0 1 0 87.48 0 a43.74 43.74 0 1 0 -87.48 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M746.29 258.6 a81.09 81.09 0 1 0 162.18 0 a81.09 81.09 0 1 0 -162.18 0z' fill='%23d3b714'%3e%3c/path%3e%3cpath d='M1370.87 526.01 a81.86 81.86 0 1 0 163.72 0 a81.86 81.86 0 1 0 -163.72 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M147.8 755.43 a78.97 78.97 0 1 0 157.94 0 a78.97 78.97 0 1 0 -157.94 0z' stroke='%23e73635'%3e%3c/path%3e%3cpath d='M1794.71 84.39L1898.8 84.39L1898.8 96.33L1794.71 96.33z' fill='%23e73635'%3e%3c/path%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1008'%3e%3crect width='1980' height='1080' fill='white'%3e%3c/rect%3e%3c/mask%3e%3c/defs%3e%3c/svg%3e");
	}
</style>