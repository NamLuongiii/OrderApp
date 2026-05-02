import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { pathToRegexp } from 'path-to-regexp'

// const PUBLIC_SCREENS = ['/login', '/', '/order/:id', '/order/new'];
const PROTECTED_SCREENS = ['/orders', 'products', 'orders/:id'];

const checkMatch = (pathname: string, patterns: string[]) => {
    return patterns.some((pattern) => {
        const regexp = pathToRegexp(pattern);
        return regexp.test(pathname);
    });
};

export function proxy(request: NextRequest) {
    const token = request.cookies.get('auth-token')?.value;
    const { pathname } = request.nextUrl;

    // const isPublicScreen = checkMatch(pathname, PUBLIC_SCREENS);
    const isProtectedScreen = checkMatch(pathname, PROTECTED_SCREENS);

    if (isProtectedScreen && !token) {
        const loginUrl = new URL('/login', request.url);
        loginUrl.searchParams.set('from', pathname);
        return NextResponse.redirect(loginUrl);
    }

    if (token && pathname === '/login') {
        return NextResponse.redirect(new URL('/orders', request.url));
    }

    return NextResponse.next();
}

export const config = {
    matcher: [
        '/((?!api|_next/static|_next/image|favicon.ico).*)',
    ],
}