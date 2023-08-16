"use client";
import * as React from "react";

// 1. import `NextUIProvider` component
import { NextUIProvider } from "@nextui-org/react";
import {Button} from "@nextui-org/button";

export default async function StaffPageTop() {
  return (
    <NextUIProvider>
      <Button>Press me</Button>
    </NextUIProvider>
  );
}