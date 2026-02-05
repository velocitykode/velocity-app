import { useEffect } from 'react';
import { usePage } from '@inertiajs/react';

const appName = 'Velocity App';

export function usePageTitle(title?: string) {
    const page = usePage();

    useEffect(() => {
        const pageTitle = title || (page.props as any).title || '';
        document.title = pageTitle ? `${pageTitle} - ${appName}` : appName;
    }, [title, page]);
}
