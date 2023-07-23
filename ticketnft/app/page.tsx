const url = "http://api:8000/";

type staff = {
    id: int
    name: string
}

// async function getStaffs() {
//     const staffs = await fetch(url).then((r) => r.json())
//     return staffs
// }

export default async function Home() {
    // const staffs = (async () => {
    //     const response = await fetch(url)
    //     const json = await response.json()
    //     // console.log(json)
    //     return json
    // })();
    // const staffs = staffsget().then((tmp) => console.log(tmp))
    // const staffs = await getStaffs()
    const staffs = await fetch(url).then((r) => r.json())
    console.log(staffs)


    return (
        <main>
            <h1>従業員一覧</h1>
            <table border="1">
                <thead>
                    <tr>
                        <th>
                            staff_id
                        </th>
                        <th>
                            staff_name
                        </th>
                    </tr>
                </thead>
                <tbody>
                {staffs.map((staff) => 
                    <tr>
                        <td>{staff.id}</td>
                        <td>{staff.name}</td>
                    </tr>
                )}
                {/*
                */}
                </tbody>
            </table>
        </main>
    )
}
