<script lang="ts">

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
			border-b-2 border-gray-200

          "
			>
				<div class=" h-15 w-12 object-contain">
					<a href="/">
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

</style>
