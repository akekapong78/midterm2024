'use client'
import axios from 'axios';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react'

type Props = {}

const RegisterPage = (props: Props) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault(); // Prevent page refresh
    try {
      if (username === '' || password === '' || role === '') {
        throw new Error('username, password, or role are empty');
      }

      const res = await axios.post(`${process.env.NEXT_PUBLIC_URL_BACKEND}/api/v1/users/register`, {
        username,
        password,
        role
      });
      console.log(res);
      router.push('/login');

    } catch (error: any) {
      console.log(error);
      if (error.response.status === 500) {
        setError(error.response.data.message);
        return
      }
      setError(error.message);
    }
  };

  // Clear error message
  useEffect(() => {
    setError('');
  }, [username, password, role]);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-lg max-w-md w-full">
        <h1 className="text-2xl font-semibold mb-6 text-center">Register</h1> 
        <form 
          onSubmit={handleRegister}
        >
          <div className="mb-4">
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
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
          <div className="mb-4">
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Roles
            </label>
            <select 
              id='role'
              className="mt-1 select select-info w-full block p-2 border focus:ring-indigo-500 focus:border-indigo-500 border-gray-300 text-white"
              onChange={(e) => setRole(e.target.value)}
              // defaultValue={''}
              required
            >
              <option disabled selected>Select roles</option>
              <option value={'USER'}>User</option>
              <option value={'ADMIN'}>Admin</option>
            </select>
          </div>
          {error && <p className="text-red-500 text-sm mb-4">{error}</p>}
          <button
            // onClick={handleRegister}
            type="submit"
            className="mt-5 w-full py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg"
          >
            Submit
          </button>

          <button
            type="button"
            className="mt-2 w-full py-2 px-4 bg-gray-600 hover:bg-gray-700 text-white rounded-lg"
            onClick={() => router.push('/login')}
          >
            Cancel
          </button>
          
        </form>
      </div>
    </div>
  );
}

export default RegisterPage