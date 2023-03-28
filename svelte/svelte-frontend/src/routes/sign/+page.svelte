<script lang="ts">
	import NavBar from '../../lib/NavBar.svelte';
	let login = false;
	export let form;
	if(form){

	}
</script>

<div>
	<NavBar {login} />
</div>

<div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
	<div class="w-full max-w-md border-2 border-slate-700 space-y-8 p-4 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
		<div class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-700">
				Sign up for a new account!
			<h2 class="mt-6 text-center text-2xl font-bold tracking-tight text-gray-100">
			
			</h2>
			<!-- <p class="mt-2 text-center text-sm text-gray-600">
          Or
          <a href="#" class="font-medium text-gray-600 hover:text-gray-500">start your 14-day free trial</a>
        </p> -->
		</div>
		<form class="mt-8 space-y-6" method="POST">
			<input type="hidden" name="remember" value="true" />
			<div class="-space-y-px rounded-md shadow-sm">
				<div>
					<label for="name" class="sr-only">Name</label>
					<input
						id="name"
						name="name"
						required
						class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-gray-500 focus:outline-none focus:ring-gray-500 sm:text-sm"
						placeholder="Name"
					/>
				</div>
				<div>
					<label for="email-address" class="sr-only">Email address</label>
					<input
						id="email-address"
						name="email"
						type="email"
						autocomplete="email"
						required
						class="relative block w-full appearance-none rounded-none border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-gray-500 focus:outline-none focus:ring-gray-500 sm:text-sm"
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
						class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-gray-500 focus:outline-none focus:ring-gray-500 sm:text-sm"
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
						class="h-4 w-4 rounded border-gray-300 text-gray-600 focus:ring-gray-500"
					/> -->
					<!-- <label for="remember-me" class="ml-2 block text-sm text-gray-900">Remember me</label> -->
				<!-- </div> -->
				<!-- probebly dont need password recovery -->

				<!-- <div class="text-sm">
					<a href="../notes" class="font-medium text-gray-600 hover:text-gray-500"
						>Forgot your password?</a
					>
				</div> -->
			<!-- </div> -->

			<div>
				<button
					type="submit"
					class="group relative flex w-full justify-center rounded-md border border-transparent bg-blue-500 py-2 px-4 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
				>
					<span class="absolute inset-y-0 left-0 flex items-center pl-3">
						<!-- Heroicon name: mini/lock-closed -->
						<svg
							class="h-5 w-5 group-hover:text-slate-200"
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
					Sign Up
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
