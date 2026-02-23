export default function AuthorizedAppCard({id, name}: {id: string, name: string}) {
  return <div className="p-4 bg-slate-900 rounded-xl">
    <div>{name}</div>
  </div>
}