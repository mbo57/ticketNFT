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

    const menuItems: string[] = ["menu1", "menu2", "menu3"];

    return (
        <Navbar>
            <NavbarContent>
                <NavbarBrand>
                    <Image src={Logo} alt="mujiqulo for staff" className="h-[50px] w-auto" />
                </NavbarBrand>
            </NavbarContent>

            <NavbarContent justify="end">
                {menuItems.map((item, index) => (
                    <NavbarItem key={`${item}-${index}`}>
                        <Link
                            color={
                                index === 2
                                    ? "primary"
                                    : index === menuItems.length - 1
                                    ? "danger"
                                    : "foreground"
                            }
                            className="w-full"
                            href="#"
                            size="lg"
                        >
                            {item}
                        </Link>
                    </NavbarItem>
                ))}
                <NavbarMenuToggle aria-label={isMenuOpen ? "Close menu" : "Open menu"} />
            </NavbarContent>

            <NavbarMenu>
                {menuItems.map((item, index) => (
                    <NavbarMenuItem key={`${item}-${index}`}>
                        <Link
                            color={
                                index === 2
                                    ? "primary"
                                    : index === menuItems.length - 1
                                    ? "danger"
                                    : "foreground"
                            }
                            className="w-full"
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
