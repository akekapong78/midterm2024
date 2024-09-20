'use client'

import axios from 'axios';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react'

type Props = {}

const LoginPage = (props: Props) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (username === '' || password === '') {
        throw new Error('username or password are empty');
      }

      const res = await axios.post(`${process.env.NEXT_PUBLIC_URL_BACKEND}/api/v1/users/login`, {
        username,
        password
      }, {
        withCredentials: true
      });
      
      console.log(res);

      if (res.status === 200) {
        router.push('/');
      }

    } catch (error: any) {
      console.log(error);
      if (error.response.status !== 200) {
        setError(error.response.data.message);
        return
      }
      setError(error.message);
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-lg max-w-md w-full">
        <h1 className="text-2xl font-semibold mb-6 text-center">Login</h1>
        
        <form onSubmit={handleLogin}>
          <div className="mb-4">
            <label htmlFor="username" className="block text-sm font-medium text-gray-700">
              Username / Email
            </label>
            <input
              placeholder='example@mail.com'
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="mt-1 block w-full p-2 border rounded-lg focus:ring-indigo-500 focus:border-indigo-500 border-gray-300 text-white"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password
            </label>
            <input
              placeholder='password123'
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mt-1 block w-full p-2 border rounded-lg focus:ring-indigo-500 focus:border-indigo-500 border-gray-300 text-white"
              required
            />
          </div>
          {error && <p className="mt-2 text-red-500 text-sm mb-4">{error}</p>}

          <button
            type="submit"
            className="w-full py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg"
          >
            Login
          </button>
          
          <p className="text-center mt-4">
            Don't have an account? 
            <Link href="/register" className="ml-2 text-blue-500 hover:text-blue-700 hover:underline">
              Sign up
            </Link>
          </p>
        </form>
      </div>
    </div>
  );
}

export default LoginPage