import Link from "next/link";

export default function Profiles() {
  return (
    <Link href='/app-groups' className='w-1/2 '>
      <div className='p-4 rounded-xl bg-slate-900'>App profiles</div>
    </Link>
  );
}
