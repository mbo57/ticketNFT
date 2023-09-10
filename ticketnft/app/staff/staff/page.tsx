"use client";
import React, { useState, useEffect } from "react";

export default function StaffStaff() {
    const [staffs, setStaffs] = useState([])

    useEffect(() => {
        const fetchStaffs = async () => {
            const response = await fetch("http://localhost:3000/api/staff");
            const data = await response.json();
            setStaffs(data.staffs);
        };
        fetchStaffs();
    }, []);

    console.log("レンダリング");
    console.log(staffs);
    
    return (
        <>
            {staffs.map((staff) => (
                <p key={staff.id}>{staff.email}</p>
            ))}
            <>{}</>
        </>
    );
}
