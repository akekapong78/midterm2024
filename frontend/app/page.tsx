'use client';
import { useUser } from "./providers/userProvider";
import { useEffect, useState } from "react";

export default function Home() {
  const userCtx = useUser();

  return (
    <div>Home: {userCtx.username}, {userCtx.role}</div>
    
  );
}
