// "use client"
import {
    LoginButton,
    LogoutButton,
} from "@/app/components/buttons";

import { getServerSession } from "next-auth/next"
import { options } from "@/app/options";

export default async function Home() {
    const session = await getServerSession(options)
    console.log(session)
    return session? (
        <main
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                height: "70vh",
            }}
        >
            <div>
                <LogoutButton/>
            </div>
        </main>
    ) : (
        <main
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                height: "70vh",
            }}
        >
        {/*
            <div>
                <SignupButton/>
            </div>
        */}
            <div>
                <LoginButton/>
            </div>
        </main>
    );
}
