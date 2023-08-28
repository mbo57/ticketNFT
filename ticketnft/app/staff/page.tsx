"use client";
import * as React from "react";

import Image from "next/image";

import { HeaderForStaff } from "@/app/components/staff/header";
import { DeleteModal, EditModal, NewModal } from "@/app/components/staff/modal";
import { Footer } from "@/app/components/footer";

import {
    NextUIProvider,
    Table,
    TableHeader,
    TableColumn,
    TableBody,
    TableRow,
    TableCell,
    Link,
    Avatar,
} from "@nextui-org/react";

import Add from "../../public/icon/add.png";
import Vuesax from "../../public/icon/vuesax.png";
import Edit from "../../public/icon/edit.png";
import Trash from "../../public/icon/trash.png";

export default async function StaffPageTop() {
    return (
        <NextUIProvider>
            <HeaderForStaff />

            <main className="mj-container">
                <section>
                    <h1 className="text-2xl mt-8 mb-4 font-bold">
                        イベント一覧
                        <NewModal />
                    </h1>

                    <Table
                        removeWrapper
                        aria-label="Example static collection table"
                    >
                        <TableHeader>
                            <TableColumn className="w-[64px]"> </TableColumn>
                            <TableColumn>イベント名</TableColumn>
                            <TableColumn>日時</TableColumn>
                            <TableColumn>場所</TableColumn>
                            <TableColumn>出演者</TableColumn>
                            <TableColumn>イベントカテゴリ</TableColumn>
                            <TableColumn>備考</TableColumn>
                            <TableColumn className=""> </TableColumn>
                        </TableHeader>
                        <TableBody>
                            <TableRow key="1">
                                <TableCell>
                                    <Avatar
                                        radius="md"
                                        src="https://material.onlineticket.jp/s/image/025431/0001/000/0254310001_3.jpg"
                                    />
                                </TableCell>
                                <TableCell className="font-bold">
                                    SUPER BEAVER 都会のラクダ TOUR 2023-2024 〜
                                    駱駝革命21 〜
                                </TableCell>
                                <TableCell>2023/09/29</TableCell>
                                <TableCell>広島文化学園HBGホール</TableCell>
                                <TableCell>SUPER BEAVER</TableCell>
                                <TableCell>音楽ライブ</TableCell>
                                <TableCell>
                                    SUPER BEAVERのライブ。SUPER
                                    BEAVERの個人的に一番好きな曲は、証明です。
                                </TableCell>
                                <TableCell>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        詳細
                                        <Image
                                            src={Vuesax}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <EditModal />
                                    <DeleteModal />
                                </TableCell>
                            </TableRow>
                            <TableRow key="2">
                                <TableCell>
                                    <Avatar
                                        radius="md"
                                        src="https://material.onlineticket.jp/s/image/025431/0001/000/0254310001_3.jpg"
                                    />
                                </TableCell>
                                <TableCell className="font-bold">
                                    SUPER BEAVER 都会のラクダ TOUR 2023-2024 〜
                                    駱駝革命21 〜
                                </TableCell>
                                <TableCell>2023/09/29</TableCell>
                                <TableCell>広島文化学園HBGホール</TableCell>
                                <TableCell>SUPER BEAVER</TableCell>
                                <TableCell>音楽ライブ</TableCell>
                                <TableCell>
                                    SUPER BEAVERのライブ。SUPER
                                    BEAVERの個人的に一番好きな曲は、証明です。
                                </TableCell>
                                <TableCell>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        詳細
                                        <Image
                                            src={Vuesax}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        編集
                                        <Image
                                            src={Edit}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link isExternal href="#" color="danger">
                                        削除
                                        <Image
                                            src={Trash}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                </TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </section>

                <section>
                    <h1 className="text-2xl mt-8 mb-4 font-bold">
                        出演者一覧
                        <Link
                            isExternal
                            href="#"
                            color="foreground"
                            className="ml-2"
                        >
                            出演者追加
                            <Image
                                src={Add}
                                alt=""
                                className="h-[16px] w-[16px]"
                            />
                        </Link>
                    </h1>

                    <Table
                        removeWrapper
                        aria-label="Example static collection table"
                    >
                        <TableHeader>
                            <TableColumn className="w-[64px]"> </TableColumn>
                            <TableColumn>出演者名</TableColumn>
                            <TableColumn>備考</TableColumn>
                            <TableColumn className="w-[150px]"> </TableColumn>
                        </TableHeader>
                        <TableBody>
                            <TableRow key="1">
                                <TableCell>
                                    <Avatar
                                        radius="md"
                                        src="https://material.onlineticket.jp/s/image/025431/0001/000/0254310001_3.jpg"
                                    />
                                </TableCell>
                                <TableCell className="font-bold">
                                    SUPER BEAVER
                                </TableCell>
                                <TableCell>
                                    SUPER BEAVER。SUPER
                                    BEAVERの個人的に一番好きな曲は、証明です。
                                </TableCell>
                                <TableCell>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        詳細
                                        <Image
                                            src={Vuesax}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        編集
                                        <Image
                                            src={Edit}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link isExternal href="#" color="danger">
                                        削除
                                        <Image
                                            src={Trash}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                </TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </section>

                <section>
                    <h1 className="text-2xl mt-8 mb-4 font-bold">
                        イベントカテゴリ一覧
                        <Link
                            isExternal
                            href="#"
                            color="foreground"
                            className="ml-2"
                        >
                            イベントカテゴリ追加
                            <Image
                                src={Add}
                                alt=""
                                className="h-[16px] w-[16px]"
                            />
                        </Link>
                    </h1>

                    <Table
                        removeWrapper
                        aria-label="Example static collection table"
                    >
                        <TableHeader>
                            <TableColumn className="w-[64px]"> </TableColumn>
                            <TableColumn>イベントカテゴリ名</TableColumn>
                            <TableColumn className="w-[150px]"> </TableColumn>
                        </TableHeader>
                        <TableBody>
                            <TableRow key="1">
                                <TableCell>
                                    <Avatar
                                        radius="md"
                                        src="https://media.istockphoto.com/id/1340928148/ja/%E3%83%93%E3%83%87%E3%82%AA/%E3%82%B3%E3%83%B3%E3%82%B5%E3%83%BC%E3%83%88%E3%81%AB%E6%84%9F%E8%AC%9D%E3%81%97%E8%88%88%E5%A5%AE%E3%81%97%E3%81%9F%E8%81%B4%E8%A1%86%E9%9F%B3%E6%A5%BD%E3%82%B3%E3%83%B3%E3%82%B5%E3%83%BC%E3%83%88%E3%83%9F%E3%83%A5%E3%83%BC%E3%82%B8%E3%82%B7%E3%83%A3%E3%83%B3%E9%9F%B3%E6%A5%BD%E7%A5%AD.jpg?s=640x640&k=20&c=el0rSbTEk9t31kdx_RM6Nc0zjXxKXW74mO7dU5lIisM="
                                    />
                                </TableCell>
                                <TableCell className="font-bold">
                                    音楽ライブ
                                </TableCell>
                                <TableCell>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        編集
                                        <Image
                                            src={Edit}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link isExternal href="#" color="danger">
                                        削除
                                        <Image
                                            src={Trash}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                </TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </section>

                <section>
                    <h1 className="text-2xl mt-8 mb-4 font-bold">
                        スタッフ一覧
                        <Link
                            isExternal
                            href="#"
                            color="foreground"
                            className="ml-2"
                        >
                            スタッフ追加
                            <Image
                                src={Add}
                                alt=""
                                className="h-[16px] w-[16px]"
                            />
                        </Link>
                    </h1>

                    <Table
                        removeWrapper
                        aria-label="Example static collection table"
                    >
                        <TableHeader>
                            <TableColumn className="w-[64px]"> </TableColumn>
                            <TableColumn>スタッフ名</TableColumn>
                            <TableColumn>Email</TableColumn>
                            <TableColumn>PassWord</TableColumn>
                            <TableColumn>Roll</TableColumn>
                            <TableColumn className="w-[150px]"> </TableColumn>
                        </TableHeader>
                        <TableBody>
                            <TableRow key="1">
                                <TableCell>
                                    <Avatar
                                        radius="md"
                                        src="https://images.unsplash.com/broken"
                                    />
                                </TableCell>
                                <TableCell className="font-bold">
                                    スタッフA
                                </TableCell>
                                <TableCell>staff@email.com</TableCell>
                                <TableCell>**</TableCell>
                                <TableCell>管理者</TableCell>
                                <TableCell>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        編集
                                        <Image
                                            src={Edit}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link isExternal href="#" color="danger">
                                        削除
                                        <Image
                                            src={Trash}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                </TableCell>
                            </TableRow>
                            <TableRow key="1">
                                <TableCell>
                                    <Avatar
                                        radius="md"
                                        src="https://images.unsplash.com/broken"
                                    />
                                </TableCell>
                                <TableCell className="font-bold">
                                    スタッフB
                                </TableCell>
                                <TableCell>staff@email.com</TableCell>
                                <TableCell>**</TableCell>
                                <TableCell>一般</TableCell>
                                <TableCell>
                                    <Link
                                        isExternal
                                        href="#"
                                        color="foreground"
                                    >
                                        編集
                                        <Image
                                            src={Edit}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                    <Link isExternal href="#" color="danger">
                                        削除
                                        <Image
                                            src={Trash}
                                            alt=""
                                            className="h-[16px] w-[16px]"
                                        />
                                    </Link>
                                </TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </section>
            </main>

            <Footer />
        </NextUIProvider>
    );
}
