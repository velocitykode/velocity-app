import tailwindcss from '@tailwindcss/vite';
import path from 'path';
import { defineConfig } from 'vite';

export default defineConfig({
    plugins: [tailwindcss()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './resources/js'),
        },
    },
    build: {
        outDir: 'public/build',
        manifest: true,
        rollupOptions: {
            input: 'resources/js/app.tsx',
        },
    },
    server: {
        port: 5173,
        strictPort: true,
        host: 'localhost',
    },
    esbuild: {
        jsx: 'automatic',
        jsxDev: false,
    },
});
