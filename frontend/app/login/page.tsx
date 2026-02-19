"use client";

import { useState } from "react";
import useAxios from "@/hooks/useAxios";
import { baseURL } from "@/lib/constants/baseURL";

export default function LoginPage() {
  const axios = useAxios();
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const [emailSent, setEmailSent] = useState(false);

  const handleFormSubmit = async (e: React.SubmitEvent) => {
    e.preventDefault();
    try {
      const res = await axios.post(`${baseURL}/auth/magic-link/request`, {
        email,
      });
      if (res.status !== 200) {
        setError("Something went wrong. Please try again.");
      } else {
        setEmail("");
        setEmailSent(true);
      }
    } catch (err) {
      setError("Something went wrong. Please try again.");
      console.error(err);
    }
  };

  return (
    <div className='h-screen flex justify-center items-center'>
      {emailSent ? (
        <div>Magic link sent. Please check your email.</div>
      ) : (
        <div className='max-w-72'>
          <form
            className='flex flex-col gap-4'
            onSubmit={(e) => handleFormSubmit(e)}>
            <input
              type='email'
              placeholder='Your email'
              className='w-72 px-2 py-2  border border-gray-700 focus:outline-none rounded-3xl'
              onChange={(e) => setEmail(e.target.value)}
            />
            <button
              type='submit'
              className='bg-orange-600 hover:bg-orange-700 w-72 py-2 rounded-3xl cursor-pointer'>
              Continue
            </button>
          </form>
          <p className='max-w-72 text-sm/4 py-4 text-red-600'>{error}</p>
        </div>
      )}
    </div>
  );
}
