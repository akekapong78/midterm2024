import { cookies } from 'next/headers'
import { UserProvider } from './userProvider'

export default function ServerUserProvider({ children }: { children: React.ReactNode }) {
  const cookieStore = cookies()
  const username = cookieStore.get('username')?.value || null
  const role = cookieStore.get('role')?.value || null
  const user = { username, role }

  return <UserProvider user={user}>{children}</UserProvider>
}