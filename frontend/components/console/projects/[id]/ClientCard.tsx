import Link from "next/link";
import { DefaultText, LargeText, TinyText } from "@/components/common/Texts";

type ClientCardProps = {
  id: string;
  name: string;
  clientId: string;
  clientSecret: string;
};

export default function ClientCard({
  id,
  name,
  clientId,
  clientSecret,
}: ClientCardProps) {
  return (
    <Link href={`/console/clients/${id}`}>
      <div className='p-4 flex flex-col bg-slate-900 rounded-xl'>
        <LargeText name>{name}</LargeText>
        <div className='py-2 flex'>
          <div className='w-1/2'>
            <TinyText>Client id</TinyText>
            <DefaultText>{clientId}</DefaultText>
          </div>
          <div className='w-1/2'>
            <TinyText>Client type</TinyText>
            <DefaultText>{clientSecret}</DefaultText>
          </div>
        </div>
      </div>
    </Link>
  );
}
