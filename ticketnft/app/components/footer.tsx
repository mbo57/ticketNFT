"use client";
import * as React from "react";

import Image from "next/image";
import Logo from "../../public/LogoBlack.png";

export const Footer = () => {

    return (
        <footer className="bg-[#333] h-[30px] flex items-center justify-center">
            <Image src={Logo} alt="mujiqulo" className="h-[20px] w-auto"></Image>
        </footer>
    );
};
