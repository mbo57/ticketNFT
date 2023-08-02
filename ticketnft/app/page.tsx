import { Staff } from "./types";

const apiUrl = "http://api:8000/";


//  SSRでStaff一覧を取得する
async function getStaffs() {
    const url = apiUrl + "staff/show"
    const res = await fetch(url,{cache: 'no-store'})

    if(!res.ok){
        throw new Error(res.statusText)
    }

    return res.json()
}

export default async function Home() {
    // const staffs = (async () => {
    //     const response = await fetch(url)
    //     const json = await response.json()
    //     // console.log(json)
    //     return json
    // })();
    // const staffs = staffsget().then((tmp) => console.log(tmp))
    // const staffs = await getStaffs()
    // const url = endpoint + "staff/show"

    // const staffs = await fetch(url).then((res) => res.json())

    const staffs = await getStaffs()
    console.log(staffs)


    return (
        <main>
            <h1>従業員一覧</h1>
            {/* 修正:border設定方法 */}
            <table border={1}>
                <thead>
                    <tr>
                        <th>
                            staff_id
                        </th>
                        <th>
                            staff_name
                        </th>
                        <th>
                            staff_email
                        </th>
                        <th>
                            staff_password
                        </th>
                    </tr>
                </thead>
                
                <tbody>
                {/* 追加：staff 型宣言・mapのkey設定 */}
                {staffs.map((staff : Staff) => 
                    <tr key={staff.id}>
                        <td>{staff.id}</td>
                        <td>{staff.name}</td>
                        <td>{staff.email}</td>
                        <td>{staff.password}</td>
                    </tr>
                )}
                {/*
                */}
                </tbody>
            </table>
        </main>
    )
}
