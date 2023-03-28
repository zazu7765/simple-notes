<script lang="ts">

	// import { onMount } from "svelte";

	// onMount(() => {

	//    return document.querySelector('#menu')
	let menu = 'hidden';
	export let login: boolean;

		let hid = "hidder";
	

	// });

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
				
			bg-[#]
          fixed-top
            flex flex-wrap
            items-center
            justify-between
            w-full
            
            md:py-0
            px-4
            text-lg text-gray-700

			

          "
			>
				<div class="link  md:p-2 my-4 block] rounded-md  text-slate-900 object-contain">
					<a href="/">
					<b>SIMPLE NOTES</b>
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
                pt-1
                text-base text-gray-700
                md:flex
                md:justify-between 
                md:pt-0"
					>
						{#if !login}
							<li>
								<a class="md:px-4 py-1 my-1 mr-3 block  border-2 px-4 border-slate-500 rounded-md  text-slate-800" href="/login"><b>Login</b></a>
							</li>
							<!-- <li> -->
								<!-- <a  class="md:p-4 py-2 my-2 block hover:border-b-2  px-2 bg-[#E9A178]  rounded-md  text-slate-100" href="/sign"
									><b>Sign Up</b></a
								> -->
							<!-- </li> -->
						{:else}
							<li>
								<a  class="md:p-4 py-2 my-4 block mr-3  border-2  px-2 border-slate-500 rounded-md  text-slate-800" href="/user/notebooks"
									><b>Open Dashboard</b></a
								>
              </li>
              <li>
			
								<a class= "md:p-4 py-2 my-4 block  px-2   border-2 border-slate-500  rounded-md  text-slate-800" href="/logout"
									><b>Logout</b></a
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
:global(body){
	background-color: #ECF2FF;
}

    .link {
  position: relative;
}
.link:after {
  content: '';
  position: absolute;
  bottom: -.4em;
  left: 50%;
  right: 50%;
  height: 1px;
  background: currentColor;
  -webkit-transition: all ease .2s;
  transition: all ease .2s;
}
.link:hover:after {
  left: 0;
  right: 0;
  height: 2px;
}
.link:before {
  position: absolute;
  -webkit-transform: translateX(-100%);
          transform: translateX(-100%);
  left: -10px;
  color: #eee;
  opacity: .3;
  font-weight: 100;
  font-size: .8em;
}

</style>
