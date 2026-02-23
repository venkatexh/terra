import useAxios from "@/hooks/useAxios";
import { useState } from "react";

export default function LoginView() {
  const axios = useAxios();
  const [email, setEmail] = useState("");
  const [sent, setSent] = useState(false);

  async function submit() {
    await axios.post("/auth/magic-link/request", {
      email,
      return_to: window.location.href,
    });

    setSent(true);
  }

  if (sent) {
    return <div>Check your email for the login link.</div>;
  }

  return (
    <div className='max-w-96 mx-auto'>
      <div className='text-xl'>Log in to continue</div>
      <form className='flex flex-col gap-4 py-12 '>
        <input
          placeholder='Your email'
          value={email}
          className='p-2 border border-gray-300 rounded-3xl focus:outline-none'
          onChange={(e) => setEmail(e.target.value)}
        />
        <button
          className='bg-orange-600 p-2 rounded-3xl hover:bg-orange-700 cursor-pointer'
          onClick={submit}>
          Continue
        </button>
      </form>
    </div>
  );
}
