import Link from "next/link";
import { DefaultText } from "./common/Texts";

export default function Console() {
  return (
    <Link href='/console/projects' className='w-1/2'>
      <div className='p-4 bg-slate-900 rounded-xl'>
        <DefaultText>Developer console</DefaultText>
      </div>
    </Link>
  );
}
