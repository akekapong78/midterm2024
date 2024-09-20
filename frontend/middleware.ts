import { NextRequest, NextResponse } from 'next/server'
import { decrypt } from '@/app/lib/session'
import { cookies } from 'next/headers'
 
// 1. Specify protected and public routes
const protectedRoutes = [ '/users', '/items']
const publicRoutes = ['/login', '/register']
 
export default async function middleware(req: NextRequest) {
  // 2. Check if the current route is protected or public
  const path = req.nextUrl.pathname
  const isRootRoute = path === '/' // Special case for root '/'
  const isProtectedRoute = isRootRoute || protectedRoutes.some((route) => path.startsWith(route))
  const isPublicRoute = publicRoutes.includes(path)

  // 3. Decrypt the session from the cookie
  const cookie = cookies().get('token')?.value
  const token = cookie?.split('Bearer+')[1]; // Extracts the token part
  // console.log("token: ", token)
  const session = await decrypt(token)
  console.log("session: ", session?.aud)
  
  // 5. Redirect to /login if the user is not authenticated
  if (isProtectedRoute && !session) {
    return NextResponse.redirect(new URL('/login', req.nextUrl))
  }

  const username = session?.aud![0] || ''
  const role = session?.aud![1] || ''
  
  // 6. Redirect to / if the user is authenticated
  if (isPublicRoute && username) {
    return NextResponse.redirect(new URL('/', req.nextUrl))
  }
 
  // Set the userId in a cookie or header to make it accessible in the app
  const response = NextResponse.next()
  response.cookies.set('username', username)
  response.cookies.set('role', role)

  return response
}
 
// Routes Middleware should not run on
export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}