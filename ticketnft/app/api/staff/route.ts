import { NextResponse } from 'next/server'
 
export async function GET() {
  const res = await fetch("http://ticketnft-api-1:8000/staff", {
    headers: {
        'Content-Type': 'application/json',
      },  
  })
  const staffs = await res.json()
  return NextResponse.json({ staffs })
}