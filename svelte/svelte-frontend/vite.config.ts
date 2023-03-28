import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';


const config: UserConfig = {
	plugins: [sveltekit() ],
	server: {
		watch: {
		  usePolling: true,
		},
		host: true, // needed for the Docker Container port mapping to work
		strictPort: true,
		port: 5173,

},

};

export default config;
