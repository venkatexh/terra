import Link from "next/link";

export default function ProjectCard({
  id,
  name,
}: {
  id: string;
  name: string;
}) {
  return (
    <Link href={`/console/projects/${id}`}>
      <div className='p-4 bg-slate-900 rounded-lg'>{name}</div>
    </Link>
  );
}
