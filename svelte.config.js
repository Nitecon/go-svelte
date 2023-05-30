import adapter from '@sveltejs/adapter-static';

const config = {
    kit: {
        prerender: {
            handleHttpError: ({ path, referrer, message }) => {
                // ignore api servers
                if (path.startsWith('/api')) {
                    return;
                }

                // otherwise fail the build
                throw new Error(message);
            }
        },
        adapter: adapter({
            pages: 'static',  // destination for pages
            assets: 'static',  // destination for assets
            fallback: null,
        })
    },
};

export default config;
