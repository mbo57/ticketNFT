"use client";
import * as React from "react";

// 1. import `NextUIProvider` component
import { NextUIProvider } from "@nextui-org/react";
import {Button} from "@nextui-org/button";

import { HeaderForStaff } from "@/app/components/staff/header";
import { Footer } from "@/app/components/footer";

export default async function StaffPageTop() {
  return (
    <NextUIProvider>
      <HeaderForStaff />
      <Button>Press me</Button>
      <Footer />
    </NextUIProvider>
  );
}