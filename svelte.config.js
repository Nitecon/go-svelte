import adapter from '@sveltejs/adapter-static';

const config = {
    kit: {
        adapter: adapter({
            pages: 'static',  // destination for pages
            assets: 'static',  // destination for assets
            fallback: null,
        })
    }
};

export default config;
