import Link from "next/link";

export default function GroupCard({ id, name }: { id: string; name: string }) {
  return (
    <Link href={`/app-groups/${id}`}>
      <div className='p-4 bg-slate-900 rounded-lg'>{name}</div>
    </Link>
  );
}
