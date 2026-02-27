import { DefaultText, TinyText } from "@/components/common/Texts";
import Link from "next/link";

export default function ProjectCard({
  id,
  name,
  description,
}: {
  id: string;
  name: string;
  description: string;
}) {
  return (
    <Link href={`/console/projects/${id}`}>
      <div className='p-4 bg-slate-900 rounded-lg'>
        <DefaultText name>{name}</DefaultText>
        <DefaultText>{description}</DefaultText>
        <div className="py-4">
          <TinyText>Project ID</TinyText>
          <DefaultText>{id}</DefaultText>
        </div>
      </div>
    </Link>
  );
}
