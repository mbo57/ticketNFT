"use client";
import * as React from "react";

import Image from "next/image";

import {
    Link,
    Modal,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalFooter,
    useDisclosure,
    Input,
    Button,
    Textarea,
} from "@nextui-org/react";

import Add from "../../../public/icon/add.png";
import Vuesax from "../../../public/icon/vuesax.png";
import Edit from "../../../public/icon/edit.png";
import Trash from "../../../public/icon/trash.png";


export function DeleteModal() {
    const { isOpen, onOpen, onOpenChange } = useDisclosure();

    return (
        <>
            <Link isExternal onClick={onOpen} color="danger">
                削除
                <Image
                    src={Trash}
                    alt=""
                    className="h-[16px] w-[16px]"
                />
            </Link>

            <Modal
                isOpen={isOpen}
                onOpenChange={onOpenChange}
                isDismissable={false}
            >
                <ModalContent>
                    {(onClose) => (
                        <>
                            <ModalHeader className="flex flex-col gap-1">
                                イベント削除
                            </ModalHeader>
                            <ModalBody>
                                「SUPER BEAVER 都会のラクダ TOUR 2023-2024 〜
                                駱駝革命21 〜」を削除しますか？
                            </ModalBody>
                            <ModalFooter>
                                <Button color="danger">削除する</Button>
                            </ModalFooter>
                        </>
                    )}
                </ModalContent>
            </Modal>
        </>
    );
}



export function EditModal() {
    const { isOpen, onOpen, onOpenChange } = useDisclosure();

    return (
        <>
            <Link isExternal onClick={onOpen} color="foreground">
                編集
                <Image
                    src={Edit}
                    alt=""
                    className="h-[16px] w-[16px]"
                />
            </Link>

            <Modal
                isOpen={isOpen}
                onOpenChange={onOpenChange}
                placement="top-center"
            >
                <ModalContent>
                    {(onClose) => (
                        <>
                            <form action="">
                                <ModalHeader className="flex flex-col gap-1">
                                    イベント編集
                                </ModalHeader>
                                <ModalBody>
                                    <Input
                                        label="イベント名"
                                        variant="bordered"
                                        placeholder=" "
                                    />
                                    <Input
                                        label="開催日"
                                        variant="bordered"
                                        type="date"
                                        placeholder=" "
                                    />
                                    <Input
                                        label="場所"
                                        variant="bordered"
                                        placeholder=" "
                                    />
                                    <div className="mj-select">
                                        <label htmlFor="">
                                            イベントカテゴリ
                                        </label>
                                        <select name="" id="">
                                            <option value="">音楽ライブ</option>
                                            <option value="">公開収録</option>
                                        </select>
                                    </div>
                                    <div className="mj-select">
                                        <label htmlFor="">
                                            出演者
                                        </label>
                                        <select name="" id="">
                                            <option value="">SUPER BEAVER</option>
                                            <option value="">the peggies</option>
                                        </select>
                                    </div>
                                    <Textarea
                                        label="備考"
                                        labelPlacement="inside"
                                        placeholder=" "
                                        defaultValue=""
                                        variant="bordered"
                                    />
                                    <Button className="bg-mj_accent text-white">送信</Button>
                                </ModalBody>
                            </form>
                        </>
                    )}
                </ModalContent>
            </Modal>
        </>
    );
}

export function NewModal() {
    const { isOpen, onOpen, onOpenChange } = useDisclosure();

    return (
        <>
            <Link isExternal onClick={onOpen} color="foreground">
                新規作成
                <Image
                    src={Edit}
                    alt=""
                    className="h-[16px] w-[16px]"
                />
            </Link>

            <Modal
                isOpen={isOpen}
                onOpenChange={onOpenChange}
                placement="top-center"
            >
                <ModalContent>
                    {(onClose) => (
                        <>
                            <form action="">
                                <ModalHeader className="flex flex-col gap-1">
                                    イベント新規作成
                                </ModalHeader>
                                <ModalBody>
                                    <Input
                                        label="イベント名"
                                        variant="bordered"
                                        placeholder=" "
                                    />
                                    <Input
                                        label="開催日"
                                        variant="bordered"
                                        type="date"
                                        placeholder=" "
                                    />
                                    <Input
                                        label="場所"
                                        variant="bordered"
                                        placeholder=" "
                                    />
                                    <div className="mj-select">
                                        <label htmlFor="">
                                            イベントカテゴリ
                                        </label>
                                        <select name="" id="">
                                            <option value="">音楽ライブ</option>
                                            <option value="">公開収録</option>
                                        </select>
                                    </div>
                                    <div className="mj-select">
                                        <label htmlFor="">
                                            出演者
                                        </label>
                                        <select name="" id="">
                                            <option value="">SUPER BEAVER</option>
                                            <option value="">the peggies</option>
                                        </select>
                                    </div>
                                    <Textarea
                                        label="備考"
                                        labelPlacement="inside"
                                        placeholder=" "
                                        defaultValue=""
                                        variant="bordered"
                                    />
                                    <Button className="bg-mj_accent text-white">送信</Button>
                                </ModalBody>
                            </form>
                        </>
                    )}
                </ModalContent>
            </Modal>
        </>
    );
}