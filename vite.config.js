import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit()],
    proxy: {
        '/api': {
            target: 'http://localhost:8888',
            changeOrigin: true,
            secure: false,
            ws: true,
        }
    }
};

export default config;
