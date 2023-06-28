const url = "http://api:8000/";
export default async function Home() {
  const json = await fetch(url).then((r) => r.json());

  return (
    <main>
      <h1>従業員一覧</h1>
      <h2>{json[0].Name}</h2>
      <h2>{json[1].Name}</h2>
    </main>
  )
}
