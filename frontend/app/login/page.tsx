"use client";

import { useState } from "react";
import useAxios from "@/hooks/useAxios";
import { useRouter } from "next/navigation";

import { baseURL } from "@/lib/constants/baseURL";

import { FormInput } from "@/components/common/Inputs";
import { DefaultButton } from "@/components/common/Buttons";

export default function LoginPage() {
  const axios = useAxios();
  const router = useRouter();

  const [otp, setOtp] = useState("");
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const [otpSent, setOtpSent] = useState(false);
  const [linkSent, setLinkSent] = useState(false);

  const ERR_MESSAGE = "Something went wrong. Please try again.";

  const requestMagicLink = async () => {
    try {
      const url = `${baseURL}/auth/magic-link/request`;
      const res = await axios.post(url, {
        email,
      });
      if (res.status !== 200) {
        setError(ERR_MESSAGE);
      } else {
        setEmail("");
        setLinkSent(true);
      }
    } catch (err) {
      setError(ERR_MESSAGE);
      console.error(err);
    }
  };

  const requestOTP = async () => {
    try {
      const url = `${baseURL}/auth/otp/request`;
      const res = await axios.post(url, {
        email,
      });
      if (res.status !== 200) {
        setError(ERR_MESSAGE);
      } else {
        setEmail("");
        setOtpSent(true);
      }
    } catch (err) {
      setError(ERR_MESSAGE);
      console.error(err);
    }
  };

  const verifyOtp = async () => {
    try {
      const url = `${baseURL}/auth/otp/verify?otp=${otp}`;
      const res = await axios.post(url);
      if (res.status !== 200) {
        setError(ERR_MESSAGE);
      } else {
        setEmail("");
        setOtpSent(false);
        router.replace("/");
      }
    } catch (err) {
      setError(ERR_MESSAGE);
      console.error(err);
    }
  };

  return (
    <div className='w-full h-full flex justify-center items-center'>
      {linkSent && <div>Magic link sent. Please check your email.</div>}
      {otpSent && (
        <div className='w-full mx-auto flex flex-col gap-2 items-center'>
          <FormInput
            type='text'
            name='otp'
            value={otp}
            placeholder='Enter OTP'
            className='w-72 mx-auto px-2 py-2 border border-gray-500 focus:outline-none rounded-3xl'
            onChange={(e) => setOtp(e.target.value)}
          />
          <DefaultButton
            className='w-72 py-2 rounded-full'
            handleButtonClick={() => {
              verifyOtp();
            }}>
            Verify OTP
          </DefaultButton>
        </div>
      )}
      {!linkSent && !otpSent && (
        <div className='max-w-72'>
          <div className='flex flex-col gap-2'>
            <input
              type='email'
              placeholder='Your email'
              className='w-72 px-2 py-2  border border-gray-700 focus:outline-none rounded-3xl'
              onChange={(e) => setEmail(e.target.value)}
            />
            <DefaultButton
              className='py-2 rounded-full'
              handleButtonClick={() => {
                requestMagicLink();
              }}>
              Request Magic Link
            </DefaultButton>
            <DefaultButton
              className='py-2 rounded-full'
              handleButtonClick={() => {
                requestOTP();
              }}>
              Request OTP
            </DefaultButton>
          </div>
          <p className='max-w-72 text-sm/4 py-4 text-red-600'>{error}</p>
        </div>
      )}
    </div>
  );
}
