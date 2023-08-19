"use client";
import * as React from "react";

import {
    Navbar,
    NavbarBrand,
    NavbarContent,
    NavbarMenu,
    NavbarMenuItem,
    NavbarMenuToggle,
    NavbarItem,
    Link,
    Button,
} from "@nextui-org/react";

import Image from "next/image";
import Logo from "../../../public/LogoForStaff.png";

export const HeaderForStaff = () => {
    const [isMenuOpen, setIsMenuOpen] = React.useState<boolean>(false);

    const menuItems: string[] = ["イベント管理", "チケット管理", "主催者管理"];

    return (
        <Navbar>
            <NavbarContent>
                <NavbarBrand>
                    <Image
                        src={Logo}
                        alt="mujiqulo for staff"
                        className="h-[50px] w-auto"
                    />
                </NavbarBrand>
            </NavbarContent>

            <NavbarContent justify="end">
                {menuItems.map((item, index) => (
                    <NavbarItem key={`${item}-${index}`}>
                        <Link
                            color= "foreground"
                            className= {
                                index === 2
                                    ? "foreground mj-header-underline"
                                    : "foreground"
                            }
                            href="#"
                            size="lg"
                        >
                            {item}
                        </Link>
                    </NavbarItem>
                ))}
                <NavbarMenuToggle
                    aria-label={isMenuOpen ? "Close menu" : "Open menu"}
                />
            </NavbarContent>

            <NavbarMenu>
                {menuItems.map((item, index) => (
                    <NavbarMenuItem key={`${item}-${index}`}>
                        <Link
                            color="foreground"
                            className={
                                index === 2
                                    ? "w-full text-mj_base"
                                    : "w-full"
                            }
                            href="#"
                            size="lg"
                        >
                            {item}
                        </Link>
                    </NavbarMenuItem>
                ))}
            </NavbarMenu>
        </Navbar>
    );
};
