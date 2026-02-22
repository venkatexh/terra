import useAxios from "@/hooks/useAxios";

export default function ConsentView() {
  const axios = useAxios();

  async function allow() {
    await axios.post("/oauth/authorize/approve");
  }

  async function deny() {
    await axios.post("/oauth/authorize/deny");
  }

  return (
    <div className='max-w-96 mx-auto'>
      <div className='text-xl'>
        TripPlanner wants access to your account details
      </div>

      {/* <ul>
        {client.scopes.map((s: string) => (
          <li key={s}>{s}</li>
        ))}
      </ul> */}
      <div className='py-12'>
        <div>Allowing this will give TripPlanner permission to</div>
        <div>
          <ul className='py-4'>
            <li className='flex gap-2 items-center justify-start'>
              <input type='checkbox' />
              <div>View and save your email</div>
            </li>
            <li className='flex gap-2 items-center justify-start'>
              <input type='checkbox' />
              <div>View and save your name</div>
            </li>
          </ul>
        </div>
        <p>
          You can revoke this access at any time by visiting your Terra account
          settings.
        </p>
      </div>
      <div className='py-4 flex justify-end gap-2'>
        <button
          className='px-4 py-1 rounded-sm bg-orange-600 hover:bg-orange-700 cursor-pointer'
          onClick={allow}>
          Allow
        </button>
        <button
          className='px-4 py-1 rounded-sm bg-gray-300 hover:bg-gray-400 text-gray-600 cursor-pointer'
          onClick={deny}>
          Deny
        </button>
      </div>
    </div>
  );
}
