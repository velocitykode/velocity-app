import '../css/app.css';

import React from 'react';
import { createInertiaApp, router } from '@inertiajs/react';
import { createRoot } from 'react-dom/client';
import { initializeTheme } from './hooks/use-appearance';
import axios from 'axios';

// Store for the current CSRF token (updated on every Inertia response)
let csrfToken: string | null = null;

// Configure axios to use the current CSRF token
axios.interceptors.request.use((config) => {
    if (csrfToken) {
        config.headers['X-CSRF-Token'] = csrfToken;
    }
    return config;
});

// Update CSRF token on every Inertia navigation
router.on('navigate', (event) => {
    const pageProps = event.detail.page.props as { csrf_token?: string };
    if (pageProps.csrf_token) {
        csrfToken = pageProps.csrf_token;
        // Also update meta tag for any non-Inertia forms
        const meta = document.querySelector('meta[name="csrf-token"]');
        if (meta) {
            meta.setAttribute('content', csrfToken);
        }
    }
});

createInertiaApp({
    resolve: async (name) => {
        const pages = import.meta.glob('./pages/**/*.tsx', { eager: true }) as Record<string, { default: React.ComponentType }>;
        const page = pages[`./pages/${name}.tsx`];
        return page.default;
    },
    setup({ el, App, props }) {
        // Initialize CSRF token from initial page props
        const initialProps = props.initialPage.props as { csrf_token?: string };
        if (initialProps.csrf_token) {
            csrfToken = initialProps.csrf_token;
        }
        const root = createRoot(el);
        root.render(<App {...props} />);
    },
    progress: {
        color: '#4B5563',
    },
});

// Initialize dark mode
initializeTheme();
