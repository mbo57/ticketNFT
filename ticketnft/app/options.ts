import type { NextAuthOptions } from "next-auth";
import GitHubProvider from "next-auth/providers/github";
import GoogleProvider from "next-auth/providers/google";
import CredentialsProvider from "next-auth/providers/credentials";

const endpoint = "http://api:8000/";

type staff = {
    id: int
    email: string
    password: string
    name: string
}


export const options: NextAuthOptions = {
        debug: true,
        session: {strategy: "jwt"},
        providers: [
            CredentialsProvider({
                    name: "email",
                    credentials: {
                        email: {
                            label: "Email",
                            type: "email",
                            placeholder: "example@example.com",
                        },
                        password: {label: "Password", type: "password"},
                    },
                    // メルアド認証処理
                    async authorize(credentials) {
                        const url = endpoint + "staff/show?email=" + credentials?.email
                        console.log(url)
                        const users = await fetch(url).then((r) => r.json())
                        if (users == null) {
                            return null;
                        }
                        const user = users[0]
                        // const users = [
                        //     {id: "1", email: "user1@example.com", password: "password"},
                        //     {id: "2", email: "user2@example.com", password: "password2"},
                        //     {id: "3", email: "abc@abc", password: "123"},
                        // ];
                        // const user = users.find(user => user.email === credentials?.email);

                        if (user && user?.password === credentials?.password) {
                            // console.log("+++++++++++sucsess");
                            return {id: user.id, name: user.email, email: user.email, role: "admin"};
                        } else {
                            // console.log("+++++++++++fail");
                            return null;
                        }
                    }
                }
            ),
            GitHubProvider({
                clientId: process.env.GITHUB_ID!,
                clientSecret: process.env.GITHUB_SECRET!,
            }),
            GoogleProvider({
                clientId: process.env.GOOGLE_CLIENT_ID!,
                clientSecret: process.env.GOOGLE_CLIENT_SECRET!
            }),
        ],
        callbacks: {
            jwt: async ({token, user, account, profile, isNewUser}) => {
                if (user) {
                    token.user = user;
                    const u = user as any
                    token.role = u.role;
                }
                if (account) {
                    token.accessToken = account.access_token
                }
                return token;
            },
            session: ({session, token}) => {
                token.accessToken
                return {
                    ...session,
                    user: {
                        ...session.user,
                        role: token.role,
                    },
                };
            },
        }
    }
;
